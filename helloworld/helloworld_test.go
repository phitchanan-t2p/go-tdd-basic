package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("test function HelloWorld should be return Hello World with message", func(t *testing.T) {
		got := HelloWorld("message")
		want := "Hello World message"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
