package filesystem

import (
	"context"
	"time"
)

func (fs *ObjStorFs) getOperationContext(opName, opClassName string) (context.Context,
	context.CancelFunc) {
	runtime := fs.opts.deadlines[opClassName]
	opDuration, exists := fs.opts.deadlines[opName]
	if exists {
		runtime = opDuration
	}

	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(runtime))
	return ctx, cancelFunc
}
