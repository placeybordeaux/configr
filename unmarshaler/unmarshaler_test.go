package configr

import (
	"fmt"
	"os"
	"testing"
)

func TestUnmarshalFromEnv(t *testing.T) {
	type C struct {
		A int64   `cli:"a" desc:"a"`
		B float64 `cli:"b" desc:"b"`
		C bool    `cli:"c" desc:"c"`
		D string  `cli:"d" desc:"d"`
	}
	os.Args = []string{"cmd", "-a", "123", "-b", "1.1", "-d", "true", "-c", "true", "-last"}
	c := C{}
	UnmarshalFromFlags(&c)
	fmt.Println(c)
	if c.A != 123 {
		t.Fatal("int64 failed to set")
	}
	if c.B != 1.1 {
		t.Fatal("float64 failed to set")
	}
	if c.C != true {
		t.Fatal("bool failed to set")
	}
	if c.D != "true" {
		t.Fatal("string failed to set")
	}
}
