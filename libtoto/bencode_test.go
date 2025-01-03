package libtoto_test

import (
	"slices"
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

	_, err = libtoto.ParseBencode("inane")
	if err == nil {
		t.Errorf("invalid int not erroring")
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

	_, err = libtoto.ParseBencode("5aaaa:")
	if err == nil {
		t.Errorf("bad length not erroring")
	}
}

func TestList(t *testing.T) {
	result, err := libtoto.ParseBencode("l1:ae")
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != libtoto.List {
		t.Errorf("failed to parse list")
	}
	if !slices.Equal(result.Val.([]any), []any{"a"}) {
		t.Errorf("failed to parse list")
	}

	result, err = libtoto.ParseBencode("l1:a2:bbe")
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != libtoto.List {
		t.Errorf("failed to parse multi-element list")
	}
	if !slices.Equal(result.Val.([]any), []any{"a", "bb"}) {
		t.Errorf("failed to parse multi-element list")
	}

	result, err = libtoto.ParseBencode("le")
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != libtoto.List {
		t.Errorf("failed to parse empty list")
	}
	if !slices.Equal(result.Val.([]any), []any{}) {
		t.Errorf("failed to parse empty list")
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

	_, err = libtoto.ParseBencode("i1")
	if err != libtoto.MissingDelimiter {
		t.Errorf("missing 'e' token for an int not erroring")
	}
}
