// Package http provides HTTP handler to give other services access to to
// permission service.
package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const prefix = "/internal/permission"

// IsAlloweder provides the IsAllowed method.
type IsAlloweder interface {
	IsAllowed(ctx context.Context, name string, userID int, data map[string]json.RawMessage) (bool, map[string]interface{}, error)
}

// IsAllowed registers a handler, to connect to the IsAllowed method.
func IsAllowed(mux *http.ServeMux, provider IsAlloweder) {
	url := prefix + "/is_allowed"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")

		var requestData struct {
			Name   string                     `json:"name"`
			UserID int                        `json:"user_id"`
			Data   map[string]json.RawMessage `json:"data"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			handleError(w, jsonError{"Can not decode request body", err})
			return
		}

		allowed, addition, err := provider.IsAllowed(context.TODO(), requestData.Name, requestData.UserID, requestData.Data)
		if err != nil {
			handleError(w, fmt.Errorf("calling IsAllowed: %w", err))
			return
		}

		responseData := struct {
			Allowed  bool                   `json:"allowed"`
			Addition map[string]interface{} `json:"addition"`
		}{
			allowed,
			addition,
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

	errType := "InternalError"
	errMsg := "Ups, something went wrong!"
	status := 500
	var clientError interface {
		Type() string
	}
	if errors.As(err, &clientError) {
		status = 400
		errType = clientError.Type()
		errMsg = err.Error()
	} else {
		log.Printf("Error: %v", err)
	}

	w.WriteHeader(status)
	fmt.Fprintf(w, `{"error": {"type": "%s", "msg": "%s"}}`, errType, errMsg)
	fmt.Fprintln(w)
}
