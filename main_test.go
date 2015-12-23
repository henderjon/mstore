package mstore

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var testPayload = "Content-Type: text/plain; charset=UTF-8" + EOL +
	"Header-One: Testing mstore" + EOL +
	"" + EOL +
	"This is the body of the message." + EOL

var testEmptyBody = "Content-Type: text/plain; charset=UTF-8" + EOL +
	"Header-One: Testing mstore" + EOL +
	// "" + EOL +
	EOL

var testEmptyHeaders = "" + EOL +
	"This is the body of the message." + EOL

func TestReadFrom(t *testing.T) {
	m := NewMessage()

	ex := bytes.NewBufferString(testPayload)
	m.ReadFrom(ex)

	if m.Meta.Get("Header-One") != "Testing mstore" {
		t.Error("header mismatch")
	}

	b, _ := ioutil.ReadAll(&m.Body)
	if !bytes.Equal(b, []byte("This is the body of the message." + EOL))  {
		t.Error("body mismatch")
	}
}

func TestReadFromEmptyBody(t *testing.T) {
	m := NewMessage()

	ex := bytes.NewBufferString(testEmptyBody)
	m.ReadFrom(ex)

	if m.Meta.Get("Header-One") != "Testing mstore" {
		t.Error("header mismatch")
	}

	b, _ := ioutil.ReadAll(&m.Body)
	if string(b) != "" {
		t.Error("body mismatch")
	}

}

func TestReadFromEmptyHeaders(t *testing.T) {
	m := NewMessage()

	ex := bytes.NewBufferString(testEmptyHeaders)
	m.ReadFrom(ex)

	b, _ := ioutil.ReadAll(&m.Body)
	if !bytes.Equal(b, []byte("This is the body of the message." + EOL))  {
		t.Error("body mismatch")
	}
}

func TestWriteTo(t *testing.T) {
	m := NewMessage()
	m.Meta.Add("Header-One", "Testing mstore")
	m.Meta.Add("Content-Type", "text/plain; charset=UTF-8")
	m.Write([]byte("This is the body of the message."))

	if m.String() != testPayload {
		t.Error("body mismatch")
	}
}
