const lengthOfLongestSubstring = (s) => {
  let answer = '';
  let substring = '';
  let prevChars = {};

  for (let i = 0; i < s.length; i++) {
    let substring = '';
    let prevChars = {};
    for (let j = i; j < s.length; j++) {
      let char = s[j];
      if (prevChars[char]) {
        if (substring.length > answer.length) {
          answer = substring;
        }
        break;
      }
      substring += char;
      prevChars[char] = true;
      if (substring.length > answer.length) {
        answer = substring;
      }
    }
  }

  return answer.length;
};