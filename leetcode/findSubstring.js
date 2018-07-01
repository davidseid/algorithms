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
  const wordLength = words[0].length;
  const indices = [];
  let startIndex = 0;
  let pointer = 0;

  while (pointer < s.length) {
    
  }
}

const exampleS = "barfoothefoobarman";

const exampleWords = ['foo', 'bar'];