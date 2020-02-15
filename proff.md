# Pro-ff

A fake web/software console for executing commands

## Requirements

### Core

- execute commands
  - linux
  - window
  - macOs
- log files && display
- cron scheduler
- sequential execution on commands that will be executed
  - `graphical arrangement`

### Others

- donate
- security
  - password for remote
- configure
  - port
  - username: password

### More

- connect or transfer files between two client

## Design

### design core

#### D1

- load config
  - port
  - password
- execute commands at backend
  - name + number
  - upload insert line
  - upload button
- record log for each action
  - done or err

#### D2

- cron scheduler
  - start
  - stop
- D1

#### D3

- arrange commands
- D1

### http support

- http
  - port: 6121

## Libraries

- github.com/robfig/cron

## Import Structure

## Classes

## Functions

## More(Notes)

## Pseudo-code
