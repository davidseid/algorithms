const isValid = (s) => {
  let stack = [];

  for (let i = 0; i < s.length; i++) {
    if (s[i] === '(' ||
        s[i] === '[' ||
        s[i] === '{') {
      stack.push(s[i]);
        } else {
          let mostRecent = stack[stack.length - 1];
          if (s[i] === ')') {
            if (mostRecent !== '(') return false;
            stack.pop();
          }
          if (s[i] === ']') {
            if (mostRecent !== '[') return false;
            stack.pop();
          }
          if (s[i] === '}') {
            if (mostRecent !== '{') return false;
            stack.pop();
          }
        }
    if (stack.length > s.length - i) return false;
  }

  if (stack.length !== 0) return false;

  return true;
}