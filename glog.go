// +build !appengine

// Package logw provides a build-time leveled log wrapper over glog and intrinsic
// cloud log packages. Each environment must provide Debugf, Infof, Warningf,
// Errorf, Fatalf and Exitf functions and a way to enable/disable debug logging.
package logw

import (
	"context"
	"fmt"

	"github.com/golang/glog"
)

func init() {
	glog.CopyStandardLogTo("INFO")
}

// Debugf formats and logs a debug message if debug logging is enabled.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	if glog.V(1) {
		glog.InfoDepth(1, fmt.Sprintf(format, args...))
	}
}

// Infof formats and logs an info message.
func Infof(ctx context.Context, format string, args ...interface{}) {
	glog.InfoDepth(1, fmt.Sprintf(format, args...))
}

// Warningf formats and logs a warning message.
func Warningf(ctx context.Context, format string, args ...interface{}) {
	glog.WarningDepth(1, fmt.Sprintf(format, args...))
}

// Errorf formats and logs an error message.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	glog.ErrorDepth(1, fmt.Sprintf(format, args...))
}

// Fatalf formats and logs an error message and halts execution is some fashion.
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	glog.FatalDepth(1, fmt.Sprintf(format, args...))
}

// FatalDepthf formats and logs an error message -- using the source information
// from 'depth' up the call stack -- and halts execution is some fashion.
func FatalDepthf(ctx context.Context, depth int, format string, args ...interface{}) {
	glog.FatalDepth(depth+1, fmt.Sprintf(format, args...))
}

// Exitf formats and logs an error message and halts execution is some fashion.
// It differs from Fatalf in that the exit condition is not considered abnormal.
func Exitf(ctx context.Context, format string, args ...interface{}) {
	glog.ExitDepth(1, fmt.Sprintf(format, args...))
}
