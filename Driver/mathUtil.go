package main

import (
	"math/rand"
	"sync"
	"time"
)

var randomMutex sync.Mutex
//随机数互斥锁（确保随机数函数不能被并发访问（因为是5个考生同时访问


func GetRandomInt(start ,end int) int {
	randomMutex.Lock()
	<- time.After(1*time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := start + r.Intn(end - start+ 1)
	randomMutex.Unlock()
	return n
}
