package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	config "github.com/Nowfloats/Bizapp/Configs"
	util "github.com/Nowfloats/Bizapp/Utils"
)

type Form struct{
	FormData FormData `json:"formdata,omitempty"`
	UploadLogoImage ImageFile `json:"uploadlogoimage,omitempty"`
}

type FormData struct {
	FPTag                 string	`json:"fptag,omitempty"`
	FPRootAliasURI        string	`json:"fprootaliasuri,omitempty"`
	FPLogoURI             string	`json:"fplogouri,omitempty"`
	FPLogoBackgroundColor string	`json:"fplogobackgroundcolor,omitempty"`
	FPLogoForegroundColor string	`json:"fplogoforegroundcolor,omitempty"`	
	Version               string	`json:"version,omitempty"`
}

type ImageFile struct{
	UploadLogoImage []byte `json:"uploadLogoImage,omitempty"`
}

//=================== Home Handler ========================
func HomeHandler(w http.ResponseWriter, r *http.Request){

	tmpl := util.InitHtmlFiles()
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
//================= End Home Handler =======================



//==================== Health Check ========================
func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	currentTime := time.Now()
	timeFormatted := currentTime.Format("15:04 Monday")

	fmt.Fprintf(w,"Checking at %q...\n\n%v", timeFormatted,"Healthy ✔️") 
}
//================== END Health Check =======================



//===================== Form handler ========================
func CreateApp(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var completeForm Form

	r.ParseMultipartForm(10 << 20)

	GetFormFile, _ , err := r.FormFile("UPLOADLOGOIMAGE"); util.CheckForNil(err)
	defer GetFormFile.Close()
	image_file, err := ioutil.ReadAll(GetFormFile); util.CheckForNil(err)

	formDataRaw := map[string]string{
		"FPTAG": r.FormValue("FPTAG"),
		"FPROOTALIASURI": r.FormValue("FPROOTALIASURI"),
		"FPLOGOURI": r.FormValue("FPLOGOURI"),
		"FPLOGOBACKGROUNDCOLOR": r.FormValue("FPLOGOBACKGROUNDCOLOR"),
		"FPLOGOFOREGROUNDCOLOR": r.FormValue("FPLOGOFOREGROUNDCOLOR"),
		"VERSION": r.FormValue("VERSION"),
		"IMAGE_NAME": r.FormValue("IMAGE_NAME"),
	}

	formFile := map[string][]byte{
		"UPLOADLOGOIMAGE": image_file,
	}

	formDataToByte, err:= json.Marshal(formDataRaw); util.CheckForNil(err)
	formFileToByte, err:= json.Marshal(formFile); util.CheckForNil(err)

	json.Unmarshal(formDataToByte, &completeForm.FormData)
	json.Unmarshal(formFileToByte, &completeForm.UploadLogoImage)

	CompleteFormToBytes, err :=  json.Marshal(&completeForm); util.CheckForNil(err)

	queue_res := config.SendToQueue(CompleteFormToBytes)

	json.NewEncoder(w).Encode(string(queue_res))
}
//=================== End Form handler ======================


















// //================= Create Application =================
// func CreateApp(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type","application/json")

// 	res, err := io.ReadAll(r.Body); util.CheckForNil(err)
// 	queue_res := config.SendToQueue(res)
	
// 	json.NewEncoder(w).Encode(string(queue_res))
// }
// //=============== END Create Application ==============



// //==================== Upload Image ===================
// func UploadImage(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type","application/json")

// 	r.ParseMultipartForm(10 << 20)
// 	file, image_details, err := r.FormFile("UPLOADLOGOIMAGE"); util.CheckForNil(err)
// 	defer file.Close()

// 	image_path, err := filepath.Abs(image_details.Filename); util.CheckForNil(err)
// 	fmt.Printf("Uploading image from: %v\n",image_path)
// 	fmt.Printf("Image total size: %vKB\n",image_details.Size)

// 	fileBytes, err := ioutil.ReadAll(file); util.CheckForNil(err)

// 	reader := bytes.NewReader(fileBytes)
// 	util.S3("icon.png", reader)
// }
// //================== END Upload Image ================