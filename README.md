# golang_missing_number

Given an array `nums` containing `n` distinct numbers in the range `[0, n]`
, return *the only number in the range that is missing from the array.*

## Examples

**Example 1:**

```
Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

```

**Example 2:**

```
Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.

```

**Example 3:**

```
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.

```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 104`
- `0 <= nums[i] <= n`
- All the numbers of `nums` are **unique**.

**Follow up:** Could you implement a solution using only `O(1)` extra space complexity and `O(n)` runtime complexity?

## 解析

給定一個整數陣列 nums 包含 n 個不同介於 [0, n] 的非負整數

要求找出 [0, n] 中沒出現在 nums 的數字

要求使用 空間複雜度 O(1), 時間複雜度 O(n) 的演算法

1 透過 sum 的方式來做

因為是 [0, n] 的整數 所以 sum 固定是 n*(n+1)/2

只要把 nums 所有數值都減去就知道少的是哪一個

這樣做只要把所有數值 loop 一便 所以是時間複雜度是 O(n)

2 透過 XOR 的特性來做

已知對於整數 XOR 有以下特性

0 XOR a = a

a XOR 0 = a

b XOR a  XOR a = a XOR b XOR a = a XOR a XOR b = b

所以知道 對任何整數在 XOR 運算中出現兩次 就會剩餘其他的數值

所以 對於 0.. n 把 nums 內所有數值做 XOR

這樣因為只會有一個數值沒有出現兩次

所以最後的值就是剩下的值

因為 0..n 有n+ 1 個數值, nums 有 n 個數值

所有總共是 2n+1個數值

時間複雜度是 O(n)

## 程式碼
```go
package sol

func missingNumberXOR(nums []int) int {
	n := len(nums)
	res := 0
	for pos := 0; pos < n; pos++ {
		res ^= ((pos + 1) ^ nums[pos])
	}
	return res
}
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	res := sum
	for pos := 0; pos < n; pos++ {
		res -= nums[pos]
	}
	return res
}

```
## 困難點

1. 思考出利用 XOR 特性來找出沒出現的值
2. 利用 sum 來找出沒出現的值

## Solve Point

- [x]  初始化 n := len(nums) , sum := n*(n+1)/2
- [x]  把 sum 逐一減去所有 nums 的值，剩下來的值就是結果