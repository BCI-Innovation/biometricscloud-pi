package camera

type CameraStillReader interface {
	TakeSnapshot() (string, error)
	GetMetricID() int
}
