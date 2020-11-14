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


## Approaches

Suppose j represents the sum that can be obtained from a subset of A[1..i], and DP[i][j] represents the number of A[i] that is not used to compose the sumof subset A[1..i]. The recurrence relation can be defined as:

```
i) When j > A[i]:

   DP[i][j] = DP[i-1][j]

ii) When j <= A[i]:

   a) When j == A[i]

      DP[i][j] = M[i] 

   b) When j < A[i]

      DP[i][j] = DP[i][j - A[i]] - 1

   c) Otherwise

      DP[i][j] = DP[i-1][j]
```

Because if there exists a subset of sum j with 1 to i-1th elements of A, both DP[i][j] and DP[i][j + A[i-1]] are certainly true.

If current element A[i-1] has value greater than j, the value cannot be added to the subset because the sum will be greater than j.
Also, if the sum of the current subset is greater than j, no further calculation is needed because it does not satfifies the problem. 

DP[i][0] is always true because the sum of the subset {} is always 0.

## Complexity Analysis

* Time Complexity: O(N*K)
* Space Complexity: O(N*K)
