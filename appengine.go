// +build appengine

package logw

import (
	"context"

	"google.golang.org/appengine/log"
)

func Debugf(ctx context.Context, format string, args ...interface{}) {
	log.Debugf(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	log.Infof(ctx, format, args...)
}

func Warningf(ctx context.Context, format string, args ...interface{}) {
	log.Warningf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(ctx, format, args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	log.Criticalf(ctx, format, args...)
}

func FatalDepthf(ctx context.Context, depth int, format string, args ...interface{}) {
	log.Criticalf(ctx, format, args...)
}

func Exitf(ctx context.Context, format string, args ...interface{}) {
	log.Criticalf(ctx, format, args...)
}
