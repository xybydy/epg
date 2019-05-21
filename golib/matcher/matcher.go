package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"

	"github.com/masatana/go-textdistance"
)

type EpgList []struct {
	ChannelName string  `json:"channel_name,omitempty"`
	EpgCode     string  `json:"epg_code,omitempty"`
	Score       float64 `json:"score,omitempty"`
}

func (e EpgList) Len() int {
	return len(e)
}

//Reversed
func (e EpgList) Less(i, j int) bool {
	if e[i].Score < e[j].Score {
		return false
	}
	return true
}

func (e EpgList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *EpgList) Match(name string) string {
	if len(*e) == 0 {
		e.GetAll("")
	}

	for i := range *e {
		(*e)[i].Score = textdistance.JaroWinklerDistance(name, (*e)[i].ChannelName)
	}

	sort.Sort(e)

	return (*e)[0].EpgCode

}

func (e *EpgList) GetAll(file string) {
	if file == "" {
		file = "merged.json"
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(b, &e)
	if err != nil {
		log.Println(err)
	}

}

// func main() {
// 	s := new(EpgList)
// 	s.Match("beIN BOX OFFICE 3 TR")
// 	log.Println((*s)[0])

// }
