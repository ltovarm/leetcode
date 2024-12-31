# Best Time to Buy and Sell Stock

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

```C++
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
```

Example 2:

```C++
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
```

Constraints:

```C++
1 <= prices.length <= 105
0 <= prices[i] <= 104
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

### Approach
<!-- Describe your approach to solving the problem. -->

### Complexity

* Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
Run time <=  0 ms Beats 100.00 % of users with C++

* Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
Memory = 97.24 MB Beat 52.48 % of user with C++

### Code

```C++
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int maxDiff = 0;
        int minSoFar = prices.at(0);
        for (auto value : prices)
        {
            if (value < minSoFar)
            {
                minSoFar = value;
            }
            if (value - minSoFar > maxDiff)
            {
                maxDiff = value - minSoFar;
            }
        }
        return maxDiff;
    }
};
```
