package main

import (
	"DesignPatterns/search"
	"DesignPatterns/sort"
	"fmt"
)

func main() {

	// circuit breaker pattern
	//cb := stability.NewCircuitBreaker(5*time.Second, 3)
	//for i := 0; i < 20; i++ {
	//	fmt.Println("Request : ", i)
	//	err := cb.SendRequest(stability.RunAPI)
	//	if err != nil {
	//		fmt.Println("Error returned : ", err)
	//	} else {
	//		fmt.Println("Successful attempt")
	//	}
	//
	//	time.Sleep(1 * time.Second)
	//}

	// sequence sage - orchestration
	//s := &saga.Saga{
	//	Sequence: []*saga.Step{
	//		{
	//			Name:         "Add order",
	//			Action:       saga.AddOrder,
	//			Compensation: saga.CancelOrder,
	//		},
	//		{
	//			Name:         "Do earmark",
	//			Action:       saga.DoEarmark,
	//			Compensation: saga.ReleaseEarmark,
	//		},
	//		{
	//			Name:         "Attempt matching",
	//			Action:       saga.PerformMatching,
	//			Compensation: saga.DoRollback,
	//		},
	//	},
	//}
	//
	//s.Execute()

	//cb := stability.NewCircuitBreakerBackOff(5*time.Second, 3)
	//for i := 0; i < 10; i++ {
	//	fmt.Println("============ Sending request ", i)
	//	_ = stability.ProcessRequestWithBackoff(cb, 5, 1*time.Second)
	//}

	// semaphore
	//semaphore := concepts.NewSemaphore(3)
	//wg := sync.WaitGroup{}
	//for i := 0; i < 10; i++ {
	//	fmt.Println("Trying to acquire semaphore ", i)
	//	semaphore.Acquire()
	//	wg.Add(1)
	//	fmt.Println("Acquired the semaphore ", i)
	//
	//	go func(id int) {
	//		defer semaphore.Release()
	//		defer wg.Done()
	//
	//		time.Sleep(2 * time.Second)
	//		fmt.Println("Done with this thread ", id)
	//	}(i)
	//}
	//
	//wg.Wait()

	// producer consumer
	//fmt.Println("\n=========== Producer Consumer")
	//pc := concepts.NewProducerConsumer(5)
	//wg := sync.WaitGroup{}
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	pc.Produce()
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	pc.Consume()
	//}()
	//
	//wg.Wait()

	input := []int{7, 3, 9, 12, 11}
	sort.PrintArray(input)
	sort.PerformBubbleSort(input)
	fmt.Println("==== Result : ", input)
	input = []int{7, 3, 9, 12, 11}
	sort.PerformSelectionSort(input)
	fmt.Println("==== Result : ", input)
	fmt.Println()

	input = []int{15, 12, 9, 3, 1}
	sort.PrintArray(input)
	sort.PerformBubbleSort(input)
	fmt.Println("==== Result : ", input)
	input = []int{15, 12, 9, 3, 1}
	sort.PerformSelectionSort(input)
	fmt.Println("==== Result : ", input)
	fmt.Println()

	input = []int{7, 12, 9, 11, 3}
	sort.PerformSelectionSort(input)
	fmt.Println("==== Result : ", input)
	fmt.Println()

	input = []int{64, 34, 25, 12, 22, 11, 90, 5}
	sort.PerformInsertionSort(input)
	fmt.Println("==== Result : ", input)
	fmt.Println()

	input = []int{64, 34, 25, 12, 22, 11, 90, 5}
	input = sort.PerformQuickSort(input)
	fmt.Println("==== Result : ", input)
	fmt.Println()

	input = []int{64, 34, 25, 12, 22, 11, 90, 5}
	input = sort.MergeSort(input)
	fmt.Println("==== Result : ", input)
	fmt.Println()

	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7
	index := search.BinarySearch(arr, target)
	fmt.Println("index:", index)

	index = search.BinarySearch(arr, 10)
	fmt.Println("index:", index)
}
