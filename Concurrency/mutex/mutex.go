package mutex

import (
	"fmt"
	"sync"
)

var message string
var wg sync.WaitGroup

func setGlobalMessage(m string, mtx *sync.Mutex) {
	defer wg.Done()
	mtx.Lock()
	message = m
	mtx.Unlock()
}
func main() {
	var mtx sync.Mutex
	wg.Add(2)
	go setGlobalMessage("I am first one trying to access global message", &mtx)
	go setGlobalMessage("I am second one trying to access global message", &mtx)
	wg.Wait()

	fmt.Println(message)
}
