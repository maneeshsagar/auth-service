package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"github.com/maneeshsagar/auth-service/persistence"
)

func AuthrizationMiddleware(persistnece persistence.Persistence) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authBearer := ctx.Request.Header["Authorization"]
		if len(authBearer) <= 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
			ctx.Abort()
			return
		}
		tokenArr := strings.Split(authBearer[0], "Bearer")

		fmt.Println(tokenArr)
		token, err := persistnece.GetToken(strings.TrimSpace(tokenArr[1]))
		if err != nil && err != orm.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "internal server error"})
			ctx.Abort()
			return
		}
		if err == orm.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
			ctx.Abort()
			return
		}

		if token.ExpiresAt.Before(time.Now()) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "token expired"})
			ctx.Abort()
			return
		}

		fmt.Println("::::::::::::::::: ", token.UserID)
		ctx.Set("userId", token.UserID)
		ctx.Next()
	}
}
