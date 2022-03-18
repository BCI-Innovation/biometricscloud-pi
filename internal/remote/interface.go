package remote

type RemoteService interface {
	SubmitPhotoSample(metricID int, bytes []byte, filePath string) error
}
