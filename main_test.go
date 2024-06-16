/**
 * Package test
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/16
 */

package main

import "testing"

func TestFakeFunc(t *testing.T) {
	want := "fake func"
	if got := FakeFunc(); got != want {
		t.Errorf("FakeFunc() = %q, want %q", got, want)
	}
}
