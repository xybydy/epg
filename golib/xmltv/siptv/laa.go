package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sort"
)

type rs struct {
	ChannelName string `json:"channel_name"`
	EpgCode     string `json:"epg_code"`
	Popularity  int    `json:"popularity"`
}

type rss []rs

// Len is the number of elements in the collection.
func (r rss) Len() int {
	return len(r)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (r rss) Less(i int, j int) bool {
	return r[i].ChannelName < r[j].ChannelName
}

// Swap swaps the elements with indexes i and j.
func (r rss) Swap(i int, j int) {
	r[i], r[j] = r[j], r[i]
}

func main() {
	qwe := rss{}
	ww := rss{}
	q, _ := filepath.Glob("*.json")
	for _, i := range q {
		b, _ := ioutil.ReadFile(i)
		json.Unmarshal(b, &ww)
		for _, k := range ww {
			qwe = append(qwe, k)
		}
	}
	sort.Sort(qwe)
	ff, _ := json.Marshal(qwe)
	ioutil.WriteFile("laaa.json", ff, 0644)
}
