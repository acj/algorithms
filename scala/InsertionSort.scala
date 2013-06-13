package algorithms

object InsertionSort {
  def sort(list: List[Int]): List[Int] = {
    def insert(e: Int, insertList: List[Int]) = {
      def insertPrime(insertPrimeList: List[Int]): List[Int] = insertPrimeList match {
        case Nil => List(e)
        case x :: xs if (e < x) => e :: x :: xs
        case x :: xs => x :: insertPrime(xs)
      }
	    
      insertPrime(insertList)
    }
	    
    def insertionSortPrime(sortedList: List[Int], unsortedList: List[Int]): List[Int] = unsortedList match {
      case Nil => sortedList
      case x :: xs => insertionSortPrime(insert(x, sortedList), xs)
    }
	
    insertionSortPrime(List[Int](), list)
  }
}
