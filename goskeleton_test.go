package goskeleton

import "testing"

var testcases = []struct {
	name string
	want string
}{
	{
		name: "Bob",
		want: "Hello, Bob",
	},
	{
		name: "Alice",
		want: "Hello, Alice",
	},
}

func dotest(t *testing.T, name, want string) {
	foo := Foo{Name: name}
	greet := foo.Greet()
	if greet != want {
		t.Errorf("name = %v; got %v; want %v", name, greet, want)
	}
}

func TestGreet(t *testing.T) {
	for _, tc := range testcases {
		dotest(t, tc.name, tc.want)
	}
}

func TestGreetParallel(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				dotest(t, tc.name, tc.want)
			})
		}
	})
}

func BenchmarkGreet(b *testing.B) {
	foo := Foo{Name: "Bob"}
	for i := 0; i < b.N; i++ {
		foo.Greet()
	}
}
