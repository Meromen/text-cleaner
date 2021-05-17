// Fast text cleaner with whitelist of char.
// Clean string in one loop over all symbols input
package text_cleaner

import (
	"io"
	"strings"
	"bytes"
)

// Config for cleaning
// if Eng == true then cleaner keep chars of english alphabet in result string.
// if Rus == true then cleaner keep chars of russian alphabet in result string.
// if Dig == true then cleaner keep digits in result string.
// AddWl it's additional whitelist of chars,
// pass ",|" if you want keep ',' and '|' in result string.
type WhiteListConfig struct {
	Eng   bool
	Rus   bool
	Dig   bool
	AddWl string
}

const (
	RuneSpace = ' '
)

// Clean return string in lower case after cleaning with whitelist of chars.
// whitelist is forming with WhiteListConfig
func Clean(r io.Reader, cfg WhiteListConfig) string {
	b := strings.Builder{}
	io.Copy(&b, r)
	return CleanString(b.String(), cfg)
}

func CleanBytes (sb []byte, cfg WhiteListConfig) string {
	b := bytes.Buffer{}
	b.Write(sb)
	return CleanString(b.String(), cfg)
}

func CleanString(str string, cfg WhiteListConfig) string {
	bl := strings.Builder{}
	bl.Grow(len(str))
	lastWrittenRune := RuneSpace
	var runeToWrite int32
	var additionalWhitelist map[int32]struct{}

	//init additional whitelist
	if cfg.AddWl != "" {
		additionalWhitelist = make(map[int32]struct{}, len(cfg.AddWl))
		for _, r := range cfg.AddWl {
			additionalWhitelist[r] = struct{}{}
		}
	}

	for _, r := range str {
		runeToWrite = -1

		switch {
		// english alphabet
		case cfg.Eng && isEnglishLowerCaseRune(r):
			runeToWrite = r
		case cfg.Eng && isEnglishUpperCaseRune(r):
			runeToWrite = r + 32

		// russian alphabet
		case cfg.Rus && isRussianLowerCaseRune(r):
			runeToWrite = r
		case cfg.Rus && isRussianUpperCaseRune(r):
			runeToWrite = r + 32

		// 	digits
		case cfg.Dig && isDigitRune(r):
			runeToWrite = r

		//	space
		case r == RuneSpace && lastWrittenRune != RuneSpace:
			runeToWrite = r
		}

		// chars from additional whitelist
		if additionalWhitelist != nil {
			if _, ok := additionalWhitelist[r]; ok {
				runeToWrite = r
			}
		}

		if runeToWrite != -1 {
			lastWrittenRune = runeToWrite
			bl.WriteRune(runeToWrite)
		} else if lastWrittenRune != RuneSpace {
			bl.WriteRune(RuneSpace)
			lastWrittenRune = RuneSpace
		}
	}

	return strings.TrimSuffix(bl.String(), " ")
}

func CleanByStopWords(words []string, stopList map[string]struct{}) []string {
	var cleanList []string
	for _, word := range words {
		if _, ok := stopList[word]; !ok {
			cleanList = append(cleanList, word)
		}
	}
	return cleanList
}

func isRussianLowerCaseRune(r int32) bool {
	return r > 1071 && r < 1104
}

func isRussianUpperCaseRune(r int32) bool {
	return r > 1039 && r < 1072
}

func isEnglishLowerCaseRune(r int32) bool {
	return r > 96 && r < 123
}

func isEnglishUpperCaseRune(r int32) bool {
	return r > 64 && r < 91
}

func isDigitRune(r int32) bool {
	return r > 47 && r < 58
}
