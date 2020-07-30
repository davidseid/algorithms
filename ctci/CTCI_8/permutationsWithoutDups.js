const getPermutations = (string, perm = '') => {

  let k = 1;
  let substring = string.substring(0, k);
  let permutations = [];

  while (k <= string.length) {
    if (substring.length === 1) {
      permutations.push(substring);
    } else {
      let nextPermutations = [];
      for (let i = 0; i < permutations.length; i++) {
        let permutation = permutations[i];
        for (let j = 0; j <= permutation.length; j++) {
          let part1 = permutation.slice(0, j);
          let part2 = permutation.slice(j);
          let nextPermutation = part1 + substring[substring.length - 1] + part2;
          nextPermutations.push(nextPermutation);
        }
      }
      permutations = nextPermutations;
    }
    k += 1;
    substring = string.substring(0, k);
  }
  return permutations;

}

console.log(getPermutations('dog'));
  //  // recursive solution
  //  let perms = []

  //  if (string.length === 0) return perm;
 
  //  for (let i = 0; i < string.length; i++) {
  //    let nextPerm = perm + string[i];
  //    let remaining = string.slice(0, i) + string.slice(i + 1); 
  //    perms = perms.concat(getPermutations(remaining, nextPerm)); 
  //  }
 
  //  return perms;