var longestPalindrome = function(s) {
  let longest = '';
  if (s.length === 1) {
    return s;
  }
  for (let i = 0; i < s.length; i++) {
    let startIndex = i;
    let endIndex = i;
    while (s[startIndex] === s[endIndex + 1]) {
      endIndex++;
    }
    while (startIndex - 1 >= 0 && endIndex + 1 < s.length && s[startIndex - 1] === s[endIndex + 1]) {
      startIndex--;
      endIndex++;
    }
    let palindrome;
    if (endIndex > s.length - 1) {
      palindrome = s[startIndex];
    } else {
      palindrome = s.substring(startIndex, endIndex + 1);
    }
    if (palindrome.length > longest.length) {
      longest = palindrome;
    }
  }
  return longest;
};