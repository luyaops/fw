package log

import "testing"

func TestInfof(t *testing.T) {
	Infof("test infof")
	Errorf("test infof")
	Debugf("test debugf")
}
