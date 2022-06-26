package middlewares

import (
	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(auth *auth.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// bearer := c.Request().Header.Get("Authorization")
			// authorizationToken := strings.Split(bearer, " ")

			// if len(authorizationToken) != 2 {
			// 	return c.String(http.StatusBadRequest, "invalid token")
			// }

			// token, err := auth.VerifyIDToken(context.Background(), authorizationToken[1])
			// if err != nil {
			// 	return c.String(http.StatusUnauthorized, "Unauthorization")
			// }
			// c.Set("uid", token.UID)
			c.Set("uid", "Ejcza0kCm9NMcExGjoX9T45cqAB2")

			return next(c)
		}
	}
}
