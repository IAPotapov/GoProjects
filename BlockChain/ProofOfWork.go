package main

type Job struct {
	Index         int
	Input         string
	Seed          uint64
	Complexity    int
	WorkFlag      *bool
	ResultChannel chan int
	Result        string
	HashString    string
	*Sha256Data
	*StringGenerator
}

func (job *Job) DoJob() {
	s := []byte(job.Input)
	job.original = make([]byte, len(s)+len(job.StringValue))
	copy(job.original, s)
	job.Start()
}

func (job *Job) Start() {

	var i uint64 = job.Seed
	for *(job.WorkFlag) {
		job.Update()

		job.CalculateSha256()
		if job.IsCorrect() {
			job.FormatResult()
			job.ResultChannel <- job.Index
			break
		}
		i++
	}
}

func (job *Job) Update() {
	job.Next()
	j := len(job.original) - 1

	for i := 0; i < 11; i++ {
		job.original[j-i] = job.StringValue[i]
	}
}

func (job *Job) IsCorrect() bool {
	for i := 0; i < job.Complexity; i++ {
		if job.hash[i] != 0 {
			return false
		}
	}
	return true
}

func (job *Job) FormatResult() {
	job.Result = job.Input + job.GetString()
	job.HashString = job.GetHashString()
}
