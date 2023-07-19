package workers

type Workers interface {
	Start()

	Stop()

	IsDone() bool
}
