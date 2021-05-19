text-cleaner is a fast text cleaner with whitelist of char. Clean string in one loop over all symbols input

## Getting started

## Installing

To start using text-cleaner, install Go and run go get:

```sh
$ go get -u github.com/Meromen/text-cleaner
```

## Usage

Clear from reader

```
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
```

Clean string

```
str := "Hello it's    fast cleaner очень быстрый, быстрее 10 самолетов....    "

cfg := WhiteListConfig{
	Eng:   true,
	Rus:   false,
	Dig:   true,
	AddWl: "'т",
}

result := CleanString(str, cfg)
fmt.Println(result) // "Hello it's fast cleaner т т 10 т"
```

Clean slice of byte

```
sb := []byte("123 asd   фыв   ")

cfg := WhiteListConfig{
	Eng:   true,
	Rus:   true,
	Dig:   false,
	AddWl: "",
}

result := CleanBytes(sb, cfg)
fmt.Println(result) // "asd фыв"
```

Clean with black list

```
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
```

## Benchmark

```
// Clean with text-cleaner 
BenchmarkClean-8        23317753                50.2 ns/op            32 B/op          1 allocs/op
// Clean with regexp
BenchmarkRegExp-8          26359             45454 ns/op            4133 B/op         11 allocs/op

``` 

## Contacts

Yuri Pysin [@Meromen](https://github.com/Meromen)

## Licence

`text-cleaner` source code is available under the MIT [Licence](/LICENSE)
