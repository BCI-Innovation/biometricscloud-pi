package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/BCI-Innovation/biometricscloud-pi/internal/camera"
)

func init() {
	snapCmd.Flags().IntVarP(&width, "width", "a", 1640, "Override initial width setting of the image")
	snapCmd.Flags().IntVarP(&height, "height", "b", 1232, "Override initial height setting of the image")
	snapCmd.Flags().StringVarP(&format, "format", "c", "jpg", "Override initial format setting of image")
	snapCmd.Flags().StringVarP(&workingDirectoryAbsoluteFilePath, "workingDir", "d", "./", "The absolute file path to the directory where all photos are saved")
	rootCmd.AddCommand(snapCmd)
}

var snapCmd = &cobra.Command{
	Use:   "snap",
	Short: "Snap a single photo with the camera",
	Long:  `-`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Taking a photo...")

		// Initialize the camera.
		cam, err := camera.NewLibCameraStill(0, width, height, format, workingDirectoryAbsoluteFilePath)
		if err != nil {
			log.Fatal(err)
		}

		// Take a photo from the camera and save it.
		_, err = cam.TakeSnapshot()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Photo was successfully taken!")
	},
}
