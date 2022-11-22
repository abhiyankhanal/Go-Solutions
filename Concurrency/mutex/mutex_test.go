package mutex

import (
	"sync"
	"testing"
)

func Test_setGlobalMessage(t *testing.T) {
	var mtx sync.Mutex
	message = "Testing"
	wg.Add(2)
	go setGlobalMessage("test message first", &mtx)
	go setGlobalMessage("test message second", &mtx)
	wg.Wait()
	if message != "test message first" {
		t.Error("Did not got expected value")
	}
}
