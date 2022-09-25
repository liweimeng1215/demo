package args_go

import (
	"polymorphism/args"
	"testing"
)

func Test_ParseAllSchema(t *testing.T) {
	args := args.New()
	args.Parse("l,d*", []string{"-l", "-d", "/a/b/c"})
	if !args.GetBool("l") {
		t.Fatal("boolean parse failed")
	}
	if args.GetString("d") != "/a/b/c" {
		t.Fatal("string parse failed")
	}
}
