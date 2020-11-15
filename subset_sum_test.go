
package subset_sum

import (
  "testing"
  . "github.com/stretchr/testify/assert"
)

func TestSubsetSum(t *testing.T) {
  var A []int
  var M []int
  var K int

  A = []int{ 3, 5, 8 }
  M = []int{ 3, 2, 2 }
  K = 17
  Equal(t, true, SubsetSum(A, M, K))

  A = []int{ 3, 5, 8 }
  M = []int{ 3, 2, 2 }
  K = 11
  Equal(t, true, SubsetSum(A, M, K))

  A = []int{ 3, 5, 8 }
  M = []int{ 3, 2, 2 }
  K = 10
  Equal(t, true, SubsetSum(A, M, K))

  A = []int{ 3, 5, 8 }
  M = []int{ 3, 2, 2 }
  K = 12
  Equal(t, false, SubsetSum(A, M, K))

  A = []int{ 3, 5 }
  M = []int{ 3, 2 }
  K = 9
  Equal(t, true, SubsetSum(A, M, K))

  A = []int{ 3, 5 }
  M = []int{ 3, 2 }
  K = 7
  Equal(t, false, SubsetSum(A, M, K))
}
