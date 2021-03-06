package main

import (
	"testing"
)

func TestCheckRegex(t *testing.T) {
	var results []string
	checks := map[string]int{
		"github.com":                                                 0,
		"github.com/user/":                                           0,
		"github.com/user -- Sys":                                     0,
		"github_api_client = \"sample key\"\naws=afewafewafewafewaf": 2,
		"aws=\"afewafewafewafewaf\"":                                 1,
		"aws\"afewafewafewafewaf\"":                                  0,
		"heroku := \"afewafewafewafewaf\"":                           1,
		"heroku_client_secret := \"afewafewafewafewaf\"":             1,
		"reddit_api_secreit = \"Fwe4fa431FgklreF\"":                  1,
	}

	for k, v := range checks {
		results = checkRegex(k)
		if v != len(results) {
			t.Errorf("regexCheck failed on string %s", k)
		}
	}
}

func TestEntropy(t *testing.T) {
	var enoughEntropy bool
	checks := map[string]bool{
		"heroku_client_secret = settings.HEROKU_CLIENT": false,
		"heroku_client_secret = conf.heroku":            false,
		"reddit_secret = settings.REDDIT_API":           false,
		"reddit_api_secret = \"Fwe4fa431FgklreF\"":      true,
		"aws_secret= \"AKIAIMNOJVGFDXXXE4OA\"":          true,
	}
	for k, v := range checks {
		enoughEntropy = checkEntropy(k)
		if v != enoughEntropy {
			t.Errorf("checkEntropy failed for %s. Expected %t, got %t", k, v, enoughEntropy)
		}
	}

}
