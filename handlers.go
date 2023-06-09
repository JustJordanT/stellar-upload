package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

const (
	filePath = "files"
)

func awsConfig(key string, secret_key string) aws.Config {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret_key, "")),
		// config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKIA6RM7A22NOYZH7JXM", "sSSmoL5NqzFj6kKiX0i6QzM4Tex94+WOVrfVK91D", "")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}

func s3Config(key string, secret_key string) *s3.Client {

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(awsConfig(key, secret_key))
	return client
}

func listObjectsHandler(c echo.Context) error {

	// Specify the S3 bucket name
	bucketName := "interstellar-block"

	// List objects in the bucket
	resp, err := s3Config(KeyAuth(c), SecretKeyAuth(c)).ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Println("Failed to list objects in S3 bucket:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		// return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Extract object keys from the response
	objectKeys := make([]string, 0, len(resp.Contents))
	for _, obj := range resp.Contents {
		objectKeys = append(objectKeys, *obj.Key)
	}

	return c.JSON(http.StatusOK, objectKeys)
}

func uploadFileHandler(c echo.Context) error {

	// Specify the S3 bucket name
	bucketName := "interstellar-block"

	// Retrieve the uploaded file from the form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to retrieve file from request")
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to open file")
	}
	defer src.Close()

	dstKet := filePath + "/" + file.Filename

	params := &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &dstKet,
		Body:   src,
	}

	_, err = s3Config(KeyAuth(c), SecretKeyAuth(c)).PutObject(context.TODO(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to upload file"})
	}
	fmt.Println("HERE")
	return c.JSON(http.StatusOK, map[string]string{"ok": fmt.Sprintf("File '%s' uploaded successfully", file.Filename)})
}

func deleteHandler(c echo.Context) error {
	return nil
}

func downloadHandler(c echo.Context) error {
	return nil
}
