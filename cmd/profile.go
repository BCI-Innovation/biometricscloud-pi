package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BCI-Innovation/biometricscloud-pi/internal/constants"
	"github.com/BCI-Innovation/biometricscloud-pi/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(profileCmd)
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Retrieve profile details.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// https://github.com/go-oauth2/oauth2/blob/b208c14e621016995debae2fa7dc20c8f0e4f6f8/example/client/client.go#L116
		token, err := cfg.Token(context.Background())
		if err != nil {
			log.Println("Retrieving profile details failed with error:")
			log.Fatal(err)
		}

		aURL := applicationAddress + constants.ProfileEndpointURL

		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + accessToken

		client := &http.Client{}
		req, _ := http.NewRequest("GET", aURL, nil)

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

		fmt.Println(utils.JsonPrettyPrint(string(responseBytes)))
	},
}
