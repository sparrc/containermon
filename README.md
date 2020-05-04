# containermon

Install with go:
```
go get github.com/sparrc/containermon
```

Container Name or ID is the only required argument
```
% ./containermon --container jovial_euler
{"timeElapsed":11,"cpuTimeElapsed":0,"percentCPUSinceStart":0.00,"percentCPUThisInterval":0.00}
{"timeElapsed":21,"cpuTimeElapsed":0,"percentCPUSinceStart":0.00,"percentCPUThisInterval":0.00}
{"timeElapsed":31,"cpuTimeElapsed":3,"percentCPUSinceStart":12.16,"percentCPUThisInterval":38.21}
{"timeElapsed":41,"cpuTimeElapsed":13,"percentCPUSinceStart":33.35,"percentCPUThisInterval":99.93}
{"timeElapsed":51,"cpuTimeElapsed":15,"percentCPUSinceStart":30.22,"percentCPUThisInterval":17.29}
{"timeElapsed":61,"cpuTimeElapsed":15,"percentCPUSinceStart":25.30,"percentCPUThisInterval":0.00}
```

You can also output to csv and change the collection interval
```
% ./containermon --container jovial_euler --interval 5 --output-format csv 
timeElapsed,cpuTimeElapsed,percentCPUSinceStart,percentCPUThisInterval
6,0,0.00,0.00
11,0,0.00,0.00
16,3,20.90,69.52
21,8,39.11,99.75
26,10,40.72,47.68
31,10,34.29,0.00
```
