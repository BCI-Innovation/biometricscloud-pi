package camera

import (
	"errors"
	"fmt"
	_ "image/jpeg" // Add support for jpeg format.
	_ "image/png"  // Add support for png format.
	"os/exec"
	"strconv"
	"time"
)

// LibCameraStill represents the Golang wrapper over the `libcamera-still` command line application that is found in the Raspberry Pi OS. For more information about this library visit: https://github.com/raspberrypi/libcamera-apps/blob/main/apps/libcamera_still.cpp
type LibCameraStill struct {
	// metricID is the  device camera attached to the user account which is a unique ID value registered at biometricscloud.net web-service.
	metricID   int
	workingDir string
	format     string
	fileExt    string
	width      int
	height     int
}

// NewLibCameraStill creates an instance of our LibCameraStill struct.
func NewLibCameraStill(metricID int, width int, height int, format string, workingDirectoryAbsoluteFilePath string) (*LibCameraStill, error) {
	// See available formats via https://www.raspberrypi.com/documentation/accessories/camera.html#encoders
	fileFormat := map[string]string{
		"png":    "png",
		"jpg":    "jpg",
		"bmp":    "bmp",
		"rgb":    "data",
		"yuv420": "data",
	}
	if fileExt, ok := fileFormat[format]; ok {
		return &LibCameraStill{
			metricID:   metricID,
			workingDir: workingDirectoryAbsoluteFilePath,
			format:     format,
			fileExt:    fileExt,
			width:      width,
			height:     height,
		}, nil
	}
	return nil, errors.New("File format does not exist")
}

// TakeSnapshot will take a snapshot with the Raspberry Pi camera module and save it to the specified file. This function is essentially a wrapper function over the `libcamera-still` command.
func (cam *LibCameraStill) TakeSnapshot() (string, error) {
	// Generate the new filename for our camera still and save.
	filename := cam.workingDir + "/" + strconv.Itoa(int(time.Now().Unix())) + "." + cam.fileExt

	// DEVELOPERS NOTE:
	// We are using the included `libcamera-still` command to handle taking a
	// snapshot of the camera and saving it to local file.
	// https://www.raspberrypi.com/documentation/accessories/camera.html#libcamera-still

	app := "libcamera-still"
	args := []string{
		"--width", strconv.Itoa(cam.width),
		"--height", strconv.Itoa(cam.height),
	}

	args = append(args, []string{"-e", cam.format}...)
	args = append(args, []string{"-o", filename}...)

	cmd := exec.Command(app, args...)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	fmt.Println(string(stdout))
	return filename, nil
}

// GetMetricID returns the unique ID value assigned by biometricscloud for this camera that is registered to a specific user's account.
func (cam *LibCameraStill) GetMetricID() int {
	return cam.metricID
}
