

const lengthOfLastWord = (s) => {

  let words = s.split(' ');

  console.log(words)

  for (let i = words.length - 1; i >= 0; i--) {
    let word = words[i];
    if (word !== '') return word.length;
  }

  return 0;
  
}

console.log(lengthOfLastWord('a '));