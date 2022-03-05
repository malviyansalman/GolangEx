package util

import (
	"github.com/gophercises/quiet_hn/hn"
	"net/url"
	"strings"
	"time"
)

func IsStoryLink(item Item) bool {
	return item.Type == "story" && item.URL != ""
}

func ParseHNItem(hnItem hn.Item) Item {
	ret := Item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type Item struct {
	hn.Item
	Host string
}

type TemplateData struct {
	Stories []Item
	Time    time.Duration
}

type Result struct {
	Item Item
	Err  error
	Idx  int
}
