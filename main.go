package main

import (
	"bytes"
	"github.com/Analyse4/stride-scheduler/job"
	ufm "github.com/Analyse4/stride-scheduler/unfairness-metric"
	"github.com/wcharczuk/go-chart"
	"io/ioutil"
	"log"
)

const BIGNUM = 1000

func main() {
	metricList := make([]*ufm.Metric, 0)
	for jobLength := 1; jobLength <= 1000; jobLength++ {
		jobList := job.CreateTwoJob(jobLength, BIGNUM)
		m := ufm.New()
		totalWorkload := 2 * jobLength
		clock := 0.0
		m.Total = float64(totalWorkload)
		for {
			if jobList[0].Pass <= jobList[1].Pass {
				jobList[0].Work()
			} else {
				jobList[1].Work()
			}
			clock += 1
			if jobList[0].IsFinish() || jobList[1].IsFinish() {
				m.First = clock
				metricList = append(metricList, m)
				break
			}
		}
	}

	xRange, yRange := ufm.ConvertToXYRange(metricList)
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xRange,
				YValues: yRange,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./chart.png", buffer.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
