package listerr

import "net/http"

var (
	UNAUTHORIZED                 = NewError(http.StatusUnauthorized, "Unauthorized", "UNAUTHORIZED")
	USERNAME_OR_PASSWORD_INVALID = NewError(http.StatusBadRequest, BAD_REQUEST, "USERNAME_OR_PASSWORD_INVALID")
)
