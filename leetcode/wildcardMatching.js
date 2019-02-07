
/*
    use recursion
*/

const isMatch = (s, p) => {
    const n = s.length;
    const m = p.length; 

    if (m === 0) {
        return n === 0;
    }

    dp = [...Array(n + 1)].map(e => Array(m + 1));

    for (let i = 0; i < n + 1; i++) {
        dp[i][0] = false;
    }

    dp[0][0] = true;

    for (let j = 1; j <= m; j++) {
        if (p.charAt(j - 1) === '*') {
            dp[0][j] = dp[0][j - 1];
        } else {
            dp[0][j] = false;
        }
    }

    for (let i = 1; i <= n; i++) {
        for (let j = 1; j <= m; j++) {
            if (p.charAt(j - 1) === '*') {
                dp[i][j] = dp[i][j - 1] || dp[i - 1][j];
            } else if (p.charAt(j - 1) === '?' || s.charAt(i - 1) === p.charAt(j - 1)) {
                dp[i][j] = dp[i - 1][j - 1];
            } else {
                dp[i][j] = false;
            }
        }
    }

    console.log(dp);
    return dp[n][m];
};

console.log(isMatch("aa", "a"));
console.log(isMatch("aa", "*"));
console.log(isMatch("cb", "?a"));
console.log(isMatch("adceb", "*a*b"));
console.log(isMatch("acdcb", "a*c?b"));