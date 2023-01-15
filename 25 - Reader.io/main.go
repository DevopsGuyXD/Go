package main

import (
	"bytes"
	"context"
	"io"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func main() {

	testFile := "This is a test message"
	testFileBytes := []byte(testFile)

	UploadtoStorage(testFileBytes)
}

func UploadtoStorage(testFileBytes []byte){

	var reader io.Reader = bytes.NewReader(testFileBytes)

	client, err := azblob.NewClientFromConnectionString("DefaultEndpointsProtocol=https;AccountName=jiwwfapicommonsastaging;AccountKey=agRLqHqO8tb7k6keXvmS78mv8GMjygt8qPrZy+DdKHZcs4+txKcilZGwXtdhb+QZ5hTnECax75s0+AStVI5PPA==;EndpointSuffix=core.windows.net", nil)
	if err != nil{ 
		log.Fatal(err) 
	}

	_, err = client.UploadStream(context.TODO(),"boost-kit-resource-production", "testfile.txt", reader, nil)
	if err != nil{ 
		log.Fatal(err) 
	}	

	log.Println("Uploading sitemap to Azure blob storage successfully")

	if err != nil {
		log.Println("Azure blob storage upload error")
	}
}