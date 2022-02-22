package driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpColor struct {
	Color string `json:"Color"`
}

func callColorServer(raspiHost, deviceName string) (string,error) {
	//host := "localhost:4000"
	//if true {
	//	return "red",nil
	//}
	resp, err := http.Get(fmt.Sprintf("http://%s/api/v2/device/name/%s/Color",raspiHost,deviceName))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(body))
	tmpColor := &httpColor{}
	if err := json.Unmarshal(body,tmpColor);err != nil {
		return "",err
	}
	return tmpColor.Color,nil
}


func setDeviceColor(raspiHost, deviceName ,val string) (string,error) {
	jsonVal, err := json.Marshal(httpColor{
		Color: val,
	})
	if err != nil {
		return "", err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/api/v2/device/name/%s/Color",raspiHost,deviceName),"application/json",bytes.NewReader(jsonVal))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body),nil
}
