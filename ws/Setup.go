package ws

import "matthewhope/real-time-forum/repo"

var (
	ClientPool *Pool
)

func Setup(repo repo.IRepository) {
	ClientPool = NewPool(repo)
	go ClientPool.Run()
}
