# golang-race

### For tag: 0.0.1

Execute this command - order to view of summation result
> `$ go run main.go`

Example output:
`261848`

Expected: `1000000`


Execute this command - order to view of "race" detection
> `$ go run -race main.go`

```
$ go run -race main.go
==================
WARNING: DATA RACE
Read at 0x00c04206e080 by goroutine 7:
  main.main.func1()
      D:/proj/golang/golang-race/main.go:11 +0x48

Previous write at 0x00c04206e080 by goroutine 6:
  main.main.func1()
      D:/proj/golang/golang-race/main.go:11 +0x5e

Goroutine 7 (running) created at:
  main.main()
      D:/proj/golang/golang-race/main.go:19 +0x8f

Goroutine 6 (running) created at:
  main.main()
      D:/proj/golang/golang-race/main.go:19 +0x8f
==================
210100
Found 1 data race(s)
exit status 66

```

### For tag: 0.0.2
Default run
>`$ go run main_atomic.go`

>`1000000`


With flag "-race"
>`$ go run -race main_atomic.go`

>`963237`

### For tag: 0.0.3

Default run
>`$ go run main_mutex.go`

> `1000000`

With flag "-race"
> `$ go run -race main_mutex.go`

> `610917`


### For tag: 0.0.4

>`$ go run main_waitgroup.go`

> `1000000`

With flag "-race"
> `$ go run -race main_waitgroup.go`

> `1000000`
