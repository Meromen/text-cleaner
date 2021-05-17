package text_cleaner

import (
	"fmt"
	"net/http"
)

func ExampleClean() {
	res, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	cfg := WhiteListConfig{
		Eng:   true,
		Rus:   true,
		Dig:   true,
		AddWl: "",
	}

	result := Clean(res.Body, cfg)
	fmt.Println(result)
}

func ExampleCleanBytes() {
	sb := []byte("123 asd   фыв   ")

	cfg := WhiteListConfig{
		Eng:   true,
		Rus:   true,
		Dig:   false,
		AddWl: "",
	}

	result := CleanBytes(sb, cfg)
	fmt.Println(result) // "asd фыв"
}

func ExampleCleanString() {
	str := "Hello it's    fast cleaner очень быстрый, быстрее 10 самолетов....    "

	cfg := WhiteListConfig{
		Eng:   true,
		Rus:   false,
		Dig:   true,
		AddWl: "'т",
	}

	result := CleanString(str, cfg)
	fmt.Println(result) // "Hello it's fast cleaner т т 10 т"
}

func ExampleCleanStringWithBlackList() {
	str := "Hello it's    fast cleaner очень быстрый, быстрее 10 самолетов....    "
	stopWordsList := map[string]struct{}{
		"fast":    {},
		"быстрый": {},
		"10":      {},
	}

	cfg := WhiteListConfig{
		Eng:   true,
		Rus:   true,
		Dig:   true,
		AddWl: "",
	}

	blCfg := BlackListConfig{
		BlackList: stopWordsList,
	}

	result := CleanStringWithBlackList(str, cfg, blCfg)
	fmt.Println(result) // "hello it s cleaner очень быстрее самолетов"
}
