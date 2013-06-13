package main
import fmt "fmt"
import rand "rand"

func insertion(a []int) []int {
	 for ndx := 2; ndx < len(a); ndx++ {
		 val := a[ndx]
		 insert_ndx := 0
		 for insert_ndx = ndx-1; insert_ndx >= 0 && a[insert_ndx] > val; insert_ndx-- {
		 	 a[insert_ndx+1] = a[insert_ndx]
		 }
		 a[insert_ndx+1] = val
	 }
	 return a
}

func main() {
	 array_size := 100000

	 arrayOfInt := make([]int, array_size)
	 for i := 0; i < len(arrayOfInt); i++ {
	 	 arrayOfInt[i] = rand.Intn(100)
	 }
	 fmt.Printf("Unsorted:\n")
	 for i,v := range arrayOfInt {
	 	 fmt.Printf("[%d] = [%d]\n", i, v)
	 }
	 insertion(arrayOfInt[:])
	 fmt.Printf("Sorted:\n")
	 for i,v := range arrayOfInt {
	 	 fmt.Printf("[%d] = [%d]\n", i, v)
	 }
}