package main

import (
	"encoding/csv"
	"fmt"
	"io"
)

func SynchronousPipeline(input *csv.Reader) {
	fmt.Println("--Synchronous pipeline----")
	input.Read()
	for {
		rec, err := input.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		out := Encode(Convert(Parse(rec)))
		fmt.Println(string(out))
	}
}
