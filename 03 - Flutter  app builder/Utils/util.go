package util

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
)

var tmpl *template.Template

//=================== Error Checking ====================
func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
//================= END Error Checking===================




//================= Env file initiation =================
func InitEnvFile(){
	err := godotenv.Load(".env"); CheckForNil(err)
}
//=============== END Env file initiation ===============




//================== AWS initialization =================
func InitAwsSession() (*session.Session){
	InitEnvFile()

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
		},
	}); CheckForNil(err)
		
	return sess
}
//=============== END AWS initialization ================



func InitHtmlFiles() *template.Template{

	tmpl = template.Must(template.ParseGlob("Static/*.html"))

	return tmpl
}



//======================== SQS ==========================
func SQS() (*sqs.SQS, *sqs.GetQueueUrlOutput){
	InitEnvFile()

	sess := InitAwsSession()
	queue_name := os.Getenv("QUEUE_NAME")

	sqsClient := sqs.New(sess)
	queue_url, err := sqsClient.GetQueueUrl(
		&sqs.GetQueueUrlInput{
			QueueName: &queue_name,
		},
	); CheckForNil(err)

	return sqsClient,
		   queue_url
}
//====================== END SQS =========================




//========================= S3 ===========================
func S3(folder_name string, reader *bytes.Reader) *s3manager.UploadOutput{
	InitEnvFile()

	sess := InitAwsSession()
	s3_bucket := os.Getenv("S3_Bucket")

	uploader := s3manager.NewUploader(sess)

	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3_bucket),
		Key: aws.String(folder_name),
		Body: reader,
	}); CheckForNil(err)

	return res
}
//====================== END S3 ==========================