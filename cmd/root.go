package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2/clientcredentials"
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
	devMetricID                      int
)

func init() {
	// Extract the envrionment variables which our application will be using.
	ra := os.Getenv("BIOMETRICSCLOUD_PI_REMOTE_SERVER_ADDRESS")
	ci := os.Getenv("BIOMETRICSCLOUD_PI_CLIENT_ID")
	cs := os.Getenv("BIOMETRICSCLOUD_PI_CLIENT_SECRET")
	tu := os.Getenv("BIOMETRICSCLOUD_PI_TOKEN_URL")
	w, err := strconv.ParseInt(os.Getenv("BIOMETRICSCLOUD_PI_WIDTH"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	ht, err := strconv.ParseInt(os.Getenv("BIOMETRICSCLOUD_PI_HEIGHT"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	f := os.Getenv("BIOMETRICSCLOUD_PI_FORMAT")
	wg := os.Getenv("BIOMETRICSCLOUD_PI_WK_GRP")
	mid, err := strconv.ParseInt(os.Getenv("BIOMETRICSCLOUD_PI_DEVICE_CAMERA_METRIC_ID"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

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
	rootCmd.PersistentFlags().IntVar(&width, "weight", int(w), "-")
	rootCmd.PersistentFlags().IntVar(&height, "height", int(ht), "-")
	rootCmd.PersistentFlags().StringVar(&format, "format", f, "")
	rootCmd.PersistentFlags().StringVar(&workingDirectoryAbsoluteFilePath, "workGroup", wg, "./")
	rootCmd.PersistentFlags().IntVar(&devMetricID, "metricID", int(mid), "-")
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
