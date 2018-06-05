const letterCombinations = (digits, combo = '') => {
  const numbersToLetters = {
    '2': 'abc',
    '3': 'def',
    '4': 'ghi',
    '5': 'jkl',
    '6': 'mno',
    '7': 'pqrs',
    '8': 'tuv',
    '9': 'wxyz'
  }

  let combinations = [];

  if (digits.length === 0) return combo;

  let digit = digits[0];
  let letters = numbersToLetters[digit];

  for (let j = 0; j < letters.length; j++) {
    let letter = letters[j];
    let newCombo = combo + letter;
    combinations = combinations.concat(letterCombinations(digits.slice(1), newCombo));
  }

  return combinations;
}

console.log(letterCombinations('23'));