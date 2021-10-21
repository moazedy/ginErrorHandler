package ginErrorHandler

import "github.com/gin-gonic/gin"

// GinError is a custom type for error to be handled in HTTP
type GinError struct {
	Message string `json:"error_message"`
	Code    int    `json:"error_code"`
}

// implementing error type interface for GinError
func (g *GinError) Error() string {
	return g.Message
}

// construction func to declare GinError as an error
func New(message string, code int) error {
	return &GinError{
		Message: message,
		Code:    code,
	}
}

// HandleTheError gets gin context refferenc and also an error as input arguments
// and responses to HTTP client using the error parameters
// TODO: allways call return, after calling this function in your scope.
func HandleTheError(ctx *gin.Context, err GinError) {
	ctx.JSON(err.Code, gin.H{
		"error_message": err.Message,
	})
}
