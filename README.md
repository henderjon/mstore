# mstore

[![GoDoc](https://godoc.org/github.com/henderjon/mstore?status.svg)](https://godoc.org/github.com/henderjon/mstore)
[![Build Status](https://travis-ci.org/henderjon/mstore.svg?branch=master)](https://travis-ci.org/henderjon/mstore)
[![License](https://img.shields.io/badge/license-BSD--3%20Clause-blue.svg)](LICENSE.md)

An all-too-simple data de/serializer. It reads and writes data in a format similar to raw email or raw http requests.

## example

```
Header-Name: val/value; type=val
Another-Header: single value

This is the payload which can actually be anything. Content-Type/Length headers aren't determinitive.
```


