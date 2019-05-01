package main

import (
	"fmt"
	// "reflect"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int = 0

var gw sync.WaitGroup

type Task interface {
	Run(param int) (int, error)
}

type DownloadTask struct {
	File string
}

func (t DownloadTask) Run(param int) (int, error) {
	mu.Lock()
	defer gw.Done()
	defer mu.Unlock()
	fmt.Printf("DownloadTask第%d个并发，第%d个执行\n", param, count)
	count++
	return param + 100, nil
}

type NomalTask struct{}

func (t NomalTask) Run(param int) (int, error) {
	mu.Lock()
	defer gw.Done()
	defer mu.Unlock()
	fmt.Printf("NomalTask第%d个并发，第%d个执行\n", param, count)
	count++
	return param + 100, nil
}

func main() {
	var task []Task
	task = append(task, NomalTask{})
	task = append(task, NomalTask{})
	task = append(task, NomalTask{})
	task = append(task, DownloadTask{File: "File 1"})
	task = append(task, DownloadTask{File: "File 2"})
	task = append(task, DownloadTask{File: "File 3"})
	task = append(task, NomalTask{})
	task = append(task, NomalTask{})
	gw.Add(len(task))

	// for i := 0; i < len(task); i++ {
	// 	go task[i].Run(i)
	// }
	var results []string = make([]string, len(task))

	for i, t := range task {
		go func(j int, dt Task) {
			d, _ := dt.Run(j)
			f, ok := dt.(DownloadTask)
			// if reflect.TypeOf(dt).String() == "main.DownloadTask" {
			if ok {
				results[j] = strconv.Itoa(d) + "," + f.File
			} else {
				results[j] = strconv.Itoa(d)
			}
		}(i, t)
	}
	gw.Wait()

	for _, d := range results {
		fmt.Printf("result %s\n", d)
	}
}
