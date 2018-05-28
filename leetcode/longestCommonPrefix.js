const longestCommonPrefix = (words) => {
  let result = '';

  if (words[0].length === 0) return result;

  let shortestWord = words[0];


  for (let i = 0; i < shortestWord.length; i++) {
    let prefix = shortestWord[i];
    for (let j = 1; j < words.length; j++) {
      if (words[j][i] !== prefix) {
        return result;
      }
    }
    result += prefix;
  }
  return result;
}

let test = ["flower","flow","flight"];
console.log(longestCommonPrefix(test));