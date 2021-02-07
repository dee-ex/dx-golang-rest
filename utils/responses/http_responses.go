package responses

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func TextResponse(w http.ResponseWriter, status int, message string) {
    w.WriteHeader(status)
    fmt.Fprintf(w, message)
}

func JSONResponse(w http.ResponseWriter, status int, message interface{}) {
    response, err := json.Marshal(message)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write([]byte(response))
}
