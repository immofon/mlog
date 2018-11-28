package mlog

import "testing"

func Test_L(t *testing.T) {
	L().Info("hello")
	L().Info("hello")
	L().Info("hello")
	L().Info("hello")
	L().Info("hello")
}
