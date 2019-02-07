
/*
each house has some money
can't rob adjacent houses
determine max amount of money can rob

[1, 2, 3, 1]





*/


const rob = (nums) => {
    if (nums.length === 0) return 0;
    if (nums.length === 1) return nums[0];
    if (nums.length === 2) return Math.max(nums[0], nums[1]);

    const dp = [];

    dp[0] = nums[0];
    dp[1] = Math.max(nums[0], nums[1]);

    for (let i = 2; i < nums.length; i++) {
        dp[i] = Math.max(dp[i - 2] + nums[i], dp[i -1]);
    }
    console.log(dp);
    return dp[nums.length - 1];
}




















// const rob = (nums) => {

//     n = nums.length;

//     if (n === 0) return 0;
//     if (n === 1) return nums[0];
//     if (n === 2) return Math.max(nums[0], nums[1]);

//     dp = [];

//     dp[0] = nums[0];
//     dp[1] = Math.max(nums[0], nums[1]);

//     for (let i = 2; i < n; i++) {
//         dp[i] = Math.max(nums[i] + dp[i - 2], dp[i - 1]);
//     }
//     return dp[n - 1];
    
// };

const test1 = [1, 2, 3, 1];
const test2 = [2, 7, 9, 3, 1];
rob(test1);
rob(test2);

