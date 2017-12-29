package errors

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("test")
	customErr := New("my custom error").CausedBy(err)
	if customErr.caused != err {
		t.Error("expected caused by error")
	}
	if customErr.line != 10 {
		t.Errorf("unexpected line; %#v", customErr)
	}

	causedBy := WithCaused(errors.New("blah"), "test")
	if causedBy.line != 18 {
		t.Errorf("unexpedted line: %#v", causedBy)
	}
}

func TestDepths(t *testing.T) {
	err := depth1().(*Error)
	if exp, got := 36, err.line; exp != got {
		t.Errorf("expected: %d got: %d", exp, got)
	}
}

func depth1() error {
	return depth2()
}

func depth2() error {
	return New("test depth")
}
