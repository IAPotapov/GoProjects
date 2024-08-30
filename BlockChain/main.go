package main

import (
	"fmt"
	"math"

	//"sync"
	"time"
)

func Calculate(config *AppConfig) {
	t1 := time.Now()
	fmt.Println("Start time:", t1.Format(time.DateTime))

	jobs := make([]Job, config.JobNumber)
	var delta uint64 = math.MaxUint64
	delta /= uint64(config.JobNumber)
	var WorkFlag bool = true
	var ResultChannel chan int = make(chan int)
	for i := 0; i < config.JobNumber; i++ {
		var seed uint64 = delta * uint64(i)
		var sha Sha256Data
		sha.Initialize()
		var sgen StringGenerator
		sgen.Initialize(seed)

		jobs[i] = Job{
			Index:           i,
			Input:           config.Text,
			Complexity:      config.Complexity,
			Seed:            seed,
			WorkFlag:        &WorkFlag,
			ResultChannel:   ResultChannel,
			Sha256Data:      &sha,
			StringGenerator: &sgen,
		}
		go jobs[i].DoJob()
	}

	// wait any result
	j := <-ResultChannel
	WorkFlag = false

	fmt.Println(jobs[j].Result)
	fmt.Println(jobs[j].HashString)

	t2 := time.Now()
	fmt.Println("Finish time:", t2.Format(time.DateTime))
	d := t2.Sub(t1)

	s := d.Seconds()
	m := s / 60
	si := int(s) % 60
	h := m / 60
	mi := int(m) % 60
	hi := int(h)
	fmt.Printf("Duration: %02d:%02d:%02d\n", hi, mi, si)
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
