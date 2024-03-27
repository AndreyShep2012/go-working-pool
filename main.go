package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

var testData = generate(5000)

func main() {
	res := ProcessDataMutex(testData)
	res.print()
}

func ProcessDataMutex(numbers []int) DataList {
	var output DataList
	var wg sync.WaitGroup
	var mq sync.Mutex

	for _, n := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			r := &Data{
				number: n,
				square: n * n,
			}

			mq.Lock()
			output = append(output, r)
			mq.Unlock()
		}(n)
	}

	wg.Wait()

	return output
}

func ProcessDataWorkerPool(numbers []int) DataList {
	var output DataList
	outputCh := make(chan *Data)
	workerCount := 5

	var wg sync.WaitGroup

	workerCh := publishData(numbers)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for v := range workerCh {
				outputCh <- &Data{
					number: v,
					square: v * v,
				}
			}
		}()
	}

	//close the channel after finishing all go routines
	go func() {
		wg.Wait()
		close(outputCh)
	}()

	// receive output
	for d := range outputCh {
		output = append(output, d)
	}

	return output
}

func publishData(numbers []int) <-chan int {
	workerCh := make(chan int)
	go func() {
		defer close(workerCh)
		for _, v := range numbers {
			workerCh <- v
		}
	}()
	return workerCh
}

type Data struct {
	number int
	square int
}

type DataList []*Data

func (d DataList) print() {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Total number: %d\n", len(d)))
	for _, v := range d {
		sb.WriteString(fmt.Sprintf("{n: %d s: %d}, ", v.number, v.square))
	}

	log.Println(sb.String())
}

func ProcessData(numbers []int) DataList {
	var outputs DataList
	outputCh := make(chan *Data)

	var wg sync.WaitGroup
	for _, v := range numbers {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()
			outputCh <- &Data{
				number: v,
				square: v * v,
			}
		}(v)
	}

	//close the channel after finishing all go routines
	go func() {
		wg.Wait()
		close(outputCh)
	}()

	// receive output
	for d := range outputCh {
		outputs = append(outputs, d)
	}

	return outputs
}

func generate(len int) []int {
	res := make([]int, len)
	for i := 0; i < len; i++ {
		res[i] = i
	}

	return res
}
