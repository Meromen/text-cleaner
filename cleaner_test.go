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
		{
			data: "a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z",
			want: "a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z",
			config: WhiteListConfig{
				Eng:   true,
				Rus:   false,
				Dig:   false,
				AddWl: ",",
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

func TestCleanByStopWords(t *testing.T) {
	cfg := WhiteListConfig{
		Eng:   true,
		Rus:   true,
		Dig:   true,
		AddWl: "",
	}
	rawStr := "<h1>Я люблю.  111ЗPorno tut...</h1>"
	expected := "люблю 111зporno tut"
	r := strings.NewReader(rawStr)
	res := Clean(r, cfg)
	stopWordsList := make(map[string]struct{})
	stopWordsList["я"] = struct{}{}
	stopWordsList["h1"] = struct{}{}
	words := strings.Split(res, " ")
	cleanResult := CleanByStopWords(words, stopWordsList)
	actual := strings.Join(cleanResult, " ")
	assert.Equal(t, expected, actual)
}

func TestExamples(t *testing.T) {
	ExampleClean()
	ExampleCleanBytes()
	ExampleCleanString()
	ExampleCleanByStopWords()
}