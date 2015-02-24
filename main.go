package main

import (
	"fmt"

	"github.com/placeybordeaux/configr/unmarshaler"
)

type SpecificConfig struct {
	URL          string  `cli:"url" env:"SF_CONFIG_URL"`
	TTL          int64   `cli:"ttl" env:"SF_CONFIG_TTL" desc:"Time to live for oAuth Tokens"`
	RedirectRate float64 `cli:"redirect-rate" env:"SF_CONFIG_REDIRECT_RATE" desc:"Probability of redirect to sister node"`
	Testing      bool    `cli:"testing" env:"SF_CONFIG_TESTING_SERVER"`
}

func main() {
	//Put your defaults here
	c := SpecificConfig{URL: "http://0.0.0.0:3245", TTL: 1, Testing: false}
	//Left out defaults go to Zero value

	//ENV takes precendce over default
	configr.UnmarshalFromEnv(&c)

	//Flags take highest precidence
	configr.UnmarshalFromFlags(&c)

	fmt.Printf("%+v", c)
}
