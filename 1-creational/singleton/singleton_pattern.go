package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}
var once sync.Once

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance already created")
		}
	} else {
		fmt.Println("single instance already created")
	}
	return singleInstance
}

func getInstanceWithOnce() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("single instance already created")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 20; i++ {
		// getInstance()
		getInstanceWithOnce()
	}
}
