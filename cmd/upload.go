package cmd

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/spf13/cobra"

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
	bytes, err := os.ReadFile(filePath) // Load entire file into memory.
	if err != nil {
		log.Fatal(err)
	}

	r := remote.New(devRemoteServerAddress, devClientID, devClientSecret, devTokenURL, cfg)
	err = r.SubmitPhotoSample(devMetricID, bytes, filePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully uploaded the image file at", filePath)
}
