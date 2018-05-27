const findMagicIndexDistinct = (arr) => {
  // given sorted array of distinct integers
  // find example of arr[i] === i;

  for (let i = 0; i < arr.length; i++) {
    if (arr[i] === i) {
      return i;
    }
  }
  return -1;
}

const findMagicIndexNotDistinct = (arr) => {

}