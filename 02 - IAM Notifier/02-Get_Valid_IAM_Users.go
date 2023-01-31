package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

//---------------- (STRUCTS) -----------------
type AccessKeyInfo struct{
	AccessKeyMetadata []AccessKeyCreationKeyDate `json:"AccessKeyMetadata"`
}

type AccessKeyCreationKeyDate struct{
	UserName string `json:"UserName"`
	AccessKeyId string `json:"AccessKeyId"`
	Status string `json:"Status"`
	CreateDate string `json:"CreateDate"`
}
//------------- (END STRUCTS) ----------------


//----------------- (METHODS) ----------------
func (u AccessKeyCreationKeyDate) GetAccessKeyCreationDates(user_name string) string{

	var accesskeyinfo AccessKeyInfo
	var return_value string

	res, err := exec.Command("aws", "iam", "list-access-keys", "--user-name", user_name).Output(); CheckForNil(err)
	err = json.Unmarshal(res, &accesskeyinfo); CheckForNil(err)
	
	if len(accesskeyinfo.AccessKeyMetadata) == 0{
		return_value = "No access key"
	}else{
		return_value = accesskeyinfo.AccessKeyMetadata[0].CreateDate[0:10]
	}

	return return_value
}
//-------------- (END METHODS) -------------


//--------------- (FUNCTION) ---------------
func FilterUsers(allUsers AllUsers) {

	var accesskeycreationkeydate AccessKeyCreationKeyDate

	var users []string
	var accesskey_dates []string

	for i := 0; i < allUsers.GetUserCount(); i++ {
		if allUsers.GetValedUsers(i) != "" {
			users = append(users, allUsers.GetValedUsers(i))
		}
	}

	fmt.Println("")
	check_status := "(1/2)Checking ."
	for i := 0; i < len(users); i++ {
		accesskey_dates = append(accesskey_dates, accesskeycreationkeydate.GetAccessKeyCreationDates(users[i]))
		fmt.Printf(check_status)
		check_status = "."
	}

	fmt.Println("")
	SendNotifications(users, accesskey_dates)

	fmt.Println("")
	fmt.Printf("\n%v\n","Completed \u2714")
}
//------------- (END FUNCTION) --------------