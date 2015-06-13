package epp

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/nbio/st"
)

func TestTime(t *testing.T) {
	var x = []byte(`<example><when>2015-05-19T06:34:21.1Z</when></example>`)
	var y struct {
		XMLName struct{} `xml:"example"`
		When    Time     `xml:"when"`
	}

	err := xml.Unmarshal(x, &y)
	st.Expect(t, err, nil)
	t.Logf("%#v\n", &y)
	tt, _ := time.Parse(time.RFC3339, "2015-05-19T06:34:21.1Z")
	st.Expect(t, y.When, Time{tt})
	z, err := xml.Marshal(&y)
	st.Expect(t, err, nil)
	st.Expect(t, string(z), string(x))
	text, err := y.When.MarshalText()
	st.Expect(t, err, nil)
	st.Expect(t, string(text), "2015-05-19T06:34:21.1Z")
}
