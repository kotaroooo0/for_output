package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	RankingCount   int  `toml:"ranking_count"`
	SlowCount      int  `toml:"slow_count"`
	ShowStdDev     bool `toml:"show_stddev"`
	ShowStatusCode bool `toml:"show_status_code"`
	Percentiles    []float64
	Scale          int
	EffectiveDigit int    `toml:"effective_digit"`
	LogFormat      string `toml:"log_format"`
	RequestIndex   int    `toml:"request_index"`
	StatusIndex    int    `toml:"status_index"`
	DurationIndex  int    `toml:"duration_index"`
	Bundle         []bundleConfig
	Replace        []replaceConfig
	Bundles        map[string]bundleConfig // for backward compatibility

	ShowBytes  bool `toml:"show_bytes"`
	BytesIndex int  `toml:"bytes_index"`
}

type bundleConfig struct {
	Name   string
	Regexp string
}

type replaceConfig struct {
	Regexp  string
	Replace string
}

var config tomlConfig

func main() {
	if _, err := toml.DecodeFile("kataribe.toml", &config); err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}
	fmt.Println(config)

	// done := make(chan struct{})

	urlNormalizeRegexps := make(map[string]*regexp.Regexp)

	// chBundle := make(chan bundleConfig)
	// go func() {
	// 	for bundle := range chBundle {
	// 		name := bundle.Name
	// 		if name == "" {
	// 			name = bundle.Regexp
	// 		}
	// 		urlNormalizeRegexps[name] = regexp.MustCompile(bundle.Regexp)
	// 	}
	// 	done <- struct{}{}
	// }()

	// for _, b := range config.Bundle {
	// 	chBundle <- b
	// }
	// for _, b := range config.Bundles {
	// 	chBundle <- b
	// }
	// close(chBundle)
	// <-done

	var bs []bundleConfig
	for _, b := range config.Bundle {
		bs = append(bs, b)
	}
	for _, b := range config.Bundles {
		bs = append(bs, b)
	}
	for _, bundle := range bs {
		name := bundle.Name
		if name == "" {
			name = bundle.Regexp
		}
		urlNormalizeRegexps[name] = regexp.MustCompile(bundle.Regexp)
	}

	fmt.Println(urlNormalizeRegexps)
}
