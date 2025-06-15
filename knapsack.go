// This is a brute force solution: all possible combinations are evaluated
// The next iteration will focus on optimizing performance

package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
  var best int

  for i := 0; i < len(items); i++ {
    for _, combination := range choose(i, items) {
      weight, value := total(combination)
      if weight <= maximumWeight && value > best {
        best = value
      }
    } 
  }
  return best
}

func total(items []Item) (weight int, value int) {
  for _, item := range items {
    weight += item.Weight
    value += item.Value
  }
  return
}

func choose[T any](n int, items []T) [][]T {
  var result [][]T

  N := len(items)
   
  a := make([]int, n)
  for i := 0; i < n; i++ {
    a[i] = i
  }
  result = add(result, a, items)
  
  i := n-1
  for i >= 0 {
    for a[i] < i + N - n {
      a[i] = a[i]+1
      i++
      for i < n {
        a[i] = a[i-1] + 1
        i++
      }
      result = add(result, a, items)
      i--
    }
    i--
  } 
  return result
}
func add[T any](a [][]T, b []int, v []T) [][]T {
  c := make([]T, len(b))
  for i := 0; i < len(b); i++ {
    c[i] = v[b[i]]
  }
  return append(a, c)
}


