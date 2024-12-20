package libtoto_test

import (
	"testing"

	"github.com/demostanis/hypertube/libtoto"
)

func TestInt(t *testing.T) {
	result, err := libtoto.ParseBencode("i123e")
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != libtoto.Int {
		t.Errorf("failed to parse int")
	}
	if result.Val != 123 {
		t.Errorf("failed to parse int")
	}
}

func TestStr(t *testing.T) {
	result, err := libtoto.ParseBencode("5:hello")
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != libtoto.Str {
		t.Errorf("failed to parse string")
	}
	if result.Val != "hello" {
		t.Errorf("failed to parse string")
	}

	_, err = libtoto.ParseBencode("5:h")
	if err != libtoto.BadLength {
		t.Errorf("bad length not erroring")
	}
}

func TestMiscErrors(t *testing.T) {
	_, err := libtoto.ParseBencode("")
	if err != libtoto.EOF {
		t.Errorf("empty input not erroring")
	}

	_, err = libtoto.ParseBencode("o")
	if err != libtoto.UnknownType {
		t.Errorf("unknown type not erroring")
	}

	_, err = libtoto.ParseBencode("18")
	if err != libtoto.MissingDelimiter {
		t.Errorf("missing ':' token for a string not erroring")
	}
}
