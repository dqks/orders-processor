package services

func ProcessByWorkers[J any, R any](
	workerNum int,
	jobs []J,
	processJobCb func(<-chan *J, chan<- R, int),
	processResultCb func(R)) {

	if workerNum <= 0 {
		return
	}

	jobNum := len(jobs)
	jobChan := make(chan *J, jobNum)
	resultChan := make(chan R, jobNum)

	for i := 1; i <= workerNum; i++ {
		go processJobCb(jobChan, resultChan, i)
	}

	for i := 0; i < jobNum; i++ {
		jobChan <- &jobs[i]
	}

	close(jobChan)

	for i := 1; i <= jobNum; i++ {
		processResultCb(<-resultChan)
	}

	close(resultChan)
}
