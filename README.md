# Bounded Subset Sum

Assume given a set A [a1, a2, a3,..., an], each element can be selected multiple times with the upper limit M [m1, m2, m3,..., mn]. Determine if there is a subset of the given set with sum equal to K.

* 1 <= n <= 100
* 1 <= ai, mi <= 100000
* 1 <= K <= 100000

## Examples

```
Input:

  A = { 3, 5, 8 }
  M = { 3, 2, 2 }
  K = 17

Output:

   True

   There is a subset { 3, 8 } with sum 17
   17 = 3 * 3 + 8
```

```
Input:

  A = { 3, 5, 8 }
  M = { 3, 2, 2 }
  K = 12

Output:

   False

   There is no such a subset with sum 12
```


## Approaches

Suppose j represents the sum that can be obtained from a subset of A[1..i], and DP[i][j] represents the number of A[i] that is not used to compose the sumof subset A[1..i]. The recurrence relation can be defined as:

```
Set DP[0][0] = 0
    DP[0][j] = -1 (1 <= j <= K)

i) When j > A[i]:

   a) when DP[i-1] >= 0

      DP[i][j] = M[i]

   b) Otherwise

      DP[i][j] = DP[i-1][j]

ii) When j <= A[i]:

   a) when DP[i-1] >= 0

      DP[i][j] = M[i]

   b) when DP[i-A[i]] >= 0

      DP[i][j] = DP[i][j - A[i]] - 1

   c) Otherwise

      DP[i][j] = DP[i-1][j]
```

## Complexity Analysis

* Time Complexity: O(N*K)
* Space Complexity: O(N*K)
