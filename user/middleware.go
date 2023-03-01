package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {

		var token string
		context.ShouldBindJSON(&token)
		if token == "" {
			data := tokenStature{
				Code:    412,
				Message: "token为空",
			}
			context.JSON(http.StatusPreconditionFailed, data)
			context.Abort()
			return
		} else {
			claim, err := ParseToken(token)
			if err != nil {
				data := tokenStature{
					Code:    500,
					Message: "token错误",
				}
				context.JSON(http.StatusInternalServerError, data)
				context.Abort()
				return
			} else if time.Now().Unix() > claim.ExpiresAt {
				data := tokenStature{
					Code:    500,
					Message: "token已过期",
				}
				context.JSON(http.StatusInternalServerError, data)
				context.Abort()
				return
			}
		}
		context.Next()
	}
}
