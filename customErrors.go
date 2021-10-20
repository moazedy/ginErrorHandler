package ginErrorHandler

import "net/http"

// error for not found data
var NotFound = GinError{
	Message: "requested data not found",
	Code:    http.StatusNotFound,
}

// error for internal server issues
var InternalError = GinError{
	Message: "internal server error",
	Code:    http.StatusInternalServerError,
}

// error for unauthorized setuations
var Unauthorized = GinError{
	Message: "requester is unauthorized",
	Code:    http.StatusUnauthorized,
}

// error for access denied conditions
var AccessDenided = GinError{
	Message: "access deined",
	Code:    http.StatusForbidden,
}

// wrong input error
var WrongInput = GinError{
	Message: " the input you've entered is wrong !",
	Code:    http.StatusBadRequest,
}
