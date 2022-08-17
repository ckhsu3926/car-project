package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"

	"car-record/entities"
)

type authorizeMiddleware struct {
	UUsecase entities.UserUsecase
}

func NewAuthorizeMiddleware(uu entities.UserUsecase) *authorizeMiddleware {
	return &authorizeMiddleware{
		UUsecase: uu,
	}
}

const AuthorizeString = "authorize"

func (m *authorizeMiddleware) Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatus(403)
			return
		}

		user, getUserErr := m.UUsecase.Authorize(ctx.Request.Context(), token)
		if getUserErr != nil {
			ctx.AbortWithStatus(403)
			return
		}

		ctx.Set(AuthorizeString, user)

		ctx.Next()
	}
}

func GetMiddlewareAuthorize(ctx *gin.Context) (user entities.User, err error) {
	value, exist := ctx.Get(AuthorizeString)
	if !exist {
		err = errors.New(`authorize not exist`)
		return
	}

	user, isUserType := value.(entities.User)
	if !isUserType {
		err = errors.New(`authorize type error`)
		return
	}

	return
}
