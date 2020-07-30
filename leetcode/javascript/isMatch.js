
let string = 'acdcb';
let pattern = 'a*c?b';


const isMatch = (s, p) => {
  if (s.length === 0 || p.length === 0) return false;

  let result = false;
  let si = 0;
  let pi = 0;

  while (si < s.length && pi < p.length) {

    let sChar = s[si];
    let pChar = p[pi];

    if (sChar === pChar) {
      si++;
      pi++;
    } else if (pChar === '?') {
      si++;
      pi++;
    } else if (pChar === '*') {
      let nextP = p[pi + 1];
      if (nextP) {
        for (let i = si; i < s.length; i++) {
          if (s[i] === nextP) {
            si = i + 1;
            pi++;
            console.log(si, pi, i)
            result = result || isMatch(s.substring(si), p.substring(pi + 1));
          }
        }
      }
    } else {
      return false;
    }
  }

  if (pi < p.length) {
    while (p[pi] === '*') {
      pi++;
    }
  } 

  if (pi >= p.length) {
    return true;
  } else {
    return false;
  }
}

console.log(isMatch(string, pattern));