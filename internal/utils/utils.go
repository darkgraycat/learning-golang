package utils

import (
	"math/rand"
	"sync"
	"time"
)

func DoStuff(wg *sync.WaitGroup, callback func()) {
	defer wg.Done()

	time.Sleep(time.Second * time.Duration(rand.Float32()))
	callback()
}
