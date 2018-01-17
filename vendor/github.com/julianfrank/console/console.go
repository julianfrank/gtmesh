// Package console tries to bring the convenience of console.log/console.error in JS/NodeJS to Go
package console

import (
	"fmt"
	"log"
	"time"
)

var (
	//LogMode Use this to turn Log output to console off or on
	LogMode = true
	//ErrorMode Use this to turn Error output to console off or on
	ErrorMode = true
	//TimedMode Use this to turn Timed Output to console off or on
	TimedMode = true
)

//Log use as console.log(pattern,vars...) return a string
func Log(pattern string, vars ...interface{}) string {
	var msg = fmt.Sprintf(pattern, vars...)
	if LogMode {
		log.Println(msg)
	}
	return msg
}

//Error return an error object as per the provided format
// Usage console.Error(pattern,vars...)
// Output: This will Return an error object
func Error(pattern string, vars ...interface{}) error {
	var err = fmt.Errorf(pattern, vars...)
	if ErrorMode {
		log.Println(err.Error())
	}
	return err
}

//Timed returns a string object appended with the time elapsed provided from the starttime
// Usage: console.Timed(starttime,pattern,vars...)
// Output: returns the formatted string appended with elapsed time and outputs the string to stdio if TimedMode is true
func Timed(startTime time.Time, pattern string, vars ...interface{}) string {
	pattern += "\t%s"
	vars = append(vars, time.Since(startTime))
	var msg = fmt.Sprintf(pattern, vars...)
	if TimedMode {
		fmt.Println(msg)
	}
	return msg
}
