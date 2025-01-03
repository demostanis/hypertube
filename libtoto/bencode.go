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
	End              = errors.New("unexpected end")
)

type BencodeType = int

const (
	Int  BencodeType = iota
	Str  BencodeType = iota
	List BencodeType = iota
)

type Bencode struct {
	Val  any
	Type BencodeType
	i    int
}

func (b *Bencode) sliceTil(data string, by byte) (string, string, error) {
	i := strings.IndexByte(data, by)
	if i < 0 {
		return "", "", MissingDelimiter
	}
	b.i += i
	return data[:i], data[i+1:], nil
}

func (b *Bencode) sliceFor(data string, l int) (string, error) {
	if len(data) < l {
		return "", BadLength
	}
	b.i += l
	return data[:int16(l)], nil
}

func ParseBencode(data string) (*Bencode, error) {
	if len(data) == 0 {
		return nil, EOF
	}

	b := new(Bencode)

	switch data[0] {
	case 'i':
		n, _, err := b.sliceTil(data[1:], 'e')
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		b.Val = i
		b.Type = Int
		return b, nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		rawl, rest, err := b.sliceTil(data, ':')
		if err != nil {
			return nil, err
		}
		l, err := strconv.Atoi(rawl)
		if err != nil {
			return nil, err
		}
		s, err := b.sliceFor(rest, l)
		if err != nil {
			return nil, err
		}
		b.Val = s
		b.Type = Str
		return b, nil
	case 'l':
		i := 1
		l := make([]any, 0)
		for {
			res, err := ParseBencode(data[i:])
			if err == End {
				b.Val = l
				b.Type = List
				return b, nil
			}
			if err != nil {
				return nil, err
			}
			l = append(l, res.Val)
			i += res.i + 1
		}
	case 'e':
		return nil, End
	}
	return nil, UnknownType
}
