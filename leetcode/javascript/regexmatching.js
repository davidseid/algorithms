var isMatch = function(s, p) {
  // dp version
  let dp = [];
  dp[s.length][p.length] = true;

  for (let i = s.length; i >= 0; i--) {
    for (let j = p.length - 1; j >= 0; j--) {
      let firstMatch = (i < s.length && (p[j] === s[i] || p[j] === '.'));
      if (j + 1 < p.length && p[j+1] === '*') {
        dp[i][j] = dp[i][j + 2] || firstMatch && dp[i+1][j];
      } else {
        dp[i][j] = firstMatch && dp[i+1][j+1];
      }
    }
  }
  return dp[0][0];
};