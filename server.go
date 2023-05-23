package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

func main() {

	// output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	// 	Bucket: aws.String("interstellar-block"),
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("All Objects in infytech2 bucket")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/time", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().UTC().String())
	})

	e.GET("/list", listObjectsHandler)

	// e.GET("/list", func(c echo.Context) error {

	// 	for _, object := range output.Contents {
	// 		("key=%s size=%d", aws.ToString(object.Key), object.Size)
	// 	}

	// 	return c.String(http.StatusOK, "Done.")
	// })

	e.Logger.Fatal(e.Start(":8080"))
}

func listObjectsHandler(c echo.Context) error {

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
