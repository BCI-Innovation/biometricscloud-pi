package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/BCI-Innovation/biometricscloud-pi/internal/constants"
)

var (
	width                            int
	height                           int
	format                           string
	workingDirectoryAbsoluteFilePath string
	devRemoteServerAddress           string
	devClientID                      string
	devClientSecret                  string
	devTokenURL                      string
	cfg                              clientcredentials.Config
)

func init() {
	// Extract the envrionment variables which our application will be using.
	ra := os.Getenv("BIOMETRICSCLOUD_PI_REMOTE_SERVER_ADDRESS")
	if ra == "" {
		log.Fatal("BIOMETRICSCLOUD_PI_REMOTE_SERVER_ADDRESS: D.N.E.")
	}
	devTokenURL = constants.TokenEndpointURL
	ci := os.Getenv("BIOMETRICSCLOUD_PI_CLIENT_ID")
	cs := os.Getenv("BIOMETRICSCLOUD_PI_CLIENT_SECRET")
	w, err := strconv.ParseInt(os.Getenv("BIOMETRICSCLOUD_PI_WIDTH"), 10, 64)
	if err != nil {
		w = 1640
	}
	ht, err := strconv.ParseInt(os.Getenv("BIOMETRICSCLOUD_PI_HEIGHT:"), 10, 64)
	if err != nil {
		ht = 1232
	}
	f := os.Getenv("BIOMETRICSCLOUD_PI_FORMAT")
	if f != "" {
		f = "jpg"
	}
	wg := os.Getenv("BIOMETRICSCLOUD_PI_WK_GRP")
	if wg == "" {
		wg = "./"
	}

	// Setup our oAuth2 client credentials authentication client.
	cfg = clientcredentials.Config{
		ClientID:     ci,
		ClientSecret: cs,
		Scopes:       []string{"all"},
		TokenURL:     devTokenURL,
	}

	// Attach
	rootCmd.PersistentFlags().StringVar(&devRemoteServerAddress, "remoteServerAddress", ra, "The address of BiometricsCloud remote server.")
	rootCmd.PersistentFlags().StringVar(&devClientID, "clientID", ci, "The oAuth2 client ID of this device.")
	rootCmd.PersistentFlags().StringVar(&devClientSecret, "clientSecret", cs, "TThe oAuth2 client secret of this device.")
	rootCmd.PersistentFlags().IntVar(&width, "weight", int(w), "-")
	rootCmd.PersistentFlags().IntVar(&height, "height", int(ht), "-")
	rootCmd.PersistentFlags().StringVar(&format, "format", f, "")
	rootCmd.PersistentFlags().StringVar(&workingDirectoryAbsoluteFilePath, "workGroup", wg, "./")
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
