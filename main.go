package pairs

import (
	"sync"
	"fmt"
)

type StructForPairs struct {
	values []int64
	sum    int64
}

var mutex = &sync.Mutex{}
var wg sync.WaitGroup

var low_index = new(int)
var high_index = new(int)

func GetPairsThatMatchesSum(pairs StructForPairs) (bool, []int64) {
	found := make(chan bool, 1)
	values := new([]int64)
	values_found := new(bool)
	values_until_sum := new([]int64)
	*values_found = false
	*low_index = 0
	*high_index = len(pairs.values) - 1
	wg.Add(2)
	go searchFromLow(found, values, values_found, pairs, values_until_sum)
	go searchFromHigh(found, values, values_found, pairs, values_until_sum)
	wg.Wait()
	fmt.Printf("is found? %v ", *values_found)
	fmt.Printf("values %v ", *values)

	return *values_found, *values
}
func searchFromLow(found chan bool, values *[]int64, values_found *bool, pairs StructForPairs, values_until_sum *[]int64) {
	for low_key := 0; low_key < len(pairs.values); low_key++ {
		select {
		case <-found:
			wg.Done()
			return
		default:
			*low_index = low_key
			if high_index != nil {
				if *high_index <= *low_index {
					wg.Done()
					return
				}
			}
			current_diff := int64(pairs.sum - pairs.values[low_key])
			for _, current_value_until_sum := range *values_until_sum {
				if pairs.values[low_key] == current_value_until_sum {
					mutex.Lock()
					*values = append(*values, current_diff)
					*values = append(*values, pairs.values[low_key])
					*values_found = true
					mutex.Unlock()

					found <- true
					wg.Done()
					return
				}
			}
			mutex.Lock()
			*values_until_sum = append(*values_until_sum, current_diff)

			mutex.Unlock()
		}

	}
	wg.Done()
}
func searchFromHigh(found chan bool, values *[]int64, values_found *bool, pairs StructForPairs, values_until_sum *[]int64) {
	for high_key := len(pairs.values) - 1; high_key >= 0; high_key-- {
		select {
		case <-found:
			wg.Done()
			return
		default:
			*high_index = high_key
			if low_index != nil {
				if *low_index > *high_index {
					wg.Done()
					return
				}
			}
			current_diff := int64(pairs.sum - pairs.values[high_key])
			for _, current_value_until_sum := range *values_until_sum {
				if pairs.values[high_key] == current_value_until_sum {
					mutex.Lock()
					*values = append(*values, pairs.values[high_key])
					*values = append(*values, current_diff)
					*values_found = true
					mutex.Unlock()

					found <- true
					wg.Done()
					return
				}
			}

			mutex.Lock()
			*values_until_sum = append(*values_until_sum, current_diff)
			mutex.Unlock()
		}
	}

	wg.Done()
}
