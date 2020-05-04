# containermon

Install with go:
```
go get -u github.com/sparrc/containermon
```

Container Name or ID is the only required argument
```
$ containermon --container jovial_euler
{"ts":"2020-05-04T18:00:51Z","timeElapsed":13.19,"cpuTimeElapsed":0.67,"percentCPUSinceStart":5.07,"percentCPUThisInterval":5.07}
{"ts":"2020-05-04T18:01:00Z","timeElapsed":22.88,"cpuTimeElapsed":0.91,"percentCPUSinceStart":3.97,"percentCPUThisInterval":2.48}
{"ts":"2020-05-04T18:01:10Z","timeElapsed":32.24,"cpuTimeElapsed":1.18,"percentCPUSinceStart":3.67,"percentCPUThisInterval":2.93}
{"ts":"2020-05-04T18:01:21Z","timeElapsed":43.32,"cpuTimeElapsed":1.76,"percentCPUSinceStart":4.06,"percentCPUThisInterval":5.19}
{"ts":"2020-05-04T18:01:30Z","timeElapsed":53.02,"cpuTimeElapsed":1.99,"percentCPUSinceStart":3.75,"percentCPUThisInterval":2.40}
{"ts":"2020-05-04T18:01:40Z","timeElapsed":62.72,"cpuTimeElapsed":2.30,"percentCPUSinceStart":3.66,"percentCPUThisInterval":3.16}
{"ts":"2020-05-04T18:01:50Z","timeElapsed":72.25,"cpuTimeElapsed":2.83,"percentCPUSinceStart":3.92,"percentCPUThisInterval":5.59}
{"ts":"2020-05-04T18:02:01Z","timeElapsed":83.30,"cpuTimeElapsed":3.08,"percentCPUSinceStart":3.70,"percentCPUThisInterval":2.28}
{"ts":"2020-05-04T18:02:10Z","timeElapsed":92.43,"cpuTimeElapsed":3.33,"percentCPUSinceStart":3.61,"percentCPUThisInterval":2.76}
```

You can also output to csv and change the collection interval
```
$ containermon --container jovial_euler --interval 5 --output-format csv
ts,timeElapsed,cpuTimeElapsed,percentCPUSinceStart,percentCPUThisInterval
2020-05-04T18:21:41Z,7.80,0.26,3.39,3.39
2020-05-04T18:21:46Z,12.75,0.68,5.31,8.34
2020-05-04T18:21:50Z,17.20,0.80,4.68,2.85
2020-05-04T18:21:56Z,23.27,0.93,4.00,2.08
2020-05-04T18:22:02Z,28.40,1.08,3.79,2.86
```
