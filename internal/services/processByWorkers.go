package services

import "fmt"

func ProcessByWorkers[J any, R any](
	workerNum int,
	jobs []J,
	processCb func(<-chan *J, chan<- R, int)) {

	jobNum := len(jobs)
	jobChan := make(chan *J, jobNum)
	resultChan := make(chan R, jobNum)

	for i := 1; i <= workerNum; i++ {
		go processCb(jobChan, resultChan, i)
	}

	for i := 0; i < jobNum; i++ {
		jobChan <- &jobs[i]
	}

	close(jobChan)

	for i := 1; i <= jobNum; i++ {
		fmt.Println(<-resultChan, " Результат")
	}
}
