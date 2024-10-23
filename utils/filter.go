package utils

import (
	"strings"
	snowballeng "github.com/kljensen/snowball/english"
)

func lowercaseFilter(tokens []string) []string{
	r := make([]string, len(tokens))
	for i, token := range tokens{
		r[i] = strings.ToLower(token)
	}
	return r
}

func stopwordFilter(tokens []string) []string{
	var stopwords = map[string]struct{}{
		"a": {}, "about": {}, "after": {}, "again": {}, "all": {}, 
		"also": {}, "an": {}, "and": {}, "any": {}, "are": {}, 
		"as": {}, "at": {}, "be": {}, "because": {}, "been": {}, 
		"before": {}, "but": {}, "by": {}, "can": {}, "could": {}, 
		"did": {}, "do": {}, "does": {}, "for": {}, "from": {}, 
		"had": {}, "has": {}, "have": {}, "he": {}, "her": {}, 
		"here": {}, "him": {}, "his": {}, "how": {}, "if": {}, 
		"in": {}, "into": {}, "is": {}, "it": {}, "its": {}, 
		"just": {}, "like": {}, "me": {}, "more": {}, "my": {}, 
		"no": {}, "not": {}, "now": {}, "of": {}, "on": {}, 
		"one": {}, "or": {}, "our": {}, "out": {}, "said": {}, 
		"she": {}, "so": {}, "some": {}, "than": {}, "that": {}, 
		"the": {}, "their": {}, "them": {}, "then": {}, "there": {}, 
		"these": {}, "they": {}, "this": {}, "to": {}, "too": {}, 
		"up": {}, "us": {}, "was": {}, "we": {}, "were": {}, 
		"what": {}, "when": {}, "where": {}, "which": {}, "who": {}, 
		"will": {}, "with": {}, "you": {}, "your": {},
	}
	r := make([]string, 0, len(tokens))
	for  _, token := range tokens{
		if _, ok :=  stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func stemmerFilter(tokens []string) []string{
	r := make([]string, len(tokens))
	for i, token := range tokens{
		r[i] = snowballeng.Stem(token, false)
	}
	return r 
}