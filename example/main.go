package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wdfky/mtmtz"
)

type T2 struct {
	ActId     string `json:"actId"`
	LinkType  string `json:"linkType"`
	Sid       string `json:"sid"`
	SkuViewId string `json:"skuViewId"`
}

func main() {
	data := T2{
		ActId:     "0",
		LinkType:  "3",
		Sid:       "test",
		SkuViewId: "324124123",
	}
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a request
	url := "https://media.meituan.com/cps_open/common/api/v1/get_referral_link"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Create a SignUtil instance
	signUtil := mtmtz.NewSignUtil("appkey", "app_secret")

	// Get sign headers
	signHeaders := signUtil.GetSignHeaders(map[string]interface{}{
		"method": "POST",
		"data":   data,
		"url":    url,
	})

	//
	fmt.Printf("%+v\n", signHeaders)
	req.Header.Set("S-Ca-App", signHeaders.SCaApp)
	req.Header.Set("S-Ca-Timestamp", signHeaders.SCaTimestamp)
	req.Header.Set("S-Ca-Signature", signHeaders.SCaSignature)
	req.Header.Set("S-Ca-Signature-Headers", signHeaders.SCaSignatureHeaders)
	req.Header.Set("Content-MD5", signHeaders.ContentMD5)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print response
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	// Read response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println("response Body:", buf.String())
}
