/*
*/

const transferDisks = (numDisks) => {
  const tower1 = [];
  const tower2 = [];
  const tower3 = [];

  for (let i = numDisks; i >= 1; i--) {
    tower1.push(i);
  }

  const recurse = (remainingDisks, source, destination) => {
    if (remainingDisks === 1) {
      destination.push(source.pop());
    } else if (remainingDisks === 2) {
      recurse(1, tower1, tower2);
      recurse(1, tower1, tower3);
      recurse(1, tower2, tower3);
    }
  } 

  recurse(numDisks, tower1, tower3);
  
  return tower3;
}

console.log(transferDisks(2));