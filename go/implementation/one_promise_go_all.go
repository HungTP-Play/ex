package implementation

import (
	"fmt"
	"sync"
	"time"
)

type Promise interface{}

type PromiseMap struct {
	sync.Mutex
	Data map[string]Promise
}

func (pm *PromiseMap) Get(key string) (value Promise, found bool) {
	pm.Lock()
	defer pm.Unlock()
	value, found = pm.Data[key]
	return value, found
}

func (pm *PromiseMap) Set(key string, value Promise) {
	pm.Lock()
	defer pm.Unlock()
	pm.Data[key] = value
}

func DBFetch(key string) Promise {
	time.Sleep(1 * time.Second)
	fmt.Println("DBFetch", key)
	return key
}

var OnePromiseForAllMap PromiseMap = PromiseMap{Data: make(map[string]Promise)}

func OnePromiseForAll(recordId string, pm *PromiseMap) Promise {
	promise, found := pm.Get(recordId)
	if found {
		return promise
	}

	resultChan := make(chan Promise)
	go func() {
		defer pm.Set(recordId, <-resultChan)
		resultChan <- DBFetch(recordId)
	}()

	return <-resultChan
}
