/*
  iterate through s
  find a match in words if it exists
  if it does, remember start index
  cross off the word and keep going until all words are crossed
  or until a non match
  add index to result if completed

  // figoured out a way to do it linear with respect to s
*/


const findSubstring = (s, words) => {

  if (s.length === 0 || words.length === 0) return [];
  const wordLength = words[0].length;
  const wordStore = {};

  for (let i = 0; i < words.length; i++) {
    let word = words[i];
    wordStore[word] ? wordStore[word]++ : wordStore[word] = 1;
  }
  const result = [];
  let pointer = 0;

  while (pointer < s.length) {
    let startIndex = pointer;
    let currentWord = s.substr(pointer, wordLength);
    let wordsInRow = 0;
    let usedWords = {};

    while (wordStore[currentWord] && (wordStore[currentWord] !== usedWords[currentWord])) {
      wordsInRow++;
      usedWords[currentWord] ? usedWords[currentWord]++ : usedWords[currentWord] = 1;
      pointer += wordLength;
      currentWord = s.substr(pointer, wordLength);
    }

    if (wordsInRow === words.length) {
      result.push(startIndex);
    }

    pointer++;;
  }
  return result;
};

const exampleS = "barfoothefoobarman";

const exampleWords = ['foo', 'bar'];

console.log(findSubstring(exampleS, exampleWords))