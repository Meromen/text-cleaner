package text_cleaner

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestClean(t *testing.T) {

	var tests = []struct {
		data   string
		want   string
		config WhiteListConfig
	}{
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "h1 я люблю 111зporno tut h1",
			config: WhiteListConfig{
				Eng:   true,
				Rus:   true,
				Dig:   true,
				AddWl: "",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "1 я люблю 111з 1",
			config: WhiteListConfig{
				Eng:   false,
				Rus:   true,
				Dig:   true,
				AddWl: "",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "h1 111 porno tut h1",
			config: WhiteListConfig{
				Eng:   true,
				Rus:   false,
				Dig:   true,
				AddWl: "",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "h я люблю зporno tut h",
			config: WhiteListConfig{
				Eng:   true,
				Rus:   true,
				Dig:   false,
				AddWl: "",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "1 111 1",
			config: WhiteListConfig{
				Eng:   false,
				Rus:   false,
				Dig:   true,
				AddWl: "",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "я люблю з",
			config: WhiteListConfig{
				Eng:   false,
				Rus:   true,
				Dig:   false,
				AddWl: "",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "h porno tut h",
			config: WhiteListConfig{
				Eng:   true,
				Rus:   false,
				Dig:   false,
				AddWl: "",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "",
			config: WhiteListConfig{
				Eng:   false,
				Rus:   false,
				Dig:   false,
				AddWl: "",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: ". ...",
			config: WhiteListConfig{
				Eng:   false,
				Rus:   false,
				Dig:   false,
				AddWl: ".",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "/",
			config: WhiteListConfig{
				Eng:   false,
				Rus:   false,
				Dig:   false,
				AddWl: "/",
			},
		},

		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "< </",
			config: WhiteListConfig{
				Eng:   false,
				Rus:   false,
				Dig:   false,
				AddWl: "</",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "< > < >",
			config: WhiteListConfig{
				Eng:   false,
				Rus:   false,
				Dig:   false,
				AddWl: "<>",
			},
		},
		{
			data: "<h1>Я люблю.  111ЗPorno tut...</h1>",
			want: "h porno tut /h",
			config: WhiteListConfig{
				Eng:   true,
				Rus:   false,
				Dig:   false,
				AddWl: "/",
			},
		},
	}

	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			r := strings.NewReader(tt.data)
			str := Clean(r, tt.config)
			assert.Equal(t, tt.want, str)
		})
	}
}
