package ginErrorHandler

import "github.com/gin-gonic/gin"

type GinError struct {
	Message string `json:"error_message"`
	Code    int    `json:"error_code"`
}

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

func HandleTheError(ctx *gin.Context, err GinError) {
	ctx.JSON(err.Code, gin.H{
		"error_message": err.Message,
	})
}
