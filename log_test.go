package ntnn

import "testing"

func TestLastLog(t *testing.T) {
	if !lastLogChanged("first", "first msg") {
		t.Errorf("first message should report changed")
	}
	lastLogUpdate("first", "first msg")
	if lastLogChanged("first", "first msg") {
		t.Errorf("first message after update should not report changed")
	}
	if lastLogChanged("first", "first msg") {
		t.Errorf("first message second time after update should not report changed")
	}

	if !lastLogChanged("second", "second msg") {
		t.Errorf("second message should report changed")
	}
}

// separate function so the origin is identical
func logger(msg string) {
	LogChanged(msg)
}

func ExampleLogChanged() {
	// only prints one "first msg"
	logger("first msg")
	logger("first msg")
	logger("first msg")
	// prints "second msg"
	logger("second msg")
	// prints "first msg" again
	logger("first msg")
	// Output:
	// ###> first msg
	// ###> second msg
	// ###> first msg
}
