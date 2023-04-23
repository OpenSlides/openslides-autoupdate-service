package authtest

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// ValidTokens creates valid tokens to be used in auth requests
func ValidTokens(cookieKey, headerKey []byte, userID int) (cookie *http.Cookie, headerName, headerValue string, err error) {
	cookieToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": "123",
	}).SignedString(cookieKey)
	if err != nil {
		return nil, "", "", fmt.Errorf("create cookie token: %w", err)
	}
	cookie = &http.Cookie{
		Name:  "refreshId",
		Value: cookieToken,
	}

	headerValue, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userID,
		"sessionId": "123",
	}).SignedString(headerKey)
	if err != nil {
		return nil, "", "", fmt.Errorf("Can not sign token token: %w", err)
	}
	headerValue = "bearer " + headerValue

	return cookie, "Authentication", headerValue, nil
}
