package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sbleks/choose"
)

var Choices = []string{
	`1st choice`,
	`2nd choice`,
	`3rd choice`,
	`4th choice`,
}

var IntChoices = []int{1, 2, 120, 55}

type Entry struct {
	Name  string  `json:"name"`
	Value float32 `json:"value,omitempty"`
}

func (e Entry) MashalJSON() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(e)
	if err != nil {
		return nil, err
	}
	byt := buf.Bytes()
	return byt[:len(byt)-1], err
}

// String implements fmt.Stringer as JSON
func (e Entry) String() string {
	byt, err := e.MashalJSON()
	// byt, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
	}
	return string(byt)
}

var EntryChoices = []Entry{
	{Name: "Sam", Value: 123},
	{Name: "Sari", Value: 11.5},
	{Name: "Josh", Value: 77.987},
	{Name: "<Josh>", Value: 77.987},
}

func main() {
	pick, err := choose.From(EntryChoices)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(pick)
	// pickInt, errInt := choose.From(IntChoices)
	// if errInt != nil {
	// 	log.Println(errInt)
	// 	os.Exit(1)
	// }
	// fmt.Println(pickInt)
}
