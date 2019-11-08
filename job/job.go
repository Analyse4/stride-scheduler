package job

import "log"

type Job struct {
	Len        int
	Pass       int
	CurrentLen int
	Stride     int
	Ticket     int
}

func CreateTwoJob(Len int, BigNum int) []*Job {
	jobList := make([]*Job, 0)
	for i := 0; i < 2; i++ {
		job := Job{
			Len:        Len,
			Pass:       0,
			CurrentLen: Len,
			Stride:     0,
			Ticket:     100,
		}
		job.Stride = BigNum / job.Ticket

		jobList = append(jobList, &job)
	}
	return jobList
}

func (j *Job) Work() {
	j.Pass += j.Stride
	j.CurrentLen -= 1
}

func (j *Job) IsFinish() bool {
	if j.CurrentLen < 0 {
		log.Fatal("invalid value of CurrentLen")
	}
	if j.CurrentLen == 0 {
		return true
	}
	return false
}
