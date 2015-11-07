package trace

import (
	"fmt"
	"io"
)

// Tracer is an interface for tracing the code
type Tracer interface {
	// Trace function will take zero or more arguments of any type
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("/n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
