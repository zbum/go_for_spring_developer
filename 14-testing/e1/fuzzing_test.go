package e1

import "testing"

func FuzzFoo(f *testing.F) {
	f.Add(5, "hello1")
	f.Add(6, "hello2")

	f.Fuzz(func(t *testing.T, i int, s string) {
		out, err := Foo(i, s)
		if err != nil && out != "" {
			t.Errorf("%q %v", out, err)
		}
	})
}
