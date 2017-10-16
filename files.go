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
