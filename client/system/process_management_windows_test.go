//go:build windows

package system

import (
	"testing"

	_ "alda.io/client/testing"
)

func TestSysProcAttr(t *testing.T) {
	attr := sysProcAttr()
	if attr != nil {
		t.Errorf("sysProcAttr() should return nil on Windows, got: %+v", attr)
	}
}
