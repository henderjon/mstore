package main

import (
	// "fmt"
	"github.com/henderjon/mstore"
	// "io/ioutil"
	"os"
)

func main() {

	m := mstore.NewMessage()
	m.Meta.Add("Content-Type", "text/plain; charset=UTF-8")
	m.Meta.Add("X-Powered-By", "henderjon/mstore; golang=v1.5.1")
	m.Write([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit"))
	m.WriteTo(os.Stdout)

	// fmt.Println("#----------------------#")

	// f, _ := os.Open("test.mbox")
	// m := mstore.NewMessage()
	// m.ReadFrom(f)
	// fmt.Println(m.Meta.Get("X-Powered-By"))

	// b, _ := ioutil.ReadAll(&m.Body)
	// fmt.Println(string(b))

}
