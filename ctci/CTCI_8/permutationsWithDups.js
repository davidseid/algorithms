const getPermutations = (str, count = false, permutation = '', result = []) => {
  if (!count) {
    count = getCount(str);
  }

  if (str.length === 0) {
    result.push(permutation);
    return;
  } 

  for (let letter in count) {
    if (count[letter] > 0) {
      count[letter]--;
      getPermutations(str.slice(1), count, permutation + letter, result);
      count[letter]++;
    }
  }

  return result;

}

const getCount = (str) => {
  let counts = {};
  for (let i = 0; i < str.length; i++) {
    counts[str[i]] = counts[str[i]] ? counts[str[i]] + 1 : 1;
  }
  return counts;
}

console.log(getPermutations('aabbc'));