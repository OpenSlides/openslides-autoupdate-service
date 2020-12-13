// Package http provides HTTP handler to give other services access to to
// permission service.
package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const prefix = "/internal/permission"

// IsAlloweder provides the IsAllowed method.
type IsAlloweder interface {
	IsAllowed(ctx context.Context, name string, userID int, dataList [](map[string]json.RawMessage)) ([](map[string]interface{}), error)
}

// IsAllowed registers a handler, to connect to the IsAllowed method.
func IsAllowed(mux *http.ServeMux, provider IsAlloweder) {
	url := prefix + "/is_allowed"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")

		bodyBytes, readErr := ioutil.ReadAll(r.Body)
		if readErr != nil {
			handleError(w, fmt.Errorf("Can't read response body: %w", readErr))
		}

		var requestData struct {
			Name     string                         `json:"name"`
			UserID   int                            `json:"user_id"`
			DataList [](map[string]json.RawMessage) `json:"data"`
		}
		if err := json.Unmarshal(bodyBytes, &requestData); err != nil {
			handleError(w, jsonError{fmt.Sprintf("Can not decode request body '%s'", string(bodyBytes)), err})
			return
		}

		additions, err := provider.IsAllowed(r.Context(), requestData.Name, requestData.UserID, requestData.DataList)

		// get reason from ClientError
		reason := ""
		var responseData interface{}
		if err != nil {
			var clientError interface {
				Type() string
			}
			if errors.As(err, &clientError) && clientError.Type() == "ClientError" {
				reason = err.Error()
			} else {
				handleError(w, fmt.Errorf("calling IsAllowed: %w", err))
				return
			}

			var indexError interface {
				Index() int
			}
			var errorIndex int
			if errors.As(err, &indexError) {
				errorIndex = indexError.Index()
			}

			responseData = struct {
				Allowed    bool   `json:"allowed"`
				Reason     string `json:"reason"`
				ErrorIndex int    `json:"error_index"`
			}{
				false,
				reason,
				errorIndex,
			}
		} else {
			responseData = struct {
				Allowed   bool                       `json:"allowed"`
				Additions [](map[string]interface{}) `json:"additions"`
			}{
				true,
				additions,
			}
		}

		if err := json.NewEncoder(w).Encode(responseData); err != nil {
			handleError(w, fmt.Errorf("decoding response: %w", err))
			return
		}
	})

	mux.Handle(url, handler)
}

// Health registers a handler, that tells, if the service is running.
func Health(mux *http.ServeMux) {
	url := prefix + "/health"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		fmt.Fprintln(w, `{"healthy": true}`)
	})

	mux.Handle(url, handler)
}

func handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/octet-stream")

	errorData := struct {
		Type string `json:"type"`
		Msg  string `json:"msg"`
	}{
		"InternalError",
		"Ups, something went wrong!",
	}

	status := 500
	var clientError interface {
		Type() string
	}
	if errors.As(err, &clientError) {
		status = 400
		errorData.Type = clientError.Type()
		errorData.Msg = err.Error()
	}

	log.Printf("Error %s status=%d: %v\n", errorData.Type, status, err)
	w.WriteHeader(status)

	jsonData, err := json.Marshal(errorData)
	if err == nil {
		fmt.Fprintf(w, `{"error":%s}`, jsonData)
	} else {
		fmt.Fprintf(w, `{"error": {"type": "InternalError", "msg": "Cannot write error response"}}`)
	}

	fmt.Fprintln(w)
}
