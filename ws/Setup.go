package ws

var (
	ClientPool *Pool
)

func Setup() {
	ClientPool = NewPool()
	go ClientPool.Run()
}
