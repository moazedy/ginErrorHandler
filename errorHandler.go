package ginErrorHandler

import "github.com/gin-gonic/gin"

type GinError struct {
	Message string `json:"error_message"`
	Code    int    `json:"error_code"`
}

func (g *GinError) Error() string {
	return g.Message
}

func HandleTheError(ctx *gin.Context, err GinError) {
	ctx.JSON(err.Code, gin.H{
		"error_message": err.Message,
	})
}
