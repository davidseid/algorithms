// Input: Number indicating number of pairs of parens
// Output: An array of all valid combinations
// Constraints: None, but eventually we should optimize
// Edge cases: n = 0 return empty array

/* Explore:
  n = 1
  result = ()
  n = 2
  result = ()(), (())
  n = 3
  result = ()()(), (())(), ()(()), (()()), ((())), 
  n = 4
  result = ()()()(), (())()(), ()(())(), ()()(()), (()()()), (()())(), ((()))(), (())(()),  ()((())), ((())()), (()(())), (((())))
*/

const findValidParensCombos = (n, stack = [], combo = '') => {
  let result = [];

  if (n === 0) {
    while (stack.length > 0) {
      combo += stack.pop();
    }
    return combo;
  }

  if (stack.length > 0) {
    let closedStack = stack.slice();
    let closedCombo = combo + closedStack.pop();
    result = result.concat(findValidParensCombos(n, closedStack, closedCombo));
  } 

  let openStack = stack.slice();
  let openCombo = combo + '(';
  openStack.push(')');
  result = result.concat(findValidParensCombos(n - 1, openStack, openCombo));
  
  return result;
}

console.log(findValidParensCombos(4));