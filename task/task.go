package task

type Tasker interface {
	Process() error
}
