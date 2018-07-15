



let testInput = [2, 3, 1, 0, -4, 7, -3, 2]; 

const merge = (firstHalf, secondHalf) => {
  let result = [];
  let i1 = 0;
  let i2 = 0;

  while (i1 < firstHalf.length && i2 < secondHalf.length) {
    if (firstHalf[i1] <= secondHalf[i2]) {
      result.push(firstHalf[i1]);
      i1++;
    } else {
      result.push(secondHalf[i2]);
      i2++;
    }
  }

  while (i1 < firstHalf.length) {
    result.push(firstHalf[i1]);
    i1++;
  }

  while (i2 < secondHalf.length) {
    result.push(secondHalf[i2]);
    i2++;
  }
  return result;
}

const mergeSort = (arr) => {

  if (arr.length === 1) {
    return arr;
  }
  
  let mid = Math.floor(arr.length / 2);

  let firstHalf = arr.slice(0, mid);
  let secondHalf = arr.slice(mid);

  return merge(mergeSort(firstHalf), mergeSort(secondHalf));

}

console.log(mergeSort(testInput));
