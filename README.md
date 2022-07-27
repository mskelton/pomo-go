# Pomo

Simple pomodoro timer CLI.

_Inspired by [rwxrob/pomo](https://github.com/rwxrob/pomo), but written from scratch because why not!_

## Installation

```bash
go install github.com/mskelton/pomo@latest
```

## Usage

### Get status

Prints the status of the current session.

```bash
pomo
```

### Start focus

Starts a new pomodoro focus session with the default duration.

```bash
pomo start
```

Or customize the session duration by providing a Go [`Duration`](https://pkg.go.dev/time#Duration).

```bash
pomo start 15m
```

### Start break

Starts a new break with the default duration.

```bash
pomo break
```

Or customize the break duration by providing a Go [`Duration`](https://pkg.go.dev/time#Duration).

```bash
pomo break 10m
```

### Change duration

Changes the duration of the active session using the specified Go [`Duration`](https://pkg.go.dev/time#Duration).

```bash
pomo duration 20m
```

### Stop session

Stops the current pomodoro session.

```bash
pomo stop
```
