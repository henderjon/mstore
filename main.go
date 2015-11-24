/*
Package mstore provides a means for serializing data in a format similar to raw
email or http requests. Messages compose headers and a body. The advantage to
this is Messages are both human readable and machine parsable. It is HIGHLY
recommended that each Message be stored in its own file as a typical MBOX file
uses a bare `From` line to seperate each message and mstore only uses two
blank lines.
*/
package mstore

import (
	"bytes"
	"io"
	"net/http"
	"net/mail"
)

const EOL = "\r\n"

// Messages have two parts: Meta and Body. Messages very closely resemble a raw
// email message or raw HTTP request with the absence of a few distinctive
// headers.
type Message struct {
	Meta http.Header
	Body bytes.Buffer
}

// Creates a new Message
func NewMessage() *Message {
	return &Message{
		http.Header{},
		bytes.Buffer{},
	}
}

// Writes to the Message.Body can be done by Message.Write. There is no
// need to force using Message.Body.Write() on the API
func (p *Message) Write(b []byte) (int, error) {
	i, err := p.Body.Write(b)
	return i, err
}

// Write to w, ending with two blank lines
func (p *Message) WriteTo(w io.Writer) error {
	var err error

	if err = p.Meta.WriteSubset(w, nil); err != nil {
		return err
	}

	w.Write([]byte(EOL))

	if _, err = p.Body.WriteTo(w); err != nil {
		return err
	}

	w.Write([]byte(EOL))
	return nil
}

// Read from r, parsing into Message{ Meta, Body }
func (p *Message) ReadFrom(r io.Reader) error {
	mm, err := mail.ReadMessage(r)
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(mm.Body); err != nil {
		return err
	}

	h := http.Header(mm.Header)

	p.Meta = h
	p.Body = buf
}

// Parse from r, parsing into a new message
// func Parse(r io.Reader) *Message {
// 	m := NewMessage()
// 	m.ReadFrom(r)
// 	return m
// }
