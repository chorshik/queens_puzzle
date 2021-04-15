package thread

// import "sync"

// // Thread
// type Thread struct {
// 	body func()
// 	mux  sync.Mutex
// }

// // NewThread ...
// func NewThread(body func()) *Thread {
// 	return &Thread{
// 		body: body,
// 	}
// }

// // Start ...
// func (thread *Thread) Start() {
// 	thread.mux.Lock()
// 	go thread.run()
// }

// // Join ...
// func (thread *Thread) Join() {
// 	thread.mux.Lock()
// 	thread.mux.Unlock()
// }

// func (thread *Thread) run() {
// 	thread.body()
// 	thread.mux.Unlock()
// }
