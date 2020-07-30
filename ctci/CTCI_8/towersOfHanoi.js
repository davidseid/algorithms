/*
*/

const transferDisks = (numDisks) => {
  const tower1 = [];
  const tower2 = [];
  const tower3 = [];

  for (let i = numDisks; i >= 1; i--) {
    tower1.push(i);
  }

  const recurse = (disks, src, tmp, dst) => {
    if (disks === 1) {
      dst.push(src.pop());
    } else if (disks === 2) {
      recurse(1, src, dst, tmp);
      recurse(1, src, tmp, dst);
      recurse(1, tmp, src, dst);
    } else {
      recurse(disks - 1, src, dst, tmp);
      recurse(1, src, tmp, dst);
      recurse(disks - 1, tmp, src, dst);
    }
  } 

  recurse(numDisks, tower1, tower2, tower3);
  
  return tower3;
}

console.log(transferDisks(12));