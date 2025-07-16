package ntnn

import (
	"context"
	"fmt"
	"runtime/pprof"
)

// WithLabels is a convenience wrapper around pprof.Do, converting the
// label pairs to string, adding them to the context and executing the
// function.
//
// If context is nil it defaults to context.Background().
//
// The label pairs should be key-value pairs:
//
//	WithLabels(ctx, fn, "key1", "value1", "key2", "value2")
func WithLabels(
	ctx context.Context,
	fn func(context.Context),
	labelPairs ...any,
) {
	if ctx == nil {
		ctx = context.Background()
	}

	strLabelPairs := make([]string, len(labelPairs))
	for i := range labelPairs {
		strLabelPairs[i] = fmt.Sprintf("%v", labelPairs[i])
	}

	pprof.Do(
		ctx,
		pprof.Labels(strLabelPairs...),
		fn,
	)
}
