import Cocoa

// http://stackoverflow.com/a/24029847/357055
extension Array {
    mutating func shuffle() {
        for i in 0..<(count - 1) {
            let j = Int(arc4random_uniform(UInt32(count - i))) + i
            swap(&self[i], &self[j])
        }
    }
}

func merge(inout array: [Int], p: Int, q: Int, r: Int) {
    assert(p >= 0 && q <= r && r < array.count, "Index out of bounds")
    
    let n1 = q - p + 1
    let n2 = r - q
    
    var slice1 = array[p...q]
    var slice2 = array[q+1...r]
    
    for i in p...r {
        if slice1.count == 0 && slice2.count == 0 {
            return
        } else if slice1.count == 0 {
            array[i...r] = slice2
            return
        } else if slice2.count == 0 {
            array[i...r] = slice1
            return
        }
        
        let num1 = slice1[0]
        let num2 = slice2[0]
        
        if num1 < num2 {
            array[i] = num1
            slice1.removeAtIndex(0)
        } else {
            array[i] = num2
            slice2.removeAtIndex(0)
        }
    }
}

func mergesort(inout array: [Int], p: Int, r: Int) {
    if p < r {
        let q = Int(floor(Float((p + r) / 2)))
        mergesort(&array, p, q)
        mergesort(&array, q+1, r)
        merge(&array, p, q, r)
    }
}

func mergesort(inout array: [Int]) {
    mergesort(&array, 0, array.count - 1)
}

let evenReferenceArray = [Int](1...20)
var evenMutableArray = [Int](evenReferenceArray)
evenMutableArray.shuffle()
mergesort(&evenMutableArray)

let oddReferenceArray = [Int](1...21)
var oddMutableArray = [Int](oddReferenceArray)
oddMutableArray.shuffle()
mergesort(&oddMutableArray)

assert(evenMutableArray == evenReferenceArray, "Array should be sorted")
assert(oddMutableArray == oddReferenceArray, "Array should be sorted")