// Fast text cleaner with whitelist of char.
// Clean string in one loop over all symbols input
package text_cleaner

import (
	"io"
	"strings"
)

// Config for cleaning
// if Eng == true then cleaner keep chars of english alphabet in result string.
// if Rus == true then cleaner keep chars of english alphabet in result string.
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

func cleanStringWithStringBuilder(str string, cfg WhiteListConfig) string {
	bl := strings.Builder{}
	bl.Grow(len(str))
	spaceStored := false
	badChar := false
	firstChar := true
	var additionalWhitelist map[int32]struct{}

	//init additional whitelist
	if cfg.AddWl != "" {
		additionalWhitelist = make(map[int32]struct{}, len(cfg.AddWl))
		for _, r := range cfg.AddWl {
			additionalWhitelist[r] = struct{}{}
		}
	}

	for _, r := range str {
		//eng
		if cfg.Eng {
			//upper case
			if r > 64 && r < 91 {
				if !spaceStored && badChar && !firstChar {
					bl.WriteRune(32)
					spaceStored = true
				}
				bl.WriteRune(r + 32)
				spaceStored = false
				badChar = false
				firstChar = false
				continue
			}

			//lower case
			if r > 96 && r < 123 {
				if !spaceStored && badChar && !firstChar {
					bl.WriteRune(32)
					spaceStored = true
				}
				bl.WriteRune(r)
				spaceStored = false
				badChar = false
				firstChar = false
				continue
			}
		}

		// rus
		if cfg.Rus {
			//upper case
			if r > 1039 && r < 1072 {
				if !spaceStored && badChar && !firstChar {
					bl.WriteRune(32)
					spaceStored = true
				}
				bl.WriteRune(r + 32)
				spaceStored = false
				badChar = false
				firstChar = false
				continue
			}

			//lower case
			if r > 1071 && r < 1104 {
				if !spaceStored && badChar && !firstChar {
					bl.WriteRune(32)
					spaceStored = true
				}
				bl.WriteRune(r)
				spaceStored = false
				badChar = false
				firstChar = false
				continue
			}

		}

		// digits
		if cfg.Dig {
			if r > 47 && r < 58 {
				if !spaceStored && badChar && !firstChar {
					bl.WriteRune(32)
					spaceStored = true
				}
				bl.WriteRune(r)
				spaceStored = false
				badChar = false
				firstChar = false
				continue
			}
		}

		// chars from additional whitelist
		if additionalWhitelist != nil {
			if _, ok := additionalWhitelist[r]; ok {
				if !spaceStored && badChar && !firstChar {
					bl.WriteRune(32)
					spaceStored = true
				}
				bl.WriteRune(r)
				spaceStored = false
				badChar = false
				firstChar = false
				continue
			}
		}

		//space
		if r == 32 && !spaceStored && !badChar {
			bl.WriteRune(r)
			spaceStored = true
			badChar = false
			firstChar = false
			continue
		}

		badChar = true
	}
	return bl.String()
}
