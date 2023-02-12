package function

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
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

func RemoteFunction(w http.ResponseWriter, r *http.Request) {
	bqReq := &RemoteFunctionRequest{}
	bqRes := &RemoteFunctionResponse{}

	if err := json.NewDecoder(r.Body).Decode(&bqReq); err != nil {
		bqRes.ErrorMessage = fmt.Sprintf("Cannot parse request body: %v", err)
	} else {
		// for debugging
		input, _ := json.Marshal(bqReq)
		fmt.Println(string(input))

		wg := &sync.WaitGroup{}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		results := make([]interface{}, len(bqReq.Calls))

		for i, args := range bqReq.Calls {
			wg.Add(1)
			go func(i int, args []interface{}) {
				defer wg.Done()
				for {
					select {
					case <-ctx.Done():
						return
					default:
						result, err := Call(args)
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

	fmt.Println(string(b)) // for debugging

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func Call(args []interface{}) (interface{}, error) {
	if l := len(args); l != 1 {
		return nil, fmt.Errorf("unexpected number of inputs. expected 1, got %d", l)
	}
	s, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("unexpected type of input. expected a string")
	}
	return len(s), nil
}
