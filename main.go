package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"
)

const EOL = "\r\n"

type Message struct {
	Header http.Header
	Body   bytes.Buffer
}

func NewMessage() Message {
	return Message{
		make(http.Header),
		bytes.Buffer{},
	}
}

func main() {
	m := NewMessage()
	m.Header.Add("test-header", "test-value")
	m.Body.Write([]byte(`this is a long string of stuff`))
	m.Put(os.Stdout)

	fmt.Println("#----------------------#")

	by := bytes.Buffer{}
	by.WriteString("Testing: this is a value\r\nAnother: value\r\n\r\nThis is a test body")
	m = Get(&by)
	m.Put(os.Stdout)
}

func (p Message) Put(w io.Writer) {
	var err error

	if err = p.Header.WriteSubset(w, nil); err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(EOL))

	if _, err = p.Body.WriteTo(w); err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(EOL))
	w.Write([]byte(EOL))
}

func Get(r io.Reader) Message {
	mm, err := mail.ReadMessage(r)

	if err != nil {
		log.Fatal(err)
	}

	by := bytes.Buffer{}
	if _, err := by.ReadFrom(mm.Body); err != nil {
		log.Fatal(err)
	}

	h := http.Header(mm.Header)

	return Message{
		h,
		by,
	}

}
