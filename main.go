package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	// Input parameters
	uploadFile := flag.String("file", "", "absolute path to the file which should be uploaded to s3")
	profileName := flag.String("profile", "admin", "aws profile name defined in ~/.aws/config")
	bucketName := flag.String("bucket", "", "bucket name where the file should be upload to")
	flag.Parse()

	sess, err := session.NewSessionWithOptions(session.Options{
		// enable shared config support.
		SharedConfigState: session.SharedConfigEnable,

		// Optionally set the profile to use from the shared config.
		Profile: *profileName,
	})
	if err != nil {
		log.Fatal(err)
	}

	uploader := s3manager.NewUploader(sess)

	// Open file
	f, err := os.Open(*uploadFile)
	if err != nil {
		log.Fatal(err)
	}

	// Upload the file to s3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:               aws.String(*bucketName),
		Key:                  aws.String(filepath.Base(*uploadFile)),
		ServerSideEncryption: aws.String("AES256"),
		Body:                 f,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file uploaded to %s\n", result.Location)
}
