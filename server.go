package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	UnauthorizedMessage   = "Unauthorized: KEY and SECRET_KEY Headers must be set."
	CallerIdentityMessage = "Issue validating credentials with aws"
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

		// TODO: conditional check for ^ these headers.
		// TODO: check if the secret key and secret key are actually able to authenticate
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion("us-west-2"),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secretKey, "")),
		)
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		svc := sts.NewFromConfig(cfg)
		input := &sts.GetCallerIdentityInput{}
		_, err = svc.GetCallerIdentity(context.TODO(), input)
		if err != nil {
			log.Println("Failed to authenticate:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return next(c)
	}
}
