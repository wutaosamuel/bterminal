# BTerminal

English | [中文](./README_CN.md)

A simple tool executes command with or without schedule.
It is self-hosted web service without dependencies.
The project primarily uses golang,
which provides good flexibility as result of cross-platform property(e.g. Windows).

## Installation

``` sh
go get -u github.com/wutaosamuel/bterminal/cmd/bterminal
```

or

``` sh
# debian/ubuntu
dpkg -i bterminal-1.0-1.deb
```

or

``` sh
# windows, tested win10 only
# use after setup
bterminal-1.0-1.msi
```

or

``` sh
# debian/utuntu
dpkg -i bterminal-1.0-1.deb
```


## Features

- Password protection
- Execute single command on web browser, running on background service.
- Cron Scheduling on all operating system, including Windows.
  - [cron wiki](https://en.wikipedia.org/wiki/Cron)
- Cookie sessions and single tokens are adopted to improve http security.

## Getting Started

### Run

#### Simple run

``` sh
# Linux
bterminal
```

``` sh
# Windows
1. bterminal.exe
2. click start button of bterminal on the right-bottom System tray
```

Web port, password and logs path can be changed by command. More details: bterminal -h | help.
if program installed by windows installer, pls, change C:/ProgramData/bterminal/config.json

### Usage

``` sh
# manage jobs on web browser
127.0.0.1:5122
```

## Figures

Figure 1: Password index page

Password protection can be setted by config or cli command
and is Base64 encoded

![BTerminalPassword](./image/bterminalPassword.png)

Figure 2: UI of Single Command Entering

The command is required,
and whereas the name and cron scheduling are optional.
The command will execute immediately when cron scheduling is empty.
Strongly recommend
that shell script or python handle multiple commands input.

The main motivation of this project is
the international output bandwidth of Chinese Telecom has been QoS by home using with some reason.
Therefore, it is beneficial to exploit such an unused bandwidth,
and schedule certain tasks to less-QoS periods, typically around 4am.  

![BTerminalShell](./image/bterminalShell.png)

Figure 3: display cron tasks

![BTerminalJob](./image/bterminalJobs.png)

Figure 4: display all tasks' log

![BTerminalLogs](./image/bterminalLogs.png)

## TODO

- [ ] -c, --clean: clean GobData.dat, html/logs.html, html/jobs.html
- [ ] restart a stopped job
- [ ] test wrong format cron
- [ ] test on Windows
- [ ] startup on Windows
- [ ] test on MacOs
- [ ] a log for watching the whole program
- [ ] limit log length on web interface
- [x] make windows installer
- [x] deb package
- [ ] support go mod (currently, fail on uuid)
- [ ] develop new web interface or a software UI

## Libraries

- github.com/patrickmn/go-cache
- github.com/satori/go.uuid
- github.com/robfig/cron

## Contributing

If you are interested, you are welcome to contribute to this project!

- if you encounter a bug, issue it first.
- if you have an idea or problem, feel free to post it on issue.
- if you can contribute code,
pls sending Pull Request to `dev branch`.
