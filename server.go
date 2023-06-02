package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	UnauthorizedMessage   = "Unauthorized: KEY and SECRET_KEY Headers must be set."
	CallerIdentityMessage = "Issue validating cradentials with aws"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(S3MiddlewareAuth)

	e.GET("/list", listObjectsHandler)
	e.POST("/upload", uploadFileHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func S3MiddlewareAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Request().Header.Get("KEY")
		secretKey := c.Request().Header.Get("SECRET_KEY")

		if len(key) == 0 || len(secretKey) == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": UnauthorizedMessage})
		}

		srv := sts.NewFromConfig(awsConfig(key, secretKey))

		// Add authentication logic here (e.g., validating keys)
		input := &sts.GetCallerIdentityInput{}
		_, err := srv.GetCallerIdentity(context.TODO(), input)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": CallerIdentityMessage})
		}

		return next(c)
	}
}