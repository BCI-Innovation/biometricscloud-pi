package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/BCI-Innovation/biometricscloud-pi/internal/idos"
	"github.com/BCI-Innovation/biometricscloud-pi/internal/remote"
)

var (
	filePath string
)

func init() {
	uploadCmd.Flags().StringVarP(&filePath, "filePath", "a", "", "-")
	uploadCmd.MarkFlagRequired("filePath")
	rootCmd.AddCommand(uploadCmd)
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload the saved file on hard disk",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		doRunUpload()
	},
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func doRunUpload() {
	// For debugging purposes.
	fmt.Println("filePath:", filePath)
	fmt.Println("width:", width)
	fmt.Println("height:", height)
	fmt.Println("format:", format)
	fmt.Println("wkgrp:", workingDirectoryAbsoluteFilePath)
	fmt.Println("remote:", devRemoteServerAddress)
	fmt.Println("devClientID:", devClientID)
	fmt.Println("devClientSecret:", devClientSecret)
	fmt.Println("devTokenURL:", devTokenURL)

	// Read our device.
	deviceBytes, err := ioutil.ReadFile("./device.json")
	if err != nil {
		log.Fatal(err)
	}
	dev := &idos.DeviceCreateResponseIDO{}
	err = json.Unmarshal(deviceBytes, &dev)
	if err != nil {
		log.Fatal(err)
	}

	// Read our camera.
	cameraBytes, err := ioutil.ReadFile("./camera.json")
	if err != nil {
		log.Fatal(err)
	}
	cm := &idos.MetricResponseIDO{}
	err = json.Unmarshal(cameraBytes, &cm)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := os.ReadFile(filePath) // Load entire file into memory.
	if err != nil {
		log.Fatal(err)
	}

	rs := remote.New(devRemoteServerAddress, dev.ClientID, dev.ClientSecret, devTokenURL)
	err = rs.SubmitPhotoSample(cm.ID, bytes, filePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully uploaded the image file at", filePath)
}
