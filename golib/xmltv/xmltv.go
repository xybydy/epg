package xmltv

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"time"
)

const timefmt = "20060102150405 -0700"
const encodingre = `encoding=".*"`

type Tv struct {
	XMLName           xml.Name     `xml:"tv"`
	Channel           []*Channel   `xml:"channel"`
	Programme         []*Programme `xml:"programme"`
	SourceInfoURL     string       `xml:"source-info-url,attr"`  // is a URL describing the data source in some human-readable form
	SourceInfoName    string       `xml:"source-info-name,attr"` // the link text for that URL
	SourceDataURL     string       `xml:"source-data-url,attr"`  //is where the actual data is grabbed from.
	GeneratorInfoName string       `xml:"generator-info-name"`
	GeneratorInfoURL  string       `xml:"generator-info-url"`
}

func NewXMLTVFile() *Tv {
	xmltvf := &Tv{}
	xmltvf.Channel = make([]*Channel, 0)
	xmltvf.Programme = make([]*Programme, 0)

	return xmltvf
}

type Channel struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"display-name"`
	Icon Icon   `xml:"icon"`
}

type Icon struct {
	Src    string `xml:"src,attr"`
	Width  string `xml:"width,attr,omitempty"`
	Height string `xml:"height,attr,omitempty"`
}

type Programme struct {
	ID       string   `xml:"id,attr,omitempty"`
	Start    string   `xml:"start,attr"`
	Stop     string   `xml:"stop,attr"`
	Channel  string   `xml:"channel,attr"`
	Title    string   `xml:"title,omitempty"`
	SubTitle string   `xml:"sub-title,omitempty"`
	Desc     string   `xml:"desc,omitempty"`
	Credits  Credits  `xml:"credits,omitempty"`
	Date     string   `xml:"date,omitempty"`
	Category []string `xml:"category,omitempty"`
	Rating   string   `xml:"rating>value,omitempty"`
}

type Credits struct {
	Director    string `xml:"director,omitempty"`
	Actor       string `xml:"actor,omitempty"`
	Writer      string `xml:"writer,omitempty"`
	Adapter     string `xml:"adapter,omitempty"`
	Producer    string `xml:"producer,omitempty"`
	Composer    string `xml:"composer,omitempty"`
	Editor      string `xml:"editor,omitempty"`
	Presenter   string `xml:"presenter,omitempty"`
	Commentator string `xml:"commentator,omitempty"`
	Guest       string `xml:"guest,omitempty"`
}

func ParseTime(t string) (time.Time, error) {
	return time.Parse(timefmt, t)
}

func TimeString(t time.Time) string {
	return t.Format(timefmt)
}

func Marshal(v interface{}) ([]byte, error) {
	data, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return data, err
	}

	data = append([]byte(xml.Header), data...)

	return data, err
}

func Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

func ReadFile(path string) (*Tv, error) {
	if _, err := os.Stat(path); err != nil {
		fmt.Errorf("%v", err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data = stripEncoding(data)
	xmltvf := NewXMLTVFile()

	err = Unmarshal(data, xmltvf)
	if err != nil {
		return nil, err
	}

	return xmltvf, nil
}

func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}

func stripEncoding(str []byte) []byte {
	reg := regexp.MustCompile(encodingre)

	res := reg.ReplaceAllString(string(str), "")

	return []byte(res)
}
