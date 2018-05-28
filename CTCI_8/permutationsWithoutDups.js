const getPermutations = (string, perm = '', perms = []) => {
  // get all permutations of a string of unique characters
  // i: string (no caps, no spaces, all a-z)
  // o: Array of all possible permutations of that string
  // c: brute first, then try to optimize
  // e: '', 'dog'

  // dog 
  // dog dgo odg ogd gdo god 
  // do
  // 12 21 (2)
  // 123 132 213 232 312 321 (6)
  // 4123 4132 4213 4232 4312 4321 1423 1432 2413 2432 3412 3421 1243 1342 2143 2342 3142 3241 1234 1324 2134 2324 3124 3214 (24)

  // if n is length of string, number of permutations = n!
  // 2 -> 2 * 1 = 2
  // 3 -> 3 * 2 = 6
  // 4 -> 4 * 3 * 2 = 24

  // recursive solution

  if (string.length === 0) return perm;

  for (let i = 0; i < string.length; i++) {
    let nextPerm = perm + string[i];
    let remaining = string.slice(0, i) + string.slice(i + 1); // og
    perms = perms.concat(getPermutations(remaining, nextPerm)); // og, [d]
  }

  return perms;
}