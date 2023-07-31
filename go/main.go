package main

import (
	"fmt"
	"sync"
	"time"
)

type PromiseMap struct {
	sync.Mutex
	Data map[string]chan interface{}
}

func (pm *PromiseMap) Get(key string) (value chan interface{}, found bool) {
	pm.Lock()
	defer pm.Unlock()
	value, found = pm.Data[key]
	return value, found
}

func (pm *PromiseMap) Set(key string, value chan interface{}) {
	pm.Lock()
	defer pm.Unlock()
	pm.Data[key] = value
}

func DBFetch(key string) interface{} {
	time.Sleep(1 * time.Second)
	fmt.Println("DBFetch", key)
	return key
}

var OnePromiseForAllMap PromiseMap = PromiseMap{Data: make(map[string]chan interface{})}

func OnePromiseForAll(recordId string, pm *PromiseMap) chan interface{} {
	promise, found := pm.Get(recordId)
	if found {
		return promise
	}

	resultChan := make(chan interface{})
	pm.Set(recordId, resultChan)
	go func() {
		resultChan <- DBFetch(recordId)
	}()
	return resultChan
}

var requests = []string{"1", "1", "3", "3"}

var promiseMap sync.Map

func OnePromiseForAll2(recordId string) chan interface{} {
	promise, found := promiseMap.Load(recordId)
	if found {
		return promise.(chan interface{})
	}

	resultChan := make(chan interface{})
	promiseMap.Store(recordId, resultChan)
	go func() {
		resultChan <- DBFetch(recordId)
	}()
	return resultChan
}

func DBRequest(Id string) chan interface{} {
	ch := make(chan interface{})
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Printf("DBRequest: %s\n", Id)
		ch <- struct{}{}
	}()
	return ch
}

// func fetchOnePromiseForAlDlDB(Id string) chan interface{} {
// 	if ch, ok := promiseMap[Id]; ok {
// 		return ch
// 	}
// 	ch := DBRequest(Id)
// 	promiseMap[Id] = ch
// 	go func() {
// 		<-ch
// 		delete(promiseMap, Id)
// 	}()
// 	return ch
// }

func main() {
	var resultChan chan interface{}
	for _, request := range requests {
		go func(request string) {
			resultChan = OnePromiseForAll2(request)
			fmt.Println("resultChan", <-resultChan)
		}(request)
	}

	time.Sleep(2 * time.Second)
}
