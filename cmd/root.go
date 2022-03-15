package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	devRemoteServerAddress string
	devClientID            string
	devClientSecret        string
	devTokenURL            string
	cfg                    clientcredentials.Config
)

func init() {
	// Extract the envrionment variables which our application will be using.
	ra := os.Getenv("BIOMETRICSCLOUD_PI_REMOTE_SERVER_ADDRESS")
	ci := os.Getenv("BIOMETRICSCLOUD_PI_CLIENT_ID")
	cs := os.Getenv("BIOMETRICSCLOUD_PI_CLIENT_SECRET")
	tu := os.Getenv("BIOMETRICSCLOUD_PI_TOKEN_URL")

	// Setup our oAuth2 client credentials authentication client.
	cfg = clientcredentials.Config{
		ClientID:     ci,
		ClientSecret: cs,
		Scopes:       []string{"all"},
		TokenURL:     tu,
	}

	// Attach
	rootCmd.PersistentFlags().StringVar(&devRemoteServerAddress, "remoteServerAddress", ra, "The address of BiometricsCloud remote server.")
	rootCmd.PersistentFlags().StringVar(&devClientID, "clientID", ci, "The oAuth2 client ID of this device.")
	rootCmd.PersistentFlags().StringVar(&devClientSecret, "clientSecret", cs, "TThe oAuth2 client secret of this device.")
	rootCmd.PersistentFlags().StringVar(&devTokenURL, "tokenURL", tu, "The oAuth2 token API endpoint of the platform")
}

var rootCmd = &cobra.Command{
	Use:   "bmcpi",
	Short: "IoT client for BiometricsCloud",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing.
	},
}

// Execute is the main entry into the application from the command line terminal.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
