/*
Package mstore provides a means for serializing data in a format similar to raw
email or http requests. Messages compose headers and a body. The advantage to
this is Messages are both human readable and machine parsable.
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"os"
)

const EOL = "\r\n"

// messages have two parts: Meta and Body. Messages very closely resemble a raw
// email message or raw HTTP request with the absence of a few distinctive
// features. Namely, there are no required headers. In the case of email `From`
// and in the case of HTTP requests the opener
type Message struct {
	Meta http.Header
	Body *bufio.ReadWriter
}

// create a new message
func NewMessage() *Message {
	b := bytes.Buffer{}
	rw := bufio.NewReadWriter(bufio.NewReader(&b), bufio.NewWriter(&b))
	h := http.Header{}
	return &Message{
		h,
		rw,
	}
}

// hide the Message.Body.Flush() from the API
func (p *Message) Write(b []byte) (int, error) {
	i, err := p.Body.Write(b)
	p.Body.Flush()
	return i, err
}

// write to w, ending with two blank lines
func (p *Message) SerializeTo(w io.Writer) {
	var err error

	if err = p.Meta.WriteSubset(w, nil); err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(EOL))

	if _, err = p.Body.WriteTo(w); err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(EOL))
	w.Write([]byte(EOL))
	w.Write([]byte(EOL))
}

// read from r, parsing into a Message{ Meta, Body }
func (p *Message) DeserializeFrom(r io.Reader) {
	mm, err := mail.ReadMessage(r)

	if err != nil {
		log.Fatal(err)
	}

	by := bytes.Buffer{}

	buf := bufio.NewReadWriter(
		bufio.NewReader(&by),
		bufio.NewWriter(&by),
	)

	if _, err := buf.ReadFrom(mm.Body); err != nil {
		log.Fatal(err)
	}

	h := http.Header(mm.Header)

	p.Meta = h
	p.Body = buf

}

func main() {

	f, _ := os.Open("test.mbox")
	m := NewMessage()
	m.DeserializeFrom(f)
	fmt.Println(m.Meta.Get("X-Powered-By"))

	b, _ := ioutil.ReadAll(m.Body)
	fmt.Println(string(b))

	// m.SerializeTo(os.Stdout)

	// m := NewMessage()
	// m.Meta.Add("test-header", "test-value")
	// m.Write([]byte(`this is a long string of stuff`))
	// m.Write([]byte(`with even more after that`))
	// m.Put(os.Stdout)

	// fmt.Println("#----------------------#")

	// by := &bytes.Buffer{}
	// by.WriteString("Testing: this is a value\r\nAnother: value\r\n\r\nThis is a test body")
	// m.Get(by)
	// m.Put(os.Stdout)
}
