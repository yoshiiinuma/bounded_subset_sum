
package subset_sum

import (
  "fmt"
)

const DEBUG = true

func showDP(DP [][]int) {
  for _, dpi := range(DP) {
    r := make([]string, len(dpi))
    for j, e :=  range(dpi) {
      r[j] = fmt.Sprintf("%2d", e)
    }
    fmt.Println(r)
  }
}

func SubsetSum(A []int, M []int, K int) bool {
  if DEBUG {
    fmt.Println("=======================================")
    fmt.Printf("K: %d, A: %v, M: %v\n", K, A, M)
    fmt.Println("=======================================")
  }
  N := len(A)
  var DP = make([][]int, N + 1);
  for i := range DP {
    DP[i] = make([]int, K + 1);
  }
  for j := 1; j <= K; j++ {
    DP[0][j] = -1
  }
  r := subsetSum(DP, A, M, K)
  if DEBUG { showDP(DP) }
  return r
}

func subsetSum(DP [][]int, A []int, M []int, K int) bool {
  N := len(A);
  var a int
  for i := 1; i <= N; i++ {
    a = A[i - 1]
    for j := 0; j <= K; j++ {
      if j < a {
        if DP[i-1][j] >= 0 {
          DP[i][j] = M[i-1]
        } else {
          DP[i][j] = DP[i-1][j]
        }
      } else {
        if DP[i-1][j] >= 0 {
          DP[i][j] = M[i-1]
        } else if DP[i][j-a] >= 0 {
          DP[i][j] = DP[i][j-a] - 1
        } else {
          DP[i][j] = DP[i-1][j]
        }
      }
    }
  }
  return DP[N][K] >= 0;
}


