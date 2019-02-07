/*
    Second attempt at solving this problem
    Use dynamic programming
    Build it from scratch
*/

const isMatch = (s, p) => {
    const n = s.length;
    const m = p.length;

    if (m === 0) {
        return n === 0;
    }

    const dp = [...Array(n + 1)].map(e => Array(m + 1));

    dp[0][0] = true;

    for (let j = 1; j <= m; j++) {
        if (p.charAt(j - 1) === '*') {
            dp[0][j] = dp[0][j - 1];
        } else {
            dp[0][j] = false;
        }
    }

    for (let i = 1; i <= n; i++) {
        dp[i][0] = false;
    }

    for (let i = 1; i <= n; i++) {
        for (let j = 1;j <= m; j++) {
            if (p.charAt(j -1) === '*') {
                dp[i][j] = dp[i - 1][j] || dp[i][j - 1];
            } else if (p.charAt(j - 1) === '.' || p.charAt(j - 1) === s.charAt(i - 1)) {
                dp[i][j] = dp[i - 1][j - 1];
            } else {
                dp[i][j] = false;
            }
        }
    }


    console.log(dp);
    console.log(dp[n][m]);
    return dp[n][m];
};

isMatch('aa', 'a');