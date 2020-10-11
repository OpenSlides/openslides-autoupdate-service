package auth_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/auth"
)

func TestFromContext(t *testing.T) {
	a := new(auth.Auth)

	t.Run("Empty Context", func(t *testing.T) {
		got := a.FromContext(context.Background())
		if got != 0 {
			t.Errorf("Got uid %d from empty context. Expected 0", got)
		}
	})

	t.Run("Context from Authenticate", func(t *testing.T) {
		ctx, cancel, err := a.Authenticate(&http.Request{})
		if err != nil {
			t.Fatalf("Can not create context from Authenticate: %v", err)
		}
		defer cancel()

		got := a.FromContext(ctx)
		if got != 1 {
			t.Errorf("Got uid %d from auth-context. Expected 1", got)
		}

	})
}
