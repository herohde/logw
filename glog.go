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

func Debugf(ctx context.Context, format string, args ...interface{}) {
	if glog.V(1) {
		glog.InfoDepth(1, fmt.Sprintf(format, args...))
	}
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	glog.InfoDepth(1, fmt.Sprintf(format, args...))
}

func Warningf(ctx context.Context, format string, args ...interface{}) {
	glog.WarningDepth(1, fmt.Sprintf(format, args...))
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	glog.ErrorDepth(1, fmt.Sprintf(format, args...))
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	glog.FatalDepth(1, fmt.Sprintf(format, args...))
}

func FatalDepthf(ctx context.Context, depth int, format string, args ...interface{}) {
	glog.FatalDepth(depth+1, fmt.Sprintf(format, args...))
}

func Exitf(ctx context.Context, format string, args ...interface{}) {
	glog.ExitDepth(1, fmt.Sprintf(format, args...))
}
