package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	width    int
	height   int
	format   string
	filename string
)

func init() {
	snapshotCmd.Flags().IntVarP(&width, "width", "a", 1640, "Width of the image")
	// snapshotCmd.MarkFlagRequired("width")
	snapshotCmd.Flags().IntVarP(&height, "height", "b", 1232, "Width of the image")
	// snapshotCmd.MarkFlagRequired("height")
	snapshotCmd.Flags().StringVarP(&format, "type", "c", "png", "Type of image")
	// snapshotCmd.MarkFlagRequired("type")
	snapshotCmd.Flags().StringVarP(&filename, "filename", "d", "image.png", "The filename to save")
	// snapshotCmd.MarkFlagRequired("filename")
	rootCmd.AddCommand(snapshotCmd)
}

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Take a snapshot from the Raspberry Pi camera module",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Taking snapshot from camera now...")
		doRunSnapshot()
		fmt.Println("Successfully took snapshot from camera.")
	},
}

func doRunSnapshot() {
	// DEVELOPERS NOTE:
	// We are using the included `libcamera-still` command to handle taking a
	// snapshot of the camera and saving it to local file.
	// https://www.raspberrypi.com/documentation/accessories/camera.html#libcamera-still

	app := "libcamera-still"
	args := []string{
		"--width", strconv.Itoa(width),
		"--height", strconv.Itoa(height),
	}

	args = append(args, []string{"-e", format}...)
	args = append(args, []string{"-o", filename}...)

	cmd := exec.Command(app, args...)
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(stdout))
}
