package jyutil

import (
	"fmt"
	"time"
)

const strProjectName string = "MyProject"

// ==============================================
// Output
// ==============================================
func Output(va ...any) {
	t := time.Now()
	str := fmt.Sprint(va...)
	fmt.Printf("%02d:%02d:%02d.%03d [ %s ] %s\n", t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000000, strProjectName, str)
}

// ==============================================
// OutputDebugString
// ==============================================
func OutputDebugString(strFormat string, va ...any) {
	t := time.Now()
	str := fmt.Sprintf(strFormat, va...)
	fmt.Printf("%02d:%02d:%02d.%03d [ %s ] %s", t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000000, strProjectName, str)
}
