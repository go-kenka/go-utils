package logger

import (
	"fmt"
	"runtime/debug"
	"time"
)

// ExceptionDump start watching stack trace
func ExceptionDump() {
	if e := recover(); e != nil {
		DumpStack(e, debug.Stack())
		panic(e)
	}
}

// DumpStack dump stack trace
func DumpStack(e interface{}, stack []byte) {

	// if err := slack.Send(e, string(stack)); err != nil {
	// 	fmt.Printf("cant send to slack %v\n", err)
	// }
	fmt.Printf("[%v] crashing...\n", PrettyPrintTime(time.Now()))
}

// PrettyPrintTime pretty print a time value to readable
func PrettyPrintTime(timeVal time.Time) string {
	return timeVal.Format("Mon Jan _2 15:04:05 2006")
}
