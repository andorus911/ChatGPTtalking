package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
  // do not forget to create .env file!
	godotenv.Load()

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Printf("S3: %s and secret key: %s", s3Bucket, secretKey)
}
