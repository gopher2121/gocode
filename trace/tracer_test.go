package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("tracer value should not be nil")
	} else {
		tracer.Trace("Hello")
		if buf.String() != "Hello/n" {
			t.Errorf("trace should not write '%s'.", buf.String())
		}
	}
}
