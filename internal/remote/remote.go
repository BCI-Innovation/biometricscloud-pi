package remote

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"

	"github.com/BCI-Innovation/biometricscloud-pi/internal/constants"
	"github.com/BCI-Innovation/biometricscloud-pi/internal/utils"
	"golang.org/x/oauth2/clientcredentials"
)

type Remote struct {
	serverAddress string
	clientID      string
	clientSecret  string
	tokenURL      string
	cfg           clientcredentials.Config
}

func New(ra string, cid string, cs string, turl string, cfg clientcredentials.Config) *Remote {
	return &Remote{
		serverAddress: ra,
		clientID:      cid,
		clientSecret:  cs,
		tokenURL:      turl,
		cfg:           cfg,
	}
}

func (r *Remote) SubmitPhotoSample(metricID int, bytes []byte, filePath string) error {
	// Convert bytes to base64 string
	content, err := utils.GetURLEncodedBase64StringFromImageBytes(bytes)
	if err != nil {
		return err
	}

	// File meta.
	filename := path.Base(filePath)
	fileExt := path.Ext(filePath)

	// Upload the data.
	return r.uploadData(content, filename, fileExt, metricID)
}

// Note: https://github.com/BCI-Innovation/biometricscloud-backend/blob/master/internal/idos/photo_sample.go

type PhotoSampleCreateRequestIDO struct {
	MetricID                  int       `json:"metric_id,omitempty"`                    // Required
	StartDate                 time.Time `json:"start_date,omitempty"`                   // Required
	EndDate                   time.Time `json:"end_date,omitempty"`                     // Required
	SampleMotionContext       string    `json:"sample_motion_context,omitempty"`        // Optional
	SampleSensorLocation      int8      `json:"sample_sensor_location,omitempty"`       // Optional
	SampleSensorLocationOther string    `json:"sample_sensor_location_other,omitempty"` // Optional
	UploadContent             string    `json:"upload_content"`
	UploadFilename            string    `json:"upload_filename"`
}

func (r *Remote) uploadData(content string, filename string, fileExt string, metricID int) error {
	// Generate the URL we will be making the submission to.
	aURL := r.serverAddress + constants.PhotoSampleListCreateEndpointURL
	// Generate our payload.
	data := &PhotoSampleCreateRequestIDO{
		MetricID:       metricID,
		StartDate:      time.Now(),
		EndDate:        time.Now(),
		UploadContent:  content,
		UploadFilename: filename,
	}

	// Get a token.
	// https://github.com/go-oauth2/oauth2/blob/b208c14e621016995debae2fa7dc20c8f0e4f6f8/example/client/client.go#L116
	t, err := r.cfg.Token(context.Background())
	if err != nil {
		return err
	}

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + t.AccessToken

	dataBytes, _ := json.Marshal(data)
	requestBodyBuf := bytes.NewBuffer(dataBytes)

	// Start preparing the sending code...
	client := &http.Client{}
	req, _ := http.NewRequest("POST", aURL, requestBodyBuf)

	// Add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Set data format.
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Read the response body
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(utils.JsonPrettyPrint(string(responseBytes)))
	return nil
}
