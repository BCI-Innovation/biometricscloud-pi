package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BCI-Innovation/biometricscloud-pi/internal/constants"
	"github.com/BCI-Innovation/biometricscloud-pi/internal/idos"
	"github.com/spf13/cobra"
)

var (
	at string
)

func init() {
	registerCmd.Flags().StringVarP(&at, "access_token", "a", "", "The access token to register with.")
	registerCmd.MarkFlagRequired("accessToken")
	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register our device",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		doRegister()
	},
}

func doRegister() {
	// Create our device.
	_, device := createDevice()

	// Save the device details to a JSON file.
	file, _ := json.MarshalIndent(device, "", " ")
	_ = ioutil.WriteFile("device.json", file, 0644)

	// Create our device.
	_, metric := createMetric(device)

	// Save the metric details to a JSON file.
	file, _ = json.MarshalIndent(metric, "", " ")
	_ = ioutil.WriteFile("camera.json", file, 0644)

	// DEVELOPERS NOTE:
	// If you want to add more metrics then follow the pattern as above.

	fmt.Println("Device registered")
}

func createDevice() ([]byte, *idos.DeviceCreateResponseIDO) {
	aURL := devRemoteServerAddress + constants.DeviceListCreateEndpointURL
	data := idos.DeviceCreateRequestIDO{
		Manufacturer: "BCI Innovation",
		DeviceCode:   "biodev",
		IsTestMode:   false,
	}

	dataBytes, _ := json.Marshal(data)
	requestBodyBuf := bytes.NewBuffer(dataBytes)

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + at

	client := &http.Client{}
	req, _ := http.NewRequest("POST", aURL, requestBodyBuf)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Set data format.
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer resp.Body.Close()

	// Read the response body
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll | An Error Occured %v", err)
	}

	var responseData idos.DeviceCreateResponseIDO

	// De-serialize bytes into our struct object.
	err = json.Unmarshal(responseBytes, &responseData)
	if err != nil {
		log.Println(string(responseBytes))
		log.Fatalf("Unmarshal | An Error Occured %v", err)
	}

	return responseBytes, &responseData
}

func createMetric(device *idos.DeviceCreateResponseIDO) ([]byte, *idos.MetricResponseIDO) {
	aURL := devRemoteServerAddress + constants.MetricListCreateEndpointURL
	deviceID := device.ID
	data := &idos.MetricCreateRequestIDO{
		DeviceID:         deviceID,
		Name:             "Rasbperry Pi Camera Module",
		SampleType:       "camera",
		QuantityType:     "jpg",
		IsTestMode:       false,
		IsContinuousData: false,
	}

	dataBytes, _ := json.Marshal(data)
	requestBodyBuf := bytes.NewBuffer(dataBytes)

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + at

	client := &http.Client{}
	req, _ := http.NewRequest("POST", aURL, requestBodyBuf)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Set data format.
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer resp.Body.Close()

	// Read the response body
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll | An Error Occured %v", err)
	}

	var responseData idos.MetricResponseIDO

	// De-serialize bytes into our struct object.
	err = json.Unmarshal(responseBytes, &responseData)
	if err != nil {
		log.Println(string(responseBytes))
		log.Fatalf("Unmarshal | An Error Occured %v", err)
	}

	return responseBytes, &responseData
}
