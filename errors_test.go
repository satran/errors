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
