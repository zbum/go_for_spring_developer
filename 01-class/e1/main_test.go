package main

import "testing"

func TestGreeter_Greet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK Manty",
			args: args{name: "Manty"},
			want: "Hello, Manty!!",
		},
		{
			name: "OK Benjamin",
			args: args{name: "Benjamin"},
			want: "Hello, Benjamin!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Greeter{}
			if got := g.Greet(tt.args.name); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
