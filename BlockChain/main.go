package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func Calculate(config *AppConfig) {
	t1 := time.Now()
	fmt.Println("Start time:", t1.Format(time.DateTime))

	jobs := make([]Job, config.JobNumber)
	var delta uint64 = math.MaxUint64
	delta /= uint64(config.JobNumber)
	var WorkFlag bool = true
	var wg sync.WaitGroup
	wg.Add(config.JobNumber)

	for i := 0; i < config.JobNumber; i++ {
		jobs[i] = Job{
			Index:           i,
			Input:           config.Text,
			Complexity:      config.Complexity,
			wg:              &wg,
			WorkFlag:        &WorkFlag,
			Sha256Data:      NewSha256Data(),
			StringGenerator: NewStringGenerator(delta * uint64(i)),
		}
		go jobs[i].DoJob()
	}

	// wait any result
	wg.Wait()

	for i := 0; i < config.JobNumber; i++ {
		if jobs[i].SuccessFlag {
			fmt.Println(jobs[i].Result)
			fmt.Println(jobs[i].HashString)
			break
		}
	}

	t2 := time.Now()
	fmt.Println("Finish time:", t2.Format(time.DateTime))
	d := t2.Sub(t1)

	s := d.Seconds()
	fmt.Printf("Duration: %v sec\n", s)
}

func main() {

	var config AppConfig
	err := config.Load()
	if err == nil {
		Calculate(&config)
	} else {
		fmt.Println(err.Error())
	}
}
