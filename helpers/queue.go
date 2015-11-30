package helpers

type Queue struct {
	WorkNum int //最大工作数量
	msg     chan interface{}
	Done    bool
}

type QueueHander interface {
	ProducerHandler()
	ConsumerHandler()
}

/**
 * 实例化
 */
func New(workNum int) *Queue {
	queue := new(Queue)
	queue.WorkNum = workNum
	return queue
}

/**
 *生产者
 */
func (this *Queue) Producer(handler func(int)) {
	for i := 0; i < this.WorkNum; i++ {
		this.msg <- handler(i)
	}
	done <- true
}

/**
 *消费者
 */
func (this *Queue) Consumer(handler func(interface{})) {
	for msg := range this.msg {
		handler(msg)
	}
}
