package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	util "github.com/Nowfloats/Bizapp/Utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/yaml.v2"
)

type FullForm struct{
	FormData *FormData `json:"formData,omitempty"`
	UploadLogoImage *ImageFile `json:"uploadlogoimage,omitempty"`
}

type FormData struct {
	Id                    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FPTag                 string	`json:"fptag,omitempty"`
	FPRootAliasURI        string	`json:"fprootaliasuri,omitempty"`
	FPLogoURI             string	`json:"fplogouri,omitempty"`
	FPLogoBackgroundColor string	`json:"fplogobackgroundcolor,omitempty"`
	FPLogoForegroundColor string	`json:"fplogoforegroundcolor,omitempty"`	
	Version               string	`json:"version,omitempty"`
}

type ImageFile struct{
	LogoImage []byte `json:"uploadLogoImage,omitempty"`
}

var wg sync.WaitGroup

//================================ Generate .env file ===================================
func (fd FullForm) GenerateEnvFile(messagedetails FormData){
	wg.Done()
	
	folder_name := fmt.Sprintf("%v/V%v/.env" ,messagedetails.FPTag, messagedetails.Version)

	envFile := map[string]string{
		"FPTAG": messagedetails.FPTag,
		"FPROOTALIASURI": messagedetails.FPRootAliasURI,
		"FPLOGOURI": messagedetails.FPLogoURI,
		"FPLOGOBACKGROUNDCOLOR": messagedetails.FPLogoBackgroundColor,
		"FPLOGOFOREGROUNDCOLOR": messagedetails.FPLogoForegroundColor,
		"VERSION": messagedetails.Version,
	}

	envFileInBytes, err := yaml.Marshal(envFile); util.CheckForNil(err)
	fmt.Printf("%v: %v\n",messagedetails.FPTag,"Preparing .envfile... ✔️")
	
	reader := bytes.NewReader(envFileInBytes)

	util.S3(folder_name, reader)
}
//=============================== END Generate .env file =================================



//============================= Generate pubspec.yaml file ===============================
func (fd FullForm) GeneratePubSpec(messagedetails FormData){
	wg.Done()

	message_version := messagedetails.Version
	folder_name := fmt.Sprintf("%v/V%v/pubspec.yaml" ,messagedetails.FPTag, message_version)

	pubspecfile := util.YamlGenerator(message_version)
	fmt.Printf("%v: %v\n",messagedetails.FPTag,"Preparing pubspecfile... ✔️")
	
	reader := bytes.NewReader(pubspecfile)

	util.S3(folder_name, reader)
}
//============================ END Generate pubspec.yaml file ==============================



//=========================== Generate name and version file ===============================
func (fd FullForm) NameAndVersion(messagedetails FormData){
	wg.Done()

	name_version := messagedetails.FPTag+"\nV"+messagedetails.Version
	folder_name := fmt.Sprintf("%v/V%v/name_version.txt" ,messagedetails.FPTag, messagedetails.Version)
	
	replacer := strings.NewReplacer(`"`,"")
	name_version = replacer.Replace(name_version)
	fmt.Printf("%v: %v\n",messagedetails.FPTag,"Generating metadata... ✔️")
	
	reader := bytes.NewReader([]byte(name_version))

	util.S3(folder_name, reader)
}
//========================= END Generate name and version file =============================



//=========================== Getting image for further use ================================
func (fd FullForm) Uploadimage(messagedetails FullForm){
	wg.Done()

	folder_name := fmt.Sprintf("%v/V%v/icon.png" ,messagedetails.FormData.FPTag, messagedetails.FormData.Version)

	fmt.Printf("%v: %v\n",messagedetails.FormData.FPTag,"Uploading image... ✔️")
	
	reader := bytes.NewReader(messagedetails.UploadLogoImage.LogoImage)

	util.S3(folder_name, reader)
}
//========================== END Getting image for further use =============================



//=================================== Send to Queue ========================================
var sqsClient, sqsQueueUrl = util.SQS()

func SendToQueue(message []byte) string{

	message_to_string := string(message)
	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{QueueUrl: sqsQueueUrl.QueueUrl, MessageBody: aws.String(message_to_string)}); util.CheckForNil(err)

	PollFromQueue()

	return "Request made successfully"
}
//================================= End Send to Queue =======================================



//================================== Poll from Queue ========================================
func PollFromQueue(){

	var messagedetails FullForm

	max_message_count, err := strconv.ParseInt(os.Getenv("MAX_MESSAGE_COUNT"),10, 64); util.CheckForNil(err)
	message, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{QueueUrl: sqsQueueUrl.QueueUrl, MaxNumberOfMessages: aws.Int64(max_message_count)})

	err = json.Unmarshal([]byte(*message.Messages[0].Body), &messagedetails); util.CheckForNil(err)
	
	wg.Add(4)

	go messagedetails.GenerateEnvFile(*messagedetails.FormData)
	go messagedetails.GeneratePubSpec(*messagedetails.FormData)
	go messagedetails.NameAndVersion(*messagedetails.FormData)
	go messagedetails.Uploadimage(messagedetails)

	wg.Wait()
}
//================================ End Poll from Queue ========================================

