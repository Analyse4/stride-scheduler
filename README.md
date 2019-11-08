# stride-scheduler
It's a program simulating stride scheduling for generating a chart. The chart describe the relationship about job length and unfairness metric(which is simply the time the first job completes divided by the time that the second job completes).
## Condition

- two jobs
- tickets of both: 100
- big number: 1000(stride: 10)
- job length from 1 to 1000
- each time slice can do a unit of work
## Chart
![chart](https://github.com/Analyse4/stride-scheduler/blob/master/chart.png)