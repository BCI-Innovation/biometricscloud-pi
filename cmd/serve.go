package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/BCI-Innovation/biometricscloud-pi/internal/app"
	"github.com/BCI-Innovation/biometricscloud-pi/internal/camera"
	"github.com/BCI-Innovation/biometricscloud-pi/internal/idos"
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

	// Initialize the camera hardware.
	cam, err := camera.NewLibCameraStill(cm.ID, width, height, format, workingDirectoryAbsoluteFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the remote service
	rs := remote.New(devRemoteServerAddress, dev.ClientID, dev.ClientSecret, devTokenURL)

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
