const checkPermutation = (str1, str2) => {
  // return bool whether one is permutation of the other
  // rather than make all permutations looking for a match
  // see if the length and character counts are the same

  if (str1.length !== str2.length) return false;

  const hash1 = {};
  const hash2 = {};
  for (let i = 0; i < str1.length; i++) {
    hash1[str1[i]] ? hash1[str1[i]]++ : hash1[str1[i]] = 1;
    hash2[str2[i]] ? hash2[str2[i]]++ : hash2[str2[i]] = 1;
  }

  for (let char in hash1) {
    if (hash1[char] !== hash2[char]) return false;
  }

  return true;
}

console.log(checkPermutation('bad', 'dka'));