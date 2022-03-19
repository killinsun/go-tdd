package dependencyInjection

import (
	"bytes"
	"testing"
)

// バッファに書き込んで、書き込まれた値をテストする
func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}