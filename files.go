package testutils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TmpDir(tb testing.TB, prefix string) (path string, cleanup func(), ok bool) {
	tb.Helper()
	path, err := ioutil.TempDir("", prefix)
	if err != nil {
		tb.Error(err)
		return
	}
	return path, func() { os.RemoveAll(path) }, true
}

func FailIf(tb testing.TB, err error) (failed bool) {
	tb.Helper()
	if failed = err != nil; failed {
		tb.Error(err)
	}
	return
}
