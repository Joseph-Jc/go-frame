package utils

type Task struct {
	handleFunc func()
}

func NewTask(handleFunc func()) *Task {
	return &Task{handleFunc: handleFunc}
}

func (t *Task) Execute() {
	t.handleFunc()
}

type Pool struct {
	EntryChannel chan *Task
	JobsChannel  chan *Task
	workNum      int
}

func NewPool(num int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		workNum:      num,
	}
}

func (p *Pool) work() {
	for task := range p.JobsChannel {
		task.Execute()
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.workNum; i++ {
		go p.work()
	}
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}
