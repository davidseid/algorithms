

/*
    Given an array of non-negative integers, you are initially positioned at the first index of the array.
    Each element in the array represents your maximum jump length at that position.
    Your goal is to reach the last index in the minimum number of jumps.
    Example:

    Input: [2,3,1,1,4]
    Output: 2
    Explanation: 
        - The minimum number of jumps to reach the last index is 2.
        - Jump 1 step from index 0 to 1, then 3 steps to the last index.
    Note:
        - You can assume that you can always reach the last index.
*/ 

/*
    Brute force, try all permutations, starting with the biggest jump? Not sure if is better
        - optimizations:
            - return one you reach the end, don't continue with permutations
        - use dynamic programming to do a bottom up solution, or work backwards from the end


    Method A: Recursively try each option
    Method B: Dynamic Programming

    fewest num jumps to 0 index is 0
    fewest num jumps to 1 index is 1
    fewest num jumps to 2 index is 1
    fewest num jumps to 3 index is 2
    fewest num jumps to 4 index is 2

    indices 0 1 2 3 4
    nums    2 3 1 1 4
*/

const jump = (nums) => {
    let jumps = 0;
    let endRange = 0;
    let furthest = 0;

    for (let i = 0; i < nums.length -1; i++) {
        furthest = Math.max(furthest, nums[i] + i);

        if (endRange === i) {
            jumps++;
            endRange = furthest;
        }
    }
    return jumps;
};























// // O(n) approach
// const jump = (nums) => {

//     // explanation:
//         // for each num
//             // establish a range
//             // find furthest point a number in that range can get to
//             // when we complete the range, increment jumps

//     let jumps = 0;
//     let end = 0;
//     let furthest = 0;

//     for (let i = 0; i < nums.length - 1; i++) {
//         furthest = Math.max(furthest, nums[i] + i);
//         if (i === end) {
//             jumps++;
//             end = furthest;
//         }
//     }
//     return jumps;
// };



// Dynamic programming approach, still not fast enough, greater than O(n)
// const jump = (nums) => {
//     const dp = [0];

//     for (let i = 0; i < nums.length; i++) {
//         let maxJump = nums[i];

//         for (let j = maxJump + i; j >= i + 1; j--) {
//             if (!dp[j]) {
//                 dp[j] = dp[i] + 1;
//             }

//             if (j === nums.length - 1) return dp[j];
//         }
//     }
// };


// recursive brute force
// const jump = (nums, curr = 0, target = nums.length - 1, jumps = 0, minJumps = target) => {

//     if (curr === target) {
//         return jumps < minJumps ? jumps : minJumps;
//     }
//     let max = nums[curr];

//     for (let i = max; i >= 1; i--) {
//         let distance = i;
//         let nextCurr = curr + distance;
//         if (nextCurr <= target) {
//             minJumps = jump(nums, nextCurr, target, jumps + 1, minJumps);
//         }
//     }

//     return minJumps;
// };

const testArray = [2, 3, 1, 1, 4];
console.log(jump(testArray));