package main
import fmt "fmt"
import rand "rand"
import math "math"

func mergesort(a []int) []int {
	 if len(a) == 1 {
	 	return a
     } else if len(a) == 2 {
	    if a[0] < a[1] {
		   return a[:]
		} else {
		   return [2]int{a[1], a[0]}[:]
		}
     }
	 array_size_f := float64(len(a))
	 left_end_ndx := int(math.Floor(array_size_f/2.0))
	 right_begin_ndx := int(math.Floor(array_size_f/2.0))
	 left := mergesort(a[:left_end_ndx])
	 right := mergesort(a[right_begin_ndx:])
	 return merge(left[:], right[:])
}

func merge(left_array []int, right_array []int) []int {
	 combined := make([]int, len(left_array) + len(right_array))
	 left_ndx := 0
	 right_ndx := 0
	 for ndx := 0; ndx < len(combined); ndx++ {
	 	left_num := 9999
		right_num := 9999
		if left_ndx < len(left_array) {
		   left_num = left_array[left_ndx]
		}
		if right_ndx < len(right_array) {
		   right_num = right_array[right_ndx]
        }

		if left_num < right_num {
		   combined[ndx] = left_num
		   left_ndx++
		} else {
		   combined[ndx] = right_num
		   right_ndx++
		}
	 }
	 return combined
}

func main() {
	 rand.Seed(42)
	 array_size := 100000
	 arrayOfInt := make([]int, array_size)
	 for i := 0; i < len(arrayOfInt); i++ {
	 	 arrayOfInt[i] = rand.Intn(100)
	 }
	 fmt.Printf("Unsorted:\n")
     for i,v := range arrayOfInt {
         fmt.Printf("[%d] = [%d]\n", i, v)
     }
     arrayOfSortedInt := mergesort(arrayOfInt[:])
     fmt.Printf("Sorted:\n")
     for i,v := range arrayOfSortedInt {
         fmt.Printf("[%d] = [%d]\n", i, v)
     }
}