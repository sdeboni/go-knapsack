// iteration 3 is a refinement of iteration 2 inspired by nmortier12's solution which is a further optimization possible by observing that only the maximum value is returned, and knowing which items were included in the total is not required.
package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {

  //dynamic programming solution
  //source: https://en.wikipedia.org/wiki/Knapsack_problem
  val := make([]int, maximumWeight+1)

  var max = func(a, b int) int {
    if a >= b {
      return a
    }
    return b
  }

  for _, item := range items {
    for w := maximumWeight; w >= item.Weight; w-- {
      val[w] = max(val[w], val[w-item.Weight] + item.Value)
    } 
  }

  return val[maximumWeight]
}
