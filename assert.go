package logw

import (
	"context"
	"fmt"
)

// Assert asserts that the predicate is true and logw.Fatals/panics otherwise.
func Assert(ctx context.Context, predicate bool, format string, args ...interface{}) {
	if !predicate {
		fail(ctx, format, args...)
	}
}

// AssertNil asserts that the err is nil and logw.Fatals/panics otherwise. The error is
// implicitly appended to the format as ": %v" and passed as a trailing arg.
func AssertNil(ctx context.Context, err error, format string, args ...interface{}) {
	if err != nil {
		args = append(args, err)
		fail(ctx, format+": %v", args...)
	}
}

func fail(ctx context.Context, format string, args ...interface{}) {
	FatalDepthf(ctx, 2, format, args...)

	// Backup panic, in case the logw.Fatal doesn't.
	panic(fmt.Sprintf(format, args...))
}
