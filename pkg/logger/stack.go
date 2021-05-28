package logger

import (
	"runtime"
	"strings"
)

func getCaller(callDepth int) (file string, line int) {
	callDepth += 2
	rpc := make([]uintptr, 10)
	n := runtime.Callers(callDepth, rpc[:])
	frames := runtime.CallersFrames(rpc[:n])
	containsToIgnore := []string{"github.com/sirupsen/logrus"}
	for {
		frame, ok := frames.Next()
		if !ok {
			return "???", 0
		}

		for _, s := range containsToIgnore {
			if strings.Contains(frame.File, s) {
				break
			}
			return frame.File, frame.Line
		}
	}
}
