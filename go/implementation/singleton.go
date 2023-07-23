package implementation

import "sync"

type Singleton struct{}

var instance *Singleton
var mu sync.Mutex

func GetInstance() *Singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
