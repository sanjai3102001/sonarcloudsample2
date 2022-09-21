package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
)

const (
	AWS_S3_REGION = "us-east-2"
	AWS_S3_BUCKET = "moviepic"
	AKID          = "AKIATICND5ZYGYCNGVXY"
	SECRET_KEY    = "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"
)

func uploadimagetos3(w http.ResponseWriter, r *http.Request) {
	//session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, ""),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Upload image
	err = uploadFile(session, "test_11zon.jpeg")
	if err != nil {
		log.Fatal(err)
	}
}

func uploadFile(session *session.Session, uploadFileDir string) error {

	upFile, err := os.Open(uploadFileDir)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(AWS_S3_BUCKET),
		Key:                  aws.String(uploadFileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", uploadimagetos3).Methods("PUT")
	log.Fatal(http.ListenAndServe("Localhost:5000", r))

}
