package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/BCI-Innovation/biometricscloud-pi/internal/app"
	"github.com/BCI-Innovation/biometricscloud-pi/internal/camera"
	"github.com/BCI-Innovation/biometricscloud-pi/internal/remote"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "-",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		doRunServe()
	},
}

func doRunServe() {
	// For debugging purposes only.
	fmt.Println("width:", width)
	fmt.Println("height:", height)
	fmt.Println("format:", format)
	fmt.Println("workingDirectoryAbsoluteFilePath:", workingDirectoryAbsoluteFilePath)
	fmt.Println("devRemoteServerAddress:", devRemoteServerAddress)
	fmt.Println("devClientID:", devClientID)
	fmt.Println("devClientSecret:", devClientSecret)
	fmt.Println("devTokenURL:", devTokenURL)
	fmt.Println()

	// Initialize the camera.
	cam, err := camera.NewLibCameraStill(devMetricID, width, height, format, workingDirectoryAbsoluteFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the remote service
	rs := remote.New(devRemoteServerAddress, devClientID, devClientSecret, devTokenURL, cfg)

	// Initialize our application.
	app, err := app.New(cam, rs)
	if err != nil {
		log.Fatal(err)
	}

	defer app.StopMainRuntimeLoop()

	// DEVELOPERS CODE:
	// The following code will create an anonymous goroutine which will have a
	// blocking chan `sigs`. This blocking chan will only unblock when the
	// golang app receives a termination command; therfore the anyomous
	// goroutine will run and terminate our running application.
	//
	// Special Thanks:
	// (1) https://gobyexample.com/signals
	// (2) https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	//
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs // Block execution until signal from terminal gets triggered here.
		fmt.Println("Starting graceful shut down now.")
		app.StopMainRuntimeLoop()
	}()

	app.RunMainRuntimeLoop()
}
