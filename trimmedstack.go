package gameboy

import (
	"fmt"
	"regexp"
	"runtime"
)

func TrimmedStack(err any, pkgRegex string) {
	re := regexp.MustCompile(pkgRegex)

	pcs := make([]uintptr, 32)
	n := runtime.Callers(3, pcs) // skip runtime.Callers + this func + defer
	frames := runtime.CallersFrames(pcs[:n])

	var lastMatchedFrame *runtime.Frame
	for {
		frame, more := frames.Next()
		if re.MatchString(frame.Function) {
			lastMatchedFrame = &frame
		} else {
			break
		}
		if !more {
			break
		}
	}

	if lastMatchedFrame != nil {
		fmt.Printf("panic: %v\nat %s\n\t%s:%d\n", err, lastMatchedFrame.Function, lastMatchedFrame.File, lastMatchedFrame.Line)
	} else {
		fmt.Printf("panic: %v\n(No matching frames)\n", err)
	}
}
