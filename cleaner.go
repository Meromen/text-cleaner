// Fast text cleaner with whitelist of char.
// Clean string in one loop over all symbols input
package text_cleaner

import (
	"io"
	"strings"
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

// Clean return string in lower case after cleaning with whitelist of chars.
// whitelist is forming with WhiteListConfig
func Clean(r io.Reader, cfg WhiteListConfig) string {
	b := strings.Builder{}
	io.Copy(&b, r)
	return cleanStringWithStringBuilder(b.String(), cfg)
}


func  cleanStringWithStringBuilder(str string, cfg WhiteListConfig) string {
	bl := strings.Builder{}
	bl.Grow(len(str))
	lastWrittenRune := ' '
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

		//eng
		if cfg.Eng {
			//lower case
			if r > 96 && r < 123 {
				runeToWrite = r
			}

			//upper case
			if r > 64 && r < 91 {
				runeToWrite = r + 32
			}
		}

		// rus
		if cfg.Rus {
			//lower case
			if r > 1071 && r < 1104 {
				runeToWrite = r
			}

			//upper case
			if r > 1039 && r < 1072 {
				runeToWrite = r + 32
			}
		}

		// digits
		if cfg.Dig && r > 47 && r < 58 {
			runeToWrite = r
		}

		// chars from additional whitelist
		if additionalWhitelist != nil {
			if _, ok := additionalWhitelist[r]; ok {
				runeToWrite = r
			}
		}

		//space
		if r == 32 && lastWrittenRune != r{
			runeToWrite = r
		}

		if runeToWrite != -1 {
			lastWrittenRune = runeToWrite
			bl.WriteRune(runeToWrite)
		} else {
			if lastWrittenRune != ' '{
				bl.WriteRune(' ')
				lastWrittenRune = ' '
			}
		}
	}

	return strings.TrimSuffix(bl.String(), " ")
}
