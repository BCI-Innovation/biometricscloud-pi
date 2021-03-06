package idos

import "time"

type DeviceCreateRequestIDO struct {
	Manufacturer string `json:"manufacturer"`
	DeviceCode   string `json:"device_code"`
	IsTestMode   bool   `json:"is_test_mode"`
}

type DeviceCreateResponseIDO struct {
	ID           uint64 `json:"id,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	ClientID     string `json:"oauth2_client_id,omitempty"`
	ClientSecret string `json:"oauth2_client_secret,omitempty"`
	RedirectURL  string `json:"oauth2_redirect_url,omitempty"`
}

type MetricCreateRequestIDO struct {
	DeviceID          uint64 `json:"device_id,omitempty"`
	Name              string `json:"name,omitempty"`
	NameOther         string `json:"name_other,omitempty"`
	SampleType        string `json:"sample_type,omitempty"`
	SampleTypeOther   string `json:"sample_type_other,omitempty"`
	QuantityType      string `json:"quantity_type,omitempty"`
	QuantityTypeOther string `json:"quantity_type_other,omitempty"`
	IsTestMode        bool   `json:"is_test_mode,omitempty"`
	IsContinuousData  bool   `json:"is_continuous_data,omitempty"`
}

type MetricResponseIDO struct {
	ID                int    `json:"id,omitempty"`
	UUID              string `json:"uuid,omitempty"`
	TenantID          uint64 `json:"tenant_id,omitempty"`
	DeviceID          uint64 `json:"device_id,omitempty"`
	DeviceUUID        string `json:"device_uuid,omitempty"`
	DeviceName        string `json:"device_name,omitempty"`
	UserID            uint64 `json:"user_id,omitempty"`
	Name              string `json:"name,omitempty"`
	NameOther         string `json:"name_other,omitempty"`
	SampleType        string `json:"sample_type,omitempty"`
	SampleTypeOther   string `json:"sample_type_other,omitempty"`
	QuantityType      string `json:"quantity_type,omitempty"`
	QuantityTypeOther string `json:"quantity_type_other,omitempty"`
	IsTestMode        bool   `json:"is_test_mode,omitempty"`
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
