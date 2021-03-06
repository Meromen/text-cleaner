package text_cleaner

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

var str = "\n\n\n<!DOCTYPE html>\n<html lang=\"en\" data-color-mode=\"dark\">\n  <head>\n    <meta charset=\"utf-8\">\n  <link rel=\"dns-prefetch\" href=\"https://github.githubassets.com\">\n  <link rel=\"dns-prefetch\" href=\"https://avatars0.githubusercontent.com\">\n  <link rel=\"dns-prefetch\" href=\"https://avatars1.githubusercontent.com\">\n  <link rel=\"dns-prefetch\" href=\"https://avatars2.githubusercontent.com\">\n  <link rel=\"dns-prefetch\" href=\"https://avatars3.githubusercontent.com\">\n  <link rel=\"dns-prefetch\" href=\"https://github-cloud.s3.amazonaws.com\">\n  <link rel=\"dns-prefetch\" href=\"https://user-images.githubusercontent.com/\">\n\n\n\n  <link crossorigin=\"anonymous\" media=\"all\" integrity=\"sha512-BSy+E+S5PJuDWKcXiIXBoFJ7uJ+88y6hFdIhZpf7nf9MVNVvnJDPUaotaxFUQi8UXCLJOcGv1uifxVMc9o5DYQ==\" rel=\"stylesheet\" href=\"https://github.githubassets.com/assets/frameworks-052cbe13e4b93c9b8358a7178885c1a0.css\" />\n  \n    <link crossorigin=\"anonymous\" media=\"all\" integrity=\"sha512-bRic7lVTQ3HMC+Xi6jXFdLpMFO5Yl6b+apchURCSN1cjAKteDgORh3nfXWuG1zLqbEVHYXWC9G/W1VQ3IYd32Q==\" rel=\"stylesheet\" href=\"https://github.githubassets.com/assets/github-6d189cee55534371cc0be5e2ea35c574.css\" />"

var Result string

func BenchmarkClean(b *testing.B) {
	r := strings.NewReader(str)
	cnf := WhiteListConfig{
		Eng:   true,
		Rus:   true,
		Dig:   true,
		AddWl: "",
	}

	for i := 0; i < b.N; i++ {
		Result = Clean(r, cnf)
	}
}

func BenchmarkRegExp(b *testing.B) {
	reg, err := regexp.Compile("[^0-9a-zA-Zа-яА-Я]+")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		Result = reg.ReplaceAllString(str, " ")
		Result = strings.ToLower(Result)
		Result = strings.TrimSpace(Result)
	}
}
