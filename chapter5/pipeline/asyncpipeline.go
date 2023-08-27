package main

import (
	"encoding/csv"
	"fmt"
	"io"
)


func PipelineStage[IN any, OUT any](input <-chan IN, output chan<- OUT, process func(IN) OUT) {
	defer close(output)
	for data:=range input {
		output <- process(data)
	}
}

func AsynchronousPipeline(input *csv.Reader){

	parseInputCh := make(chan []string)
	convertInputCh := make(chan Record)
	encodeInputCh := make(chan Record)
	outputCh := make(chan []byte)
	done := make(chan struct{})


	go PipelineStage(parseInputCh, convertInputCh, Parse)
	go PipelineStage(convertInputCh, encodeInputCh, Convert)
	go PipelineStage(encodeInputCh, outputCh, Encode)

	go func() {
		for data := range outputCh {
			fmt.Println(string(data))
		}
		close(done)
	}()

	input.Read()
	for {
		rec, err := input.Read()
		if err == io.EOF {
			close(parseInputCh)
			break
		}
		if err != nil {
			panic(err)
		}
		parseInputCh <- rec
	}

	<-done

}