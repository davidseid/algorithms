const findMagicIndexDistinct = (arr) => {
  // given sorted array of distinct integers
  // find example of arr[i] === i;

  let midIndex = arr.length / 2;
  
  if (arr[midIndex] === midIndex) {
    return midIndex;
  } else {
    if (arr[midIndex] < midIndex) {
      return findMagicIndexDistinct(arr.slice(midIndex));
    } else {
      return findMagicIndexDistinct(arr.slice(0, midIndex));
    }
  }
}

const distinctTestArr = [-40, -12, -3, 0, 4, 6, 7, 8];
const notDistinctTestArr = [-10, -2, 2, 4, 4, 6, 7];

// const findMagicIndexNotDistinct = (arr) => {
//   let midIndex = arr.length / 2;
//   if (arr[midIndex] === midIndex) {
//     return midIndex;
//   } else {
//     return findMagicIndexNotDistinct 
//   }


// }