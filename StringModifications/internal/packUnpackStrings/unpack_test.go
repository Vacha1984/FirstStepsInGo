package packUnpackStrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"startsWithNumber", args{"3xxx"}, ""},
		{"simpleStr", args{"xxx"}, "xxx"},
		{"hasNumber", args{"x4z"}, "xxxxz"},
		{"hasNumberInEnd", args{"abcd5"}, "abcddddd"},
		{"hasNull", args{"abcd0"}, "abc"},
		{"hasBigNumber", args{"abcd15"}, "abcddddddddddddddd"},
		{"hasSpecChar", args{"abc\n2"}, "abc\n\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Unpack(tt.args.s), "%s != %s", tt.want, Unpack(tt.args.s))
			//if got := Unpack(tt.args.s); got != tt.want {
			//t.Errorf("Unpack() = %v, want %v", got, tt.want)
			//}
		})
	}
}
