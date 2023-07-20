## Dynamic Programing

Sure, I can provide a brief description and an example for each of the 100 dynamic programming problems I listed earlier. Here are the first 50 problems:

1. Fibonacci sequence: Given an integer n, find the nth Fibonacci number. The Fibonacci sequence is defined as follows: F0 = 0, F1 = 1, and Fn = Fn-1 + Fn-2 for n > 1. Example: If n = 6, then the output should be 8 (since the 6th Fibonacci number is 8).

2. Longest Increasing Subsequence: Given an unsorted array of integers, find the length of the longest increasing subsequence. Example: If the input array is [10, 9, 2, 5, 3, 7, 101, 18], then the longest increasing subsequence is [2, 3, 7, 101], and the length of this subsequence is 4.

3. Longest Common Subsequence: Given two strings s and t, find the length of their longest common subsequence. Example: If s = "abcde" and t = "ace", then the longest common subsequence is "ace", and the length of this subsequence is 3.

4. Knapsack Problem: Given a list of items with weights and values, and a maximum weight capacity, find the maximum value that can be obtained by selecting a subset of the items such that the total weight is less than or equal to the maximum weight capacity. Example: If the items are [(w=1, v=6), (w=2, v=10), (w=3, v=12)] and the maximum weight capacity is 5, then the maximum value that can be obtained is 22 (by selecting items 1 and 3).

5. Edit Distance: Given two strings s and t, find the minimum number of operations required to convert s into t, where an operation is defined as either inserting a character, deleting a character, or replacing a character. Example: If s = "horse" and t = "ros", then the minimum number of operations required is 3 (by deleting the "h" and the "e" from s, and replacing the "r" with "o").

6. Maximum Subarray Problem: Given an array of integers, find the contiguous subarray with the largest sum. Example: If the input array is [-2, 1, -3, 4, -1, 2, 1, -5, 4], then the maximum subarray is [4, -1, 2, 1], and the sum of this subarray is 6.

7. Coin Change Problem: Given a list of coin denominations and a target amount, find the minimum number of coins required to make the target amount, assuming an infinite supply of each coin denomination. Example: If the coin denominations are [1, 2, 5] and the target amount is 11, then the minimum number of coins required is 3 (by using two 5-cent coins and one 1-cent coin).

8. Rod Cutting Problem: Given a rod of length n and a list of prices for each possible cut, find the maximum profit that can be obtained by cutting the rod into pieces and selling the pieces. Example: If the rod length is 8 and the prices for each cut are [1, 5, 8, 9, 10, 17, 17, 20], then the maximum profit that can be obtained is 22 (by cutting the rod into two pieces of length 2 and 6, and selling them for prices of 5 and 17).

9. Binomial Coefficient: Given two integers n and k, find the binomial coefficient C(n, k), which represents the number of ways to choose k items from a set of n items. Example: If n = 5 and k = 2, then C(n, k) = 10 (since there are 10 ways to choose 2 items from a set of 5 items).

10. Minimum Cost Path: Given a two-dimensional grid of integers, find the minimum cost path from the top-left corner to the bottom-right corner, where the cost of a path is defined as the sum of the integers along the path. Example: If the input grid is [[1, 3, 1], [1, 5, 1], [4, 2, 1]], then the minimum cost path is [1, 3, 1, 1, 1], and the cost of this path is 7.

11. Partition Equal Subset Sum: Given an array of integers, determine whether it can be partitioned into two subsets such that the sum of the elements in each subset is equal. Example: If the input array is [1, 5, 11, 5], then the answer is true, sincethe array can be partitioned into [1, 5, 5] and [11], which both have a sum of 11.

12. Word Break Problem: Given a string s and a dictionary of words, determine whether s can be segmented into a space-separated sequence of dictionary words. Example: If the input string is "leetcode" and the dictionary is ["leet", "code"], then the answer is true, since "leetcode" can be segmented into "leet" and "code".

13. Unique Paths: Given a two-dimensional grid of size m x n, find the number of unique paths from the top-left corner to the bottom-right corner, where a path can only move down or right. Example: If the grid size is 3 x 7, then the number of unique paths is 28.

14. Climbing Stairs: Given an integer n, find the number of distinct ways to climb n stairs, where each step can be either 1 or 2. Example: If n = 4, then the number of distinct ways to climb the stairs is 5 (by taking steps of size 1-1-1-1, 1-1-2, 1-2-1, 2-1-1, or 2-2).

15. House Robber: Given an array of non-negative integers representing the amount of money at each house, determine the maximum amount of money you can rob without robbing adjacent houses. Example: If the input array is [1, 2, 3, 1], then the maximum amount of money that can be robbed is 4 (by robbing the first and third houses).

16. Regular Expression Matching: Given a string s and a pattern p, determine whether s matches p, where p can contain the following wildcards: "." (matches any single character) and "*" (matches zero or more of the preceding character). Example: If s = "aa" and p = "a*", then the answer is true.

17. Palindrome Partitioning: Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible partitions. Example: If s = "aab", then the possible partitions are [["a", "a", "b"], ["aa", "b"]].

18. Longest Palindromic Subsequence: Given a string s, find the length of the longest palindromic subsequence in s. Example: If s = "bbbab", then the longest palindromic subsequence is "bbbb", and the length of this subsequence is 4.

19. Maximum Product Subarray: Given an array of integers, find the contiguous subarray with the largest product. Example: If the input array is [2, 3, -2, 4], then the maximum product subarray is [2, 3], and the product of this subarray is 6.

20. Wildcard Matching: Given a string s and a pattern p, determine whether s matches p, where p can contain the following wildcards: "?" (matches any single character) and "*" (matches zero or more of any characters). Example: If s = "adceb" and p = "*a*b", then the answer is true.

21. Maximum Length of Pair Chain: Given a list of pairs, find the length of the longest chain of pairs where each pair (a, b) is such that a < b. Example: If the input pairs are [(1,2), (2,3), (3,4)], then the length of the longest chain is 3 (by selecting all three pairs).

22. Decode Ways: Given a string s containing digits from '0'-'9', determine the number of ways to decode it into letters. Example: If s = "226", then there are three ways to decode it: "BZ", "VF", or "BBF".

23. Unique Binary Search Trees: Given an integer n, determine the number of unique binary search trees that can be formed using the numbers 1 to n as nodes. Example: If n = 3, then there are 5 unique binary search trees: 1-2-3, 1-3-2, 2-1-3, 3-1-2, and 3-2-1.

24. Palindrome Partitioning II: Given a string s, partition s such that every substring of the partition is a palindrome. Return the minimum cuts needed for a palindrome partitioning of s. Example: If s = "aab", then the minimum cuts needed for a palindrome partitioning are 1 (by partitioning into ["aa", "b"]).

25. Jump Game: Given an array of non-negative integers, determine whether you can reach the last index if you can only jump a maximum of k steps from each index. Example: If the input array is [2, 3, 1, 1, 4] and k = 2, then the answer is true.

26. Unique Paths II: Given a two-dimensional grid with obstacles, find the number of unique paths from the top-left corner to the bottom-right corner, where a path can only move down or right, and obstacles are represented by 1's. Example: If the input grid is [[0,0,0], [0,1,0], [0,0,0]], then the number of unique paths is 2.

27. Minimum Path Sum: Given a two-dimensional grid of integers, find the minimum path sum from the top-left corner to the bottom-right corner, where a path can only move down or right. Example: If the input grid is [[1,3,1], [1,5,1], [4,2,1]], then the minimum path sum is 7.

28. Word Break II: Given a string s and a dictionary of words, return all possible space-separated word breakings of s using the words from the dictionary. Example: If s = "catsanddog" and the dictionary is ["cat", "cats", "and", "sand", "dog"], then the possible word breakings are ["cats and dog", "cat sand dog"].

29. Maximum Subarray Sum with One Deletion: Given an array of integers, find the maximum sum of any contiguous subarray of length at most n, where one element can be deleted from the subarray. Example: If the input array is [1,-2,0,3], then the maximum sum of any contiguous subarray of length at most 3 with one element deleted is 4 (by selecting the subarray [1,0,3]).

30. Partition to K Equal Sum Subsets: Given an array of integers and an integer k, determine whether it is possible to partition the array into k non-empty subsets with equal sums. Example: If the input array is [4, 3, 2, 3, 5, 2, 1] and k = 4, then the answer is true, since the array can be partitioned into [4], [3, 1], [2, 2], and [3, 5].

31. Minimum Size Subarray Sum: Given an array of integers and a target sum, find the minimum size of a contiguous subarray of the array that sums up to the target sum. Example: If the input array is [2,3,1,2,4,3] and the target sum is 7, then the minimum size of a contiguous subarray that sums up to 7 is 2 (by selecting the subarray [4,3]).

32. Longest Increasing Path in a Matrix: Given a two-dimensional matrix of integers, find the length of the longest increasing path. Example: If the input matrix is [[9,9,4], [6,6,8], [2,1,1]], then the length of the longest increasing path is 4 (by selecting the path [9, 6, 2, 1]).

33. Counting Bits: Given a non-negative integer n, return an array of integers representing the number of 1's in the binary representation of every numberfrom 0 to n. Example: If n = 5, then the array of counts is [0, 1, 1, 2, 1, 2].

34. Find Median from Data Stream: Design a data structure that supports adding and finding the median of elements in a continuous data stream. Example: If the stream of elements is [2, 1, 3, 4, 5], then the median at each step is [2, 1.5, 2, 2.5, 3].

35. Longest Palindromic Subsequence: Given a string s, find the length of the longest palindromic subsequence of s. Example: If s = "bbbab", then the length of the longest palindromic subsequence is 4 (by selecting the subsequence "bbbb").

36. Serialize and Deserialize Binary Tree: Design an algorithm to serialize and deserialize a binary tree. Example: If the binary tree is [1,2,3,null,null,4,5], then the serialized string is "1,2,3,null,null,4,5" and the deserialized tree is:

       1
     /   \
    2     3
         / \
        4   5

37. Reconstruct Itinerary: Given a list of airline tickets represented by pairs of departure and arrival airports [from, to], reconstruct the itinerary in order. Example: If the input tickets are [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]], then the itinerary is ["JFK","ATL","JFK","SFO","ATL","SFO"].

38. Number of Islands: Given a two-dimensional grid of 0's and 1's, count the number of islands, where an island is a group of connected 1's (horizontally or vertically). Example: If the input grid is [[1,1,0,0,0], [1,1,0,0,0], [0,0,1,0,0], [0,0,0,1,1]], then the number of islands is 3.

39. Maximum Profit in Stock Market: Given an array of stock prices, find the maximum profit that can be obtained by buying and selling stocks on different days. Example: If the input array is [7,1,5,3,6,4], then the maximum profit is 5 (by buying on day 2 and selling on day 5).

40. Combination Sum: Given a set of candidate numbers (without duplicates) and a target sum, find all unique combinations of the candidate numbers that sum up to the target. Example: If the input candidates are [2,3,6,7] and the target sum is 7, then the possible combinations are [[2,2,3], [7]].

41. Longest Repeating Character Replacement: Given a string s and an integer k, find the length of the longest substring in s that can be replaced with any character for up to k times, such that the resulting substring contains only repeating characters. Example: If s = "ABAB" and k = 2, then the length of the longest substring that can be replaced with any character for up to 2 times is 4 (by replacing both "A"s with "B"s).

42. Course Schedule II: Given the total number of courses and a list of prerequisite pairs, return the order in which the courses can be taken. Example: If there are 4 courses and the prerequisite pairs are [[1,0],[2,0],[3,1],[3,2]], then one possible order is [0,1,2,3].

43. Maximum Points You Can Obtain from Cards: Given an array of integers representing the values of cards in a deck, and an integer k representing the number of cards you can draw, return the maximum score you can obtain. Example: If the input array is [1,2,3,4,5,6,1] and k = 3, then the maximum score you can obtain is 12 (by selecting the cards with values 6, 5, and 1).

44. Best Time to Buy and Sell Stock with Cooldown: Given an array of stock prices, find the maximum profit that can be obtained by buying and selling stocks on different days, with a cooldown period of one day between any buy and sell actions. Example: If the input array is [1,2,3,0,2], then the maximum profit is 3 (by buying on day 1 and selling on day 2, then buying on day 4 and selling on day 5).

45. Subarray Sum Equals K: Given an array of integers and an integer k, find the total number of continuous subarrays whose sum equals k. Example: If the input array is [1,1,1] and k = 2, then the total number of continuous subarrays whose sum equals k is 2 (by selecting the subarrays [1,1] and [1,1]).

46. Course Schedule III: Given the total number of courses and a list of course durations, find the maximum number of courses you can take such that the total course duration is less than or equal to a certain number. Example: If there are 5 courses with durations [5,4,3,2,1] and the maximum course duration is 10, then the maximum number of courses you can take is 5 (by taking all the courses in decreasing order of duration).

47. Minimum Cost to Hire K Workers: Given two arrays representing the quality and wage of workers, and an integer k representing the number of workers you can hire, find the minimum cost to hire k workers such that each worker is paid according to their quality, and the ratio of wage to quality is the same for all workers. Example: If the quality array is [10,20,5] and the wage array is [70,50,30] with k = 2, then the minimum cost to hire k workers is 105 (by hiring workers 0 and 2).

48. Unique Paths III: Given a two-dimensional grid with obstacles and starting and ending points, find the number of unique paths from the starting point to the ending point, where a path can only move up, down, left, or right, and must visit every non-obstacle square exactly once. Example: If the input grid is [[1,0,0,0], [0,0,0,0], [0,0,2,-1]], then the number of unique paths from the starting point to the ending point is 2.

49. Sliding Window Maximum: Given an array of integers and a window size k, find the maximum value in each window of size k as it slides from left to right. Example: If the input array is [1,3,-1,-3,5,3,6,7] and k = 3, then the maximum values in each window of size k are [3,3,5,5,6,7].

50. Remove K Digits: Given a non-negative integer num represented as a string, remove k digits from the number such that the resulting number is the smallest possible. Example: If num = "1432219" and k = 3, then the smallest possible number after removing k digits is "1219".

51. Pacific Atlantic Water Flow: Given an m x n matrix representing the heights of cells in a map, find all cells that can flow to both the Pacific and Atlantic oceans. Example: If the input matrix is [[1,2,2,3,5], [3,2,3,4,4], [2,4,5,3,1], [6,7,1,4,5], [5,1,1,2,4]], then the cells that can flow to both oceans are [(0,4), (1,3), (1,4), (2,2), (3,0), (3,1), (4,0)].

52. Longest Consecutive Sequence: Given an unsorted array of integers, find the length of the longest consecutive elements sequence. Example: If the input array is [100, 4, 200, 1, 3, 2], then the length of the longest consecutive elements sequence is 4 (by selecting the sequence [1, 2, 3, 4]).

53. Word Ladder II: Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that: Only one letter can be changed at a time, Each transformed word must exist in the word list, and the sequence must contain only unique words. Example: If beginWord = "hit", endWord = "cog", and wordList = ["hot","dot","dog","lot","log","cog"], then one possible shortest transformation sequence is ["hit","hot","dot","dog","cog"].

54. Reconstruct Itinerary: Given a list of airline tickets represented by pairs of departure and arrival airports, reconstruct the itinerary in order. All of the tickets belong to a man who departs from JFK. Thus, the itinerary must begin with JFK. Example: If the input list is [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]], then one possible itinerary is ["JFK", "MUC", "LHR", "SFO", "SJC"].

55. Minimum Window Substring: Given two strings s and t, find the minimum window in s which will contain all the characters in t. Example: If s = "ADOBECODEBANC" and t = "ABC", then the minimum window substring in s which contains all characters in t is "BANC".

56. Palindrome Partitioning II: Given a string s, partition s such that every substring of the partition is a palindrome. Return the minimum cuts needed for a palindrome partitioning of s. Example: If s = "aab", then the minimum cuts needed for a palindrome partitioning of s is 1 (by selecting the partition ["aa","b"]).

57. Jump Game II: Given an array of non-negative integers, you are initially positioned at the first index of the array. Each element in the array represents your maximum jump length at that position. Your goal is to reach the last index in the minimum number of jumps. Example: If the input array is [2,3,1,1,4], then the minimum number of jumps needed to reach the last index is 2 (by jumping from index 0 to index 1, then jumping from index 1 to index 4).

58. Kth Smallest Element in a Sorted Matrix: Given an n x n matrix where each row and column is sorted in ascending order, find the kth smallest element in the matrix. Example: If the input matrix is [[1,5,9],[10,11,13],[12,13,15]] and k = 8, then the kth smallest element in the matrix is 13.

59. Spiral Matrix II: Given a positive integer n, generate an n x n matrix filled with elements from 1 to n^2 in spiral order. Example: If n = 3, then one possible spiral matrix is [[1,2,3],[8,9,4],[7,6,5]].

60. Search a 2D Matrix: Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties: Integers in each row are sorted from left to right, and the first integer of each row is greater than the last integer of the previous row. Example: If the input matrix is [[1,3,5,7],[10,11,16,20],[23,30,34,50]] and target = 3, then the algorithm should return true.

61. Word Break II: Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word.Return all such possible sentences. Example: If s = "catsanddog" and wordDict = ["cat", "cats", "and", "sand", "dog"], then one possible sentence is ["cats and dog", "cat sand dog"].

62. Number of Islands: Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. Example: If the input grid is [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]], then the number of islands is 1.

63. House Robber III: Given the root of a binary tree, return the maximum amount of money you can rob without alerting the police. Each house in the binary tree represents a node in the tree. The money represents the amount of money in each house. If you rob a house, you cannot rob its adjacent houses (i.e., its parent and its children). Example: If the input binary tree is [3,2,3,null,3,null,1], then the maximum amount of money you can rob is 7 (by robbing the houses at nodes 3, 3, and 1).

64. Decode Ways II: A message containing letters from A-Z can be encoded into numbers using the following mapping: 'A' -> "1", 'B' -> "2", ..., 'Z' -> "26". Given a string s containing digits and '*' characters, return the number of ways to decode it. The '*' character can represent any digit from 1 to 9. Example: If s = "1*2*3", then the number of ways to decode it is 18.

65. Longest Increasing Path in a Matrix: Given an m x n integers matrix, return the length of the longest increasing path in the matrix. A path is increasing if each cell in the path is greater than the previous cell. Example: If the input matrix is [[9,9,4],[6,6,8],[2,1,1]], then the length of the longest increasing path in the matrix is 4 (by selecting the path [6, 9, 9, 8]).

66. Serialize and Deserialize Binary Tree: Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later. Deserialization is the process of converting the sequence of bits back into the original data structure or object. Design an algorithm to serialize and deserialize a binary tree. Example: If the input binary tree is [1,2,3,null,null,4,5], then one possible serialization is "1,2,3,null,null,4,5".

67. Maximal Rectangle: Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area. Example: If the input matrix is [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]], then the largest rectangle containing only 1's has an area of 6.

68. Design Twitter: Design a simplified version of Twitter where users can post tweets, follow other users, and see the tweets of the users they follow. Example: If user 1 follows users 2 and 3, and user 2 posts a tweet, then user 1 can see the tweet in their timeline.

69. Valid Sudoku: Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules: Each row must contain the digits 1-9 without repetition, Each column must contain the digits 1-9 without repetition, and Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition. Example: If the input Sudoku board is [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]], then the Sudoku board is valid.

70. Range Sum Query - Mutable: Given an integer array nums, handle multiple range sum queries efficiently. Implement the NumArray class: NumArray(int[] nums) initializes the object with the integer array nums, void update(int index, int val) updates the value of nums[index] to be val, and int sumRange(int left, int right) returns the sum of the elements of nums between indices left and right (inclusive). Example: If the input array is [1, 3, 5], then one possible sequence of operations is: NumArray obj = new NumArray(nums); obj.sumRange(0, 2); // returns 9 obj.update(1, 2); obj.sumRange(0, 2); // returns 8

71. Count of Smaller Numbers After Self: You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i]. Example: If the input array is [5,2,6,1], then the counts array would be [2,1,1,0], since there are 2 smaller elements to the right of 5, 1 smaller element to the right of 2, 1 smaller element to the right of 6, and 0 smaller elements to the right of 1.

72. Intersection of Two Arrays II: Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order. Example: If nums1 = [1,2,2,1] and nums2 = [2,2], then the intersection array would be [2,2].

73. Perfect Squares: Given an integer n, return the least number of perfect square numbers that sum to n. A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. Example: If n = 12, then the least number of perfect square numbers that sum to n is 3 (by selecting the perfect square numbers 4, 4, and 4).

74. Course Schedule II: There are a total of n courses you have to take, labeled from 0 to n-1. Some courses may have prerequisites, for example, to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]. Given the total number of courses and a list of prerequisite pairs, return the ordering of courses you should take to finish all courses. Example: If n = 4 and prerequisites = [[1,0],[2,0],[3,1],[3,2]], then one possible ordering of courses is [0,1,2,3].

75. Largest Number: Given a list of non-negative integers nums, arrange them such that they form the largest number. Example: If the input list is [3,30,34,5,9], then the largest number that can be formed is 9534330.

76. Sort Colors: Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue. Example: If the input array is [2,0,2,1,1,0], then one possible sorted array is [0,0,1,1,2,2].

77. Binary Tree Zigzag Level Order Traversal: Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between). Example: If the input binary tree is [3,9,20,null,null,15,7], then the zigzag level order traversal is [[3],[20,9],[15,7]].

78. Best Time to Buy and Sell Stock: You are given an array prices where prices[i] is the price of a given stock on the ith day. You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock. Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0. Example: If the input array is [7,1,5,3,6,4], then the maximum profit you can achieve is 5 (by buying on day 2 and selling on day 5).

79. Remove Invalid Parentheses: Given a string s that contains parentheses and letters, remove the minimum number of invalid parentheses to make the input string valid. A string is valid if every open parenthesis has a corresponding closed parenthesis and vice versa. Example: If s = "()())()", then one possible valid string after removing the minimum number of invalid parentheses is "()()()".

80. Longest Palindromic Substring:Given a string s, return the longest palindromic substring in s. A palindrome is a word, phrase, number, or other sequence of characters that reads the same backward as forward. Example: If s = "babad", then the longest palindromic substring in s is "bab" or "aba".

81. Merge k Sorted Lists: You are given an array of k linked-lists lists, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted linked-list and return it. Example: If the input array is [[1,4,5],[1,3,4],[2,6]], then one possible merged linked-list is [1,1,2,3,4,4,5,6].

82. Search in Rotated Sorted Array: You are given an integer array nums sorted in ascending order, and an integer target. Suppose nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]). You are given the target value to search for. If found in the array return its index, otherwise return -1. Example: If nums = [4,5,6,7,0,1,2] and target = 0, then the index of target in nums is 4.

83. Unique Paths II: You are given a m x n grid filled with non-negative numbers. You are also given a start cell (startX, startY) and an end cell (endX, endY). You can only move either down or right at any point in time. Some obstacles are added to the grid. Find the number of unique paths from the start cell to the end cell. An obstacle and empty space are marked as 1 and 0 respectively in the grid. Example: If the input grid is [[0,0,0],[0,1,0],[0,0,0]] and the start and end cells are (0,0) and (2,2), then the number of unique paths from the start cell to the end cell is 2.

84. Product of Array Except Self: Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i]. The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer. You must write an algorithm that runs in O(n) time and without using the division operation. Example: If the input array is [1,2,3,4], then the output array would be [24,12,8,6].

85. Longest Common Subsequence: Given two strings text1 and text2, return the length of their longest common subsequence. A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements. Example: If text1 = "abcde" and text2 = "ace", then the longest common subsequence of text1 and text2 is "ace" and its length is 3.

86. Jump Game: Given an array of non-negative integers nums, you are initially positioned at the first index of the array. Each element in the array represents your maximum jump length at that position. Determine if you are able to reach the last index. Example: If the input array is [2,3,1,1,4], then you are able to reach the last index.

87. Counting Bits: Given an integer num, return an array answer such that answer[i] is the number of bits that the integer i has in its binary representation. Example: If num = 5, then the output array would be [0,1,1,2,1,2].

88. Course Schedule III: There are n different online courses numbered from 1 to n. Each course has some duration(course length) t and closed on dth day. A course should be taken continuously for t days and must be finished before or on the dth day. You will start at the 1st day. Given n online courses represented by pairs (t,d), your task is to find the maximal number of courses that can be taken. Example: If the input array is [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]], then the maximal number of courses that can be taken is 3.

89. Validate Binary Search Tree: Given the root of a binary tree, determine if it is a valid binary search tree (BST). A valid BST is defined as follows: The left subtree of a node contains only nodes with keys less than the node's key. The right subtree of a node contains only nodes with keys greater than the node's key. Both the left and right subtrees must also be binary search trees. Example: If the input binary tree is [2,1,3], then it is a valid binary search tree.

90. Word Ladder: Given two words beginWord and endWord, and a dictionary wordList, return the length of the shortest transformation sequence from beginWord to endWord, such that: Only one letter can be changed at a time. Each transformed word must exist in the wordList. Return 0 if there is no such transformation sequence. Example: If beginWord = "hit", endWord = "cog", and wordList = ["hot","dot","dog","lot","log","cog"], then the length of the shortest transformation sequence from beginWord to endWord is 5.

91. Maximum Subarray: Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum. Example: If the input array is [-2,1,-3,4,-1,2,1,-5,4], then the contiguous subarray with the largest sum is [4,-1,2,1] and its sum is 6.

92. Reorder Data in Log Files: You have an array of logs. Each log is a space-delimited string of words, where the first word is the identifier and the rest of the words are the content. Logs with content consisting of digits are called digit-logs and logs with content consisting of letters are called letter-logs. Reorder the logs so that all of the letter-logs come before any digit-log. The letter-logs are ordered lexicographically ignoring identifier, with the identifier used in case of ties. The digit-logs should be put in their original order. Return the final order of the logs. Example: If the input log array is ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"], then one possible final order of the logs is ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"].

93. Merge Intervals: Given an array of intervals where intervals[i] = [start_i, end_i], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input. Example: If the input array is [[1,3],[2,6],[8,10],[15,18]], then one possible output array of the non-overlapping intervals is [[1,6],[8,10],[15,18]].

94. Kth Largest Element in an Array: Given an integer array nums and an integer k, return the kth largest element in the array. Note that it is the kth largest element in the sorted order, not the kth distinct element. Example: If the input array is [3,2,1,5,6,4] and k = 2, then the 2nd largest element in the array is 5.

95. Insertion Sort List: Given the head of a singly linked list, sort the list using insertion sort, and return the sorted list's head. Example: If the input linked list is [4,2,1,3], then one possible sorted linked list is [1,2,3,4].

96. Subsets: Given an integer array nums, return all possible subsets (the power set). The solution set must not contain duplicate subsets. Return the subsets in any order. Example: If the input array is [1,2,3], then one possible output set of subsets is [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]].

97. Jump Game II: Given an array of non-negative integers nums, you are initially positioned at the first index of the array. Each element in the array represents your maximum jump length at that position. Your goal is to reach the last index in the minimum number of jumps. Example: If the input array is [2,3,1,1,4], then the minimum number of jumps required to reach the last index is 2 (by starting at index 0 and jumping to index 1, then jumping from index 1 to index 4).

98. First Missing Positive: Given an unsorted integer array nums, find the smallest missing positive integer. You must implement an algorithm that runs in O(n) time and uses constant extra space. Example: If the input array is [1,2,0], then the smallest missing positive integer is 3.

99. Valid Parentheses: Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid. An input string is valid if: Open brackets must be closed by the same type of brackets. Open brackets must be closed in the correct order.Example: If the input string is "()[]{}", then it is valid.

100. Symmetric Tree: Given the root of a binary tree, check whether it is a mirror of itself (i.e. symmetric around its center). Example: If the input binary tree is [1,2,2,3,4,4,3], then it is a symmetric tree.

101. Binary Tree Level Order Traversal: Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level). Example: If the input binary tree is [3,9,20,null,null,15,7], then one possible level order traversal is [[3],[9,20],[15,7]].

102. Maximum Depth of Binary Tree: Given the root of a binary tree, return its maximum depth. A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node. Example: If the input binary tree is [3,9,20,null,null,15,7], then its maximum depth is 3.

103. Binary Tree Zigzag Level Order Traversal: Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between). Example: If the input binary tree is [3,9,20,null,null,15,7], then one possible zigzag level order traversal is [[3],[20,9],[15,7]].

104. Symmetric Tree: Given the root of a binary tree, check whether it is a mirror of itself (i.e. symmetric around its center). Example: If the input binary tree is [1,2,2,3,4,4,3], then it is a symmetric tree.

105. Construct Binary Tree from Preorder and Inorder Traversal: Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree. Example: If the input preorder array is [3,9,20,15,7] and inorder array is [9,3,15,20,7], then one possible binary tree constructed from those arrays is [3,9,20,null,null,15,7].

106. Construct Binary Tree from Inorder and Postorder Traversal: Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree. Example: If the input inorder array is [9,3,15,20,7] and postorder array is [9,15,7,20,3], then one possible binary tree constructed from those arrays is [3,9,20,null,null,15,7].

107. Binary Tree Level Order Traversal II: Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root). Example: If the input binary tree is [3,9,20,null,null,15,7], then one possible bottom-up level order traversal is [[15,7],[9,20],[3]].

108. Convert Sorted Array to Binary Search Tree: Given an integer array nums that is sorted in ascending order, convert it to a height-balanced binary search tree. A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one. Example: If the input sorted array is [-10,-3,0,5,9], then one possible height-balanced binary search tree constructed from that array is [0,-3,9,-10,null,5].

109. Convert Sorted List to Binary Search Tree: Given the head of a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST. For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1. Example: If the input linked list is [-10,-3,0,5,9], then one possible height-balanced binary search tree constructed from that list is [0,-3,9,-10,null,5].

110. Balanced Binary Tree: Given a binary tree, determine if it is height-balanced. For this problem, a height-balanced binary tree is defined as: a binary tree in which the left and right subtrees of every node differ in height by no more than 1. Example: If the input binary tree is [3,9,20,null,null,15,7], then it is a balanced binary tree.
