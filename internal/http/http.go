// Package http provides HTTP handler to give other services access to to
// permission service.
package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const prefix = "/internal/permission"

// IsAlloweder provides the IsAllowed method.
type IsAlloweder interface {
	IsAllowed(ctx context.Context, name string, userID int, dataList [](map[string]json.RawMessage)) (bool, error)
}

// IsAllowed registers a handler, to connect to the IsAllowed method.
//
// It returns the string `true` or `false` that can be encoded as json.
//
// If an error happens, a json-error-string is returned with status code 500.
func IsAllowed(mux *http.ServeMux, provider IsAlloweder) {
	mux.Handle(prefix+"/is_allowed", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			jsonError(w, "Can't read request body: "+err.Error())
			return
		}

		var requestData struct {
			Name     string                         `json:"name"`
			UserID   int                            `json:"user_id"`
			DataList [](map[string]json.RawMessage) `json:"data"`
		}
		if err := json.Unmarshal(b, &requestData); err != nil {
			jsonError(w, fmt.Sprintf("Can not decode request body '%s': %v", b, err))
			return
		}

		allowed, err := provider.IsAllowed(r.Context(), requestData.Name, requestData.UserID, requestData.DataList)

		if err != nil {
			jsonError(w, err.Error())
			return
		}

		value := "false"
		if allowed {
			value = "true"
		}
		fmt.Fprintln(w, value)
	}))
}

type allrouter interface {
	AllRoutes() ([]string, []string)
}

// Health registers a handler, that tells, if the service is running.
//
// It also returns all known collections and actions.
func Health(mux *http.ServeMux, router allrouter) {
	mux.Handle(prefix+"/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var rData struct {
			Info struct {
				Routes struct {
					Collections []string `json:"collections"`
					Actions     []string `json:"actions"`
				} `json:"routes"`
			} `json:"healthinfo"`
		}
		rData.Info.Routes.Collections, rData.Info.Routes.Actions = router.AllRoutes()
		if err := json.NewEncoder(w).Encode(rData); err != nil {
			jsonError(w, "Something went wrong")
			return
		}
	}))
}

// jsonError writes an error to the client as json object.
func jsonError(w http.ResponseWriter, msg string) {
	b, err := json.Marshal("Internal Error. Norman, Do not sent it to client: " + msg)
	if err != nil {
		b = []byte(`"Very internal error"`)
	}

	w.WriteHeader(500)
	w.Write(b)
}
