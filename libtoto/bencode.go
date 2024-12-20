package libtoto

import (
	"errors"
	"strconv"
	"strings"
)

var (
	EOF              = errors.New("EOF")
	MissingDelimiter = errors.New("missing delimiter")
	UnknownType      = errors.New("unknown type")
	BadLength        = errors.New("bad length")
)

type BencodeType = int

const (
	Int BencodeType = iota
	Str BencodeType = iota
)

type Bencode struct {
	Val  any
	Type BencodeType
}

func sliceTil(data string, b byte) (string, string, error) {
	i := strings.IndexByte(data, b)
	if i < 0 {
		return "", "", MissingDelimiter
	}
	return data[:i], data[i+1:], nil
}

func sliceFor(data string, l int) (string, error) {
	if len(data) < l {
		return "", BadLength
	}
	return data[:int16(l)], nil
}

func ParseBencode(data string) (*Bencode, error) {
	if len(data) == 0 {
		return nil, EOF
	}
	switch data[0] {
	case 'i':
		n, _, err := sliceTil(data[1:], 'e')
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		return &Bencode{i, Int}, nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		rawl, rest, err := sliceTil(data, ':')
		if err != nil {
			return nil, err
		}
		l, err := strconv.Atoi(rawl)
		if err != nil {
			return nil, err
		}
		s, err := sliceFor(rest, l)
		if err != nil {
			return nil, err
		}
		return &Bencode{s, Str}, nil
	}
	return nil, UnknownType
}
