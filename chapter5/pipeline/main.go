package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type Record struct {
	Row int `json:"row"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}


func NewRecord(in []string) (rec Record, err error){
	rec.Row,err = strconv.Atoi(in[0])
	if err != nil {
		return
	}
	rec.Height,err = strconv.ParseFloat(in[1],64)
	if err != nil {
		return
	}
	rec.Weight,err = strconv.ParseFloat(in[2],64)
	return 
	
}

func Parse(input []string) Record {
	rec, err := NewRecord(input)
	if err != nil {
		panic(err)
	}
	return rec
}

func Convert(r Record) Record {
	r.Height = 2.54 * r.Height
	r.Weight = 0.45 * r.Weight
	return r
}

func Encode(r Record) []byte {
	data,err:=json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return data
}


func main() {
	f,err := os.Open("sample.csv")
	if err != nil {
		panic(err)
	}
	input := csv.NewReader(f)
	//SynchronousPipeline(input)
	AsynchronousPipeline(input)
}


