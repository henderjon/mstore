package mstore

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var testPayload = "Content-Type: text/plain; charset=UTF-8\r\n" +
	"Header-One: Testing mstore\r\n" +
	"\r\n" +
	"This is the body of the message.\r\n"

func TestReadFrom(t *testing.T) {
	m := NewMessage()

	ex := bytes.NewBufferString(testPayload)
	m.ReadFrom(ex)

	if m.Meta.Get("Header-One") != "Testing mstore" {
		t.Error("header mismatch")
	}

	b, _ := ioutil.ReadAll(&m.Body)
	if string(b) != "This is the body of the message.\r\n" {
		t.Error("body mismatch")
	}
}

func TestWriteTo(t *testing.T) {
	m := NewMessage()
	m.Meta.Add("Header-One", "Testing mstore")
	m.Meta.Add("Content-Type", "text/plain; charset=UTF-8")
	m.Write([]byte("This is the body of the message."))

	buf := make([]byte, 0)
	ex := bytes.NewBuffer(buf)
	m.WriteTo(ex)

	if ex.String() != testPayload {
		t.Error("body mismatch")
	}
}
