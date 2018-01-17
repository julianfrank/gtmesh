[![GoDoc](https://godoc.org/github.com/julianfrank/console?status.svg)](https://godoc.org/github.com/julianfrank/console)
 [![Coverage Status](https://coveralls.io/repos/github/julianfrank/console/badge.svg?branch=master)](https://coveralls.io/github/julianfrank/console?branch=master) [![Build Status](https://travis-ci.org/julianfrank/console.svg?branch=master)](https://travis-ci.org/julianfrank/console)
 [![Go Report Card](https://goreportcard.com/badge/github.com/julianfrank/console)](https://goreportcard.com/report/github.com/julianfrank/console)
[![Gitter](https://img.shields.io/badge/gitter-join-brightgreen.svg)](https://gitter.im/jfopensource/console)


# console package for those used to JS console.log(...)
Simple Console Log Package useful for anyone coming from Javascript background

## Usage
Just import this package
```go
import "github.com/julianfrank/console"
```

### Log
In your Application use the familiar console.Log ...Just remember that the L is capital
```go
comment:="Awsome!"
dude:="Julian Frank"
console.Log("This comment (%s) was made by %s",comment,dude)

mylog:=console.Log("This comment (%s) was made by %s",comment,dude)
//Output mylog=> "This comment (Awsome!) was made by Julian Frank"
```

### Error
In your Application use the familiar console.Error ...Just remember that the E is capital
```go
comment:="Awsome!"
dude:="Julian Frank"
myError:=console.Error("This comment (%s) was made by %s")
//Output myError=> "This comment (Awsome!) was made by Julian Frank" 
//of type error
//Use err.Error() to retreive error string 
```


### Timed
This tries to replicate what the console timestart/timeend does in JS... it is slightly different... here you will need to record the start time in a variable at the point where you want to start measuring... in the point were you need the measure use 
console.Timed(<the timedstart variable>,<the pattern>,the variables)
The elabpsed time will be automatically appended to the pattern with a tab
```go
ts:=time.Now()
console.TimedMode=true
comment:="Awsome!"
dude:="Julian Frank"
console.Timed(ts,"This comment (%s) was made by %s",comment,dude)

mylog:=console.Log(ts,"This comment (%s) was made by %s",comment,dude)
//Output mylog=> "This comment (Awsome!) was made by Julian Frank    0s"
```


### Console Display Flags 
If You want to selectively enable or disable console display of Log or Error you can simply do that by setting console.LogMode and console.ErrorMode to true or false. True is default
```go
//To make screen completely quiet
console.LogMode=false
console.ErrorMode=false
console.TimedMode=false
//To Display only errors and not the Logs
console.LogMode=false
console.ErrorMode=true
console.TimedMode=false
```

## License
Apache 2.0
