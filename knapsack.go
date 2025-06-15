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
  var m = make([][]int, len(items)+1)
  for i := 0; i < len(m); i++ {
    m[i] = make([]int, maximumWeight+1)
  } 

  var max = func(a, b int) int {
    if a >= b {
      return a
    }
    return b
  }

  for i := 1; i < len(m); i++ {
    for j := 1; j <= maximumWeight; j++ {
      if items[i-1].Weight > j {
        m[i][j] = m[i-1][j]
      } else {
        m[i][j] = max(m[i-1][j], m[i-1][j-items[i-1].Weight]+ items[i-1].Value)
      }
    }
  }

  return m[len(items)][maximumWeight]
}
