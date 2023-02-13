package function

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type RemoteFunctionRequest struct {
	RequestId          string            `json:"requestId"`
	Caller             string            `json:"caller"`
	SessionUser        string            `json:"sessionUser"`
	UserDefinedContext map[string]string `json:"userDefinedContext"`
	Calls              [][]interface{}   `json:"calls"`
}

type RemoteFunctionResponse struct {
	Replies      []int64 `json:"replies,omitempty"`
	ErrorMessage string  `json:"errorMessage,omitempty"`
}

var sem = semaphore.NewWeighted(3)

func init() {
	log.SetFlags(0)
}

func Function(w http.ResponseWriter, r *http.Request) {
	bqReq := &RemoteFunctionRequest{}
	bqRes := &RemoteFunctionResponse{}

	if err := json.NewDecoder(r.Body).Decode(&bqReq); err != nil {
		bqRes.ErrorMessage = fmt.Sprintf("Cannot parse request body: %v", err)
	} else {
		log.Printf("%s, %s, %d\n", bqReq.RequestId, bqReq.Caller, len(bqReq.Calls))
		CollectBookmarkCount(bqReq, bqRes)
	}

	if bqRes.ErrorMessage != "" {
		bqRes.Replies = nil
	}

	b, err := json.Marshal(bqRes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot convert to json: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func CollectBookmarkCount(req *RemoteFunctionRequest, res *RemoteFunctionResponse) {
	// validate
	urls := make([]string, len(req.Calls))
	for i, args := range req.Calls {
		if l := len(args); l != 1 {
			res.ErrorMessage = fmt.Sprintf("Unexpected number of inputs for call %d. expected 1, got %d", i, l)
			return
		}
		s, ok := args[0].(string)
		if !ok {
			res.ErrorMessage = fmt.Sprintf("Unexpected type of input for call %d. expected string", i)
			return
		}
		_, err := url.Parse(s)
		if err != nil {
			res.ErrorMessage = fmt.Sprintf("Unexpected input, input must be a url for call %d", i)
			return
		}
		urls[i] = s
	}

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make([]int64, len(urls))

	for i, u := range urls {
		wg.Add(1)
		go func(i int, u string) {
			defer wg.Done()

			if err := sem.Acquire(ctx, 1); err != nil {
				res.ErrorMessage = fmt.Sprintf("Error occured for call %d: %v", i, err)
				cancel()
				return
			}
			defer sem.Release(1)

			for {
				select {
				case <-ctx.Done():
					return
				default:
					result, err := Call(ctx, u)
					if err != nil {
						res.ErrorMessage = fmt.Sprintf("Error occured for call %d: %v", i, err)
						cancel()
						return
					}
					results[i] = result
					return
				}
			}
		}(i, u)
	}

	wg.Wait()
	res.Replies = results
}

func Call(ctx context.Context, entryURL string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	u, _ := url.Parse("https://bookmark.hatenaapis.com/count/entry")
	q := u.Query()
	q.Add("url", entryURL)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", "bq-remote-function")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	log.Printf("%s, %s", entryURL, string(body))

	return strconv.ParseInt(string(body), 10, 64)
}
