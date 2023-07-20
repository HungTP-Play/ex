## Back tracking

Sure, here are 100 different backtracking problems:

1. Letter Combinations of a Phone Number: Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order. Example: If the input string is "23", then one possible output list of letter combinations is ["ad","ae","af","bd","be","bf","cd","ce","cf"].

2. N-Queens: Given an integer n, return all distinct solutions to the n-queens puzzle. Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively. Example: If n is 4, then one possible output list of solutions is [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]].

3. Combination Sum II: Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target. Each number in candidates may only be used once in the combination. Example: If the input array is [10,1,2,7,6,1,5] and the target is 8, then one possible output list of combinations is [[1,1,6],[1,2,5],[1,7],[2,6]].

4. Palindrome Partitioning: Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s. Example: If the input string is "aab", then one possible output list of partitions is [["a","a","b"],["aa","b"]].

5. Restore IP Addresses: Given a string s containing only digits, return all possible valid IP addresses that can be obtained from s. You can return them in any order. Example: If the input string is "25525511135", then one possible output list of IP addresses is ["255.255.11.135","255.255.111.35"].

6. Subsets II: Given an integer array nums that may contain duplicates, return all possible subsets (the power set). The solution set must not contain duplicate subsets. Return the subsets in any order. Example: If the input array is [1,2,2], then one possible output set of subsets is [[],[1],[1,2],[1,2,2],[2],[2,2]].

7. Permutations II: Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order. Example: If the input array is [1,1,2], then one possible output list of permutations is [[1,1,2],[1,2,1],[2,1,1]].

8. Combination Sum III: Find all valid combinations of k numbers that sum up to n such that the following conditions are true: Only numbers 1 through 9 are used. Each number is used at most once. Return a list of all possible valid combinations. Example: If k is 3 and n is 7, then one possible output list of combinations is [[1,2,4]].

9. Word Search II: Given an m x n grid of characters board and a list of strings words, return all words on the board. Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word. Example: If the input grid is [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]] and the input word list is ["oath","pea","eat","rain"], then one possible output list of words is ["eat","oath"].

10. Sudoku Solver: Write a program to solve a Sudoku puzzle by filling the empty cells. Empty cells are indicated by the character '.'. You may assume that there will be only one unique solution. Example: If the input puzzle is [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]], then one possible output solved puzzle is [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1. Sudoku Solver (continued): "1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","8","1","5","3","7","2","4","."],["2","4","5",".","1","9","6","8","3"],["3",".","7","2","8","6","1",".","."],["1",".",".","4",".",".",".",".","."],[".",".",".",".",".",".",".",".","."]].

11. Generate Parentheses: Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses. Example: If n is 3, then one possible output list of parentheses combinations is ["((()))","(()())","(())()","()(())","()()()"].

12. Combinations: Given two integers n and k, return all possible combinations of k numbers out of 1 ... n. Example: If n is 4 and k is 2, then one possible output list of combinations is [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]].

13. Permutation Sequence: The set [1,2,3,...,n] contains a total of n! unique permutations. Given n and k, return the kth permutation sequence. Example: If n is 3 and k is 3, then one possible output permutation sequence is "213".

14. Word Search: Given an m x n grid of characters board and a string word, return true if word exists in the grid. The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once. Example: If the input grid is [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]] and the input word is "ABCCED", then the output should be true.

15. Gray Code: The gray code is a binary numeral system where two successive values differ in only one bit. Given an integer n representing the total number of bits in the code, return any sequence of gray code. A gray code sequence must begin with 0. Example: If n is 2, then one possible output sequence of gray code is [0,1,3,2].

16. Letter Tile Possibilities: You have n tiles, where each tile has one letter tiles[i] printed on it. Return the number of possible non-empty sequences of letters you can make using the letters printed on those tiles. Example: If the input array is ["A","A","B"], then the output should be 8.

17. Beautiful Arrangement: Suppose you have n integers from 1 to n. We define a beautiful arrangement as an array that is constructed by these n numbers successfully if one of the following is true for the ith position (1 <= i <= n) in this array: The number at the ith position is divisible by i. i is divisible by the number at the ith position. Given an integer n, return the number of the beautiful arrangements that you can construct. Example: If n is 2, then the output should be 2.

18. Combination Sum IV: Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target. The answer is guaranteed to fit in a 32-bit integer. Example: If the input array is [1,2,3] and the target is 4, then the output should be 7.

19. Partition Equal Subset Sum: Given a non-empty array nums containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal. Example: If the input array is [1,5,11,5], then the output should be true.

20. Permutations III: Given an integer array nums of size n, return all the permutations of the elements in nums, which include n! distinct permutations. You can return the answer in any order. Example: If the input array is [1,2,3], then one possible output list of permutations is [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]].

21. Increasing Subsequences: Given an integer array nums, return all the different possible increasing subsequences of the given array with at least two elements. You may return the answer in any order. Example: If the input array is [4,6,7,7], then one possible output list of increasing subsequences is [[4,6],[4,7],[4,6,7],[4,6,7,7],[6,7],[6,7,7],[7,7],[4,7,7]].
