package function

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// App is exposed as a function
func App(w http.ResponseWriter, r *http.Request) {
	payload := map[string]interface{}{
		"runtime": os.Getenv("GCF_RUNTIME"),
		"key":     "value",
		"array":   []int{1, 2, 3},
	}
	output, _ := json.Marshal(payload)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	mw := io.MultiWriter(w, os.Stdout)
	fmt.Fprintf(mw, "%s\n", output)
}
