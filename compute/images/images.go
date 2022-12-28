package images

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var TOKEN = "gAAAAABjrF1WE-wQRLFWujRO3QX6qT9g84itU9gLBEGmk7zh4Tlce5Rov9VldmXMpB2D4l1AeBbrgsfWkPw06KZ4IRpv1pVCzzF3EsEznlr1SdCOCBV525xCOFoW06F_jM9R91Xz0_TMeKiuEMSN51TlvYHz4ILG7N3OvMFHLP42HKKLImmWkLc"


//images struct
type Images_Struct struct {
	Images []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
			Type string `json:"type,omitempty"`
		} `json:"links"`
	} `json:"images"`
}

// images list
func getAllImages() Images_Struct{
	URL := fmt.Sprintf("http://voscontrol:8774/v2.1/images")
    
	fmt.Println("----------------All Images----------------")

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token",TOKEN)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Images_Struct
	json.Unmarshal(bodyBytes, &responseObject)
	
	return responseObject
}

// print all images
func showImages() {
	images := getAllImages()
	for _, value := range images.Images {
		fmt.Printf("Name: %+v\n", value.Name)
	}
}