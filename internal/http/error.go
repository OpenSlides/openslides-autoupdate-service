package http

import "fmt"

type invalidRequestError struct {
	err error
}

func (e invalidRequestError) Error() string {
	return fmt.Sprintf("Invalid request: %v", e.err)
}

func (e invalidRequestError) Type() string {
	return "invalid_request"
}

// logoutError is sent when a session is terminated due to server-initiated logout
// (e.g., backchannel logout from Keycloak).
type logoutError struct{}

func (e logoutError) Error() string {
	return "Session logged out"
}

func (e logoutError) Type() string {
	return "logout"
}

// Reason returns the reason identifier for the client to handle logout redirection.
func (e logoutError) Reason() string {
	return "Logout"
}
