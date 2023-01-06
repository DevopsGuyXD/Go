package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Handling URL")

	my_url := "https://www.google.com/search?q=lamborghini&sxsrf=ALiCzsZ6-yRH24bDpeMpJd-dHWDwws-2fw%3A1671534008501&source=hp&ei=uJWhY-XuG--12roP99--mA8&iflsig=AJiK0e8AAAAAY6GjyJgnCSVq-OURkNdjrLC3GpXfGyLB&oq=lamboring&gs_lcp=Cgdnd3Mtd2l6EAMYADIHCAAQgAQQCjIKCAAQgAQQsQMQCjIHCAAQgAQQCjINCAAQgAQQsQMQyQMQCjIKCAAQgAQQsQMQCjIHCAAQgAQQCjIKCAAQgAQQsQMQCjIKCAAQgAQQsQMQCjIHCAAQgAQQCjIKCAAQgAQQsQMQCjoHCCMQ6gIQJzoECCMQJzoFCAAQkQI6CAguEIMBELEDOhEILhCABBCxAxCDARDHARDRAzoICC4Q1AIQgAQ6CAgAEIAEELEDOgUIABCxAzoLCAAQgAQQsQMQgwE6CAguEIAEELEDOg4ILhCABBCxAxDHARDRAzoFCAAQgAQ6BQguEIAEOggILhCxAxCDAToLCC4QxwEQrwEQkQI6CAgAELEDEIMBOgsIABCABBCxAxDJAzoLCC4QgAQQsQMQgwE6CwguEIAEELEDENQCOhEILhCABBCxAxDHARCvARDUAjoOCC4QxwEQsQMQ0QMQgAQ6EAguEIAEELEDEMcBENEDEApQ1gdY8yNgzixoAnAAeACAAZkBiAHvCpIBBDAuMTCYAQCgAQGwAQo&sclient=gws-wiz"

	result, err := url.Parse(my_url); if err != nil{
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Println(qparams["q"])
}