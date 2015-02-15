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

func insertionSort(inout array: [Int]) {
    for index in 1..<array.count {
        let key = array[index]
        var placeIndex = index - 1
        while placeIndex >= 0 && array[placeIndex] > key {
            array[placeIndex + 1] = array[placeIndex]
            placeIndex--
        }
        
        array[placeIndex + 1] = key
    }
}

let referenceArray = [Int](1...20)
var mutableArray = [Int](referenceArray)
mutableArray.shuffle()


insertionSort(&mutableArray)


assert(mutableArray == referenceArray, "Array should be sorted")