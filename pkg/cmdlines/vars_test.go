package cmdlines

import (
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/bytedance/mockey"
)

func Test_getValidShellVarName(t *testing.T) {
	tests := []struct {
		input string
		want0 string
		want1 int
	}{
		{
			input: "abc.",
			want0: "abc",
			want1: 3,
		},
		{
			input: "$abc",
			want0: "",
			want1: 0,
		},
		{
			input: "123_abc_$abc",
			want0: "123_abc_",
			want1: 8,
		},
		{
			input: "_01234abc_edfZEJ",
			want0: "_01234abc_edfZEJ",
			want1: len("_01234abc_edfZEJ"),
		},
		{
			input: "123.05",
			want0: "123",
			want1: 3,
		},
		{
			input: "",
			want0: "",
			want1: 0,
		},
		{
			input: ".",
			want0: "",
			want1: 0,
		},
		{
			"",
			"",
			0,
		},
		{
			"{",
			"",
			0,
		},
		{
			"{}",
			"",
			2,
		},
		{
			"{aaa}",
			"aaa",
			5,
		},
		{
			"{abc23",
			"",
			0,
		},
		{
			"abc23}",
			"abc23",
			5,
		},
		{
			"abc23{}",
			"abc23",
			5,
		},
		{
			"{ }",
			"",
			0,
		},
		{
			"{}}",
			"",
			2,
		},
	}
	for i, tt := range tests {
		t.Run("test #"+strconv.Itoa(i), func(t *testing.T) {
			if got, got1 := GetValidShellVarName(tt.input); got != tt.want0 || got1 != tt.want1 {
				t.Errorf("GetValidShellVarName() = %v, %v, want %v, %v", got, got1, tt.want0, tt.want1)
			}
		})
	}
}

func Test_splitShellVars(t *testing.T) {
	mockey.Mock(os.LookupEnv).Return("", false).Build()
	cache := map[string]string{"": ""}
	tests := []struct {
		input string
		cache map[string]string
		want  []string
	}{
		{
			"abc",
			cache,
			[]string{"abc"},
		},
		{
			"$abc",
			cache,
			[]string{""},
		},
		{
			"$abc",
			map[string]string{"": "", "abc": "abc"},
			[]string{"abc"},
		},
		{
			"abc$abc$$${abcd}",
			map[string]string{"": "", "abc": "abc"},
			[]string{"abc", "abc", "$$", ""},
		},
		{
			"print env abc:\n${abc}",
			map[string]string{"": "", "abc": "abc"},
			[]string{"print env abc:\n", "abc"},
		},
	}
	for i, tt := range tests {
		t.Run("case #"+strconv.Itoa(i), func(t *testing.T) {
			if got := SplitShellVars(tt.cache, tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitShellVars() = %v, want %v", got, tt.want)
			}
		})
	}
}
