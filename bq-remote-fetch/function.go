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
	Replies      []interface{} `json:"replies,omitempty"`
	ErrorMessage string        `json:"errorMessage,omitempty"`
}

var sem = semaphore.NewWeighted(3)

func init() {
	log.SetFlags(0)
}

func RemoteFunction(w http.ResponseWriter, r *http.Request) {
	bqReq := &RemoteFunctionRequest{}
	bqRes := &RemoteFunctionResponse{}

	if err := json.NewDecoder(r.Body).Decode(&bqReq); err != nil {
		bqRes.ErrorMessage = fmt.Sprintf("Cannot parse request body: %v", err)
	} else {
		log.Printf("%s, %s, %d\n", bqReq.RequestId, bqReq.Caller, len(bqReq.Calls))

		wg := &sync.WaitGroup{}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		results := make([]interface{}, len(bqReq.Calls))

		for i, args := range bqReq.Calls {
			wg.Add(1)
			go func(i int, args []interface{}) {
				defer wg.Done()

				if err := sem.Acquire(ctx, 1); err != nil {
					bqRes.ErrorMessage = fmt.Sprintf("Error occured for call %d: %v", i, err)
					cancel()
					return
				}
				defer sem.Release(1)

				for {
					select {
					case <-ctx.Done():
						return
					default:
						result, err := Call(ctx, args)
						if err != nil {
							bqRes.ErrorMessage = fmt.Sprintf("Error occured for call %d: %v", i, err)
							cancel()
							return
						}
						results[i] = result
						return
					}
				}
			}(i, args)
		}

		wg.Wait()
		if bqRes.ErrorMessage != "" {
			bqRes.Replies = nil
		} else {
			bqRes.Replies = results
		}
	}

	b, err := json.Marshal(bqRes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot convert to json: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func Call(ctx context.Context, args []interface{}) (interface{}, error) {
	if l := len(args); l != 1 {
		return nil, fmt.Errorf("unexpected number of inputs. expected 1, got %d", l)
	}
	s, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("unexpected type of input. expected a string")
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	u, _ := url.Parse("https://bookmark.hatenaapis.com/count/entry")
	q := u.Query()
	q.Add("url", s)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "bq-remote-function")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("%s, %s", u.String(), string(body))
	time.Sleep(1 * time.Second)

	return strconv.ParseInt(string(body), 10, 64)
}
