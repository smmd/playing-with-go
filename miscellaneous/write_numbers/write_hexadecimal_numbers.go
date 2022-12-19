// source: Chio Code
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("concurrent_hexadecimals.txt")
	if err != nil {
		panic(err)
	}

	outCh := make(chan string)
	doneWrite := make(chan struct{})
	// Writer
	go func() {
		for s := range outCh {
			_, err := file.WriteString(s)
			if err != nil {
				panic(err)
			}
		}
		doneWrite <- struct{}{}
	}()

	numGoRoutines := 10
	doneCh := make(chan struct{})

	final := 16777215
	for i := 0; i <= final; i = i + (final / numGoRoutines) + 1 {
		step := i + (final / numGoRoutines)
		if step > final {
			step = final
		}
		fmt.Printf("executing %d %d\n", i, step)
		go numCalculation(i, step, outCh, doneCh)
	}

	doneNum := 0
	for doneNum < numGoRoutines {
		<-doneCh
		fmt.Println("finish step")
		doneNum++
	}
	close(outCh)
	<-doneWrite
	fmt.Println("Finished!")
}

func numCalculation(start, end int, resultCh chan string, doneCh chan struct{}) {
	var sBuilder strings.Builder
	for i := start; i <= end; i++ {
		fmt.Fprint(&sBuilder, fmt.Sprintf("%06x\n", i))
	}
	resultCh <- sBuilder.String()
	doneCh <- struct{}{}
}
