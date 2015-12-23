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
func NewMessage() Message {
	return Message{
		http.Header{},
		bytes.Buffer{},
	}
}

// Writes to the Message.Body can be done by message.Write. There is no
// need to force using Message.Body.Write() on the API
func (p *Message) Write(b []byte) (int, error) {
	i, err := p.Body.Write(b)
	return i, err
}

// String provides a simple API to viewing the final payload (Headers + Body)
// as a string.
func (p *Message) String() string {
	buf := &bytes.Buffer{}
	p.WriteTo(buf)
	return buf.String()
}

// Write to w; an empty header or empty body has to be accounted from with an
// additional newline
func (p *Message) WriteTo(w io.Writer) error {
	var err error
	var b int64

	if err = p.Meta.WriteSubset(w, nil); err != nil {
		return err
	}

	if len(p.Meta) < 1 {
		w.Write([]byte(EOL))
	}

	w.Write([]byte(EOL))

	b, err = p.Body.WriteTo(w)
	if err != nil {
		return err
	}

	if b < 1 {
		w.Write([]byte(EOL))
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
	return nil
}
