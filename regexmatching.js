var isMatch = function(s, p) {
  // make sIndex
  // loop through p
  // pChar
  // pNext
  // if pChar === s[sIndex] && pNext === *
    // while s[sIndex] === pChar 
      // sIndex++
  // if pChar === a-z && pChar !== s[sIndex] pNext === *
    // sIndex++
  // if pChar === s[sIndex] && pNext !== *
    // sIndex++
  // if pChar === . && pNext === *
    // sChar
    // while s[sIndex] === sChar
      // sIndex++
  // if pChar === . && pNext !== *
    // sIndex++

  let sIndex = 0;
  for (let i = 0; i < p.length; i++) {
    let pChar = p[i];
    let pNext = p[i + 1];

    if (pChar === s[sIndex] && pNext === '*') {
      while (s[sIndex] === pChar) {
        sIndex++;
      }
    } else if (pChar >= 'a' && pChar <= 'z' && pChar !== s[sIndex] && pNext === '*') {
      sIndex++;
    } else if (pChar === s[sIndex] && pNext !== '*') {
      sIndex++;
    } else if (pChar === '.' && pNext === '*') {
      let sChar = s[sIndex];
      while (s[sIndex] === sChar) {
        sIndex++;
      }
    } else if (pChar === '.' && pNext !== '*') {
      sIndex++;
    } else {
      return false;
    }
  }

  return true;

};