package xmltv

import (
	"fmt"
	"io/ioutil"
	"sort"
	"testing"
)

func TestReadFile(t *testing.T) {
	qq := []string{}
	v := ""
	t.Run("Turkey", func(t *testing.T) {

		for _, i := range []string{"epg.xml", "epghiz.xml", "siptv.xml"} {
			q, err := ReadFile(i)
			if err != nil {
				t.Errorf("READ: %v", err)
			}
			for _, k := range q.Channels {
				// v += fmt.Sprintf("%s\n", k.Id)
				qq = append(qq, k.Id)
			}
		}
		sort.Strings(qq)
		for _, k := range qq {
			v += fmt.Sprintf("%s\n", k)
		}

		ioutil.WriteFile("cikti.txt", []byte(v), 0644)
	})
}
