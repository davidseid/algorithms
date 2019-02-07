/*
    - because there are many calls to sum range
    - it is much faster to iterate through once to get all the options
    - rather than iterate each time sumRange is called
    - ex [-2, 0, 3, -5, 2, -1]


    start index 
        0    1    2    3    4    5
    0   -2   -2   1    -4   -2   -3
    1        0    3    -2   0    -1
    2             3    -2   0    -1
    3                  -5   -3   -4
    4                       2     1 
    5                            -1

    // need to refactor to make this linear! 
*/

const NumArray = function(nums) {
    this.dp = [...Array(nums.length)].map(e => new Array(nums.length));
    this.nums = nums;

    for (let k = 0; k < this.nums.length; k++) {
        
        for (let l = 0; l <= k; l++) {
            
            if (!this.dp[l][k - 1]) {
                this.dp[l][k] = this.nums[k];
            } else {
                this.dp[l][k] = this.dp[l][k - 1] + this.nums[k];
            }
        }

    }
};

NumArray.prototype.sumRange = function(i, j) {
    console.log(this.dp[i][j]);
    return this.dp[i][j];
};

const myArray = new NumArray([-2, 0, 3, -5, 2, -1]);

/*
    0    1    2    3    4    5
0   -2   -2   1   -4   -2   -3
1   -    0    3   -2    0   -1
2   -    -    3   -2    0   -1
3   -    -    -   -5   -3   -4
4   -    -    -    -    2    1
5   -    -    -    -    -   -1

*/ 

myArray.sumRange(0, 0);
