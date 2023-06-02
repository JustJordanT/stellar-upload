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

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("", "", "")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)
	// Create a new AWS session
	// sess, err := session.newsession(&aws.config{
	// 	region: aws.string("us-west-2"), // update with your desired aws region
	// })
	// if err != nil {
	// 	log.println("failed to create aws session:", err)
	// 	return c.string(http.statusinternalservererror, "internal server error")
	// }

	// Create a new S3 service client
	// svc := s3.New(sess)

	// Specify the S3 bucket name
	bucketName := "interstellar-block"

	// List objects in the bucket
	resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Println("Failed to list objects in S3 bucket:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Extract object keys from the response
	objectKeys := make([]string, 0, len(resp.Contents))
	for _, obj := range resp.Contents {
		objectKeys = append(objectKeys, *obj.Key)
	}

	return c.JSON(http.StatusOK, objectKeys)
}
