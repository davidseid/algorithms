
const mergeSort = (arr) => {
  let result = [];

  if (arr.length === 1) {
    result = result.concat(arr);
    return result;
  } else {
    let middle = arr.length / 2;
    let leftArr = arr.slice(0, middle);
    let rightArr = arr.slice(middle);

    return merge(mergeSort(leftArr), mergeSort(rightArr));
  }
}

const merge = (left, right) => {
  let leftIndex = 0;
  let rightIndex = 0;

  let result = [];

  while (leftIndex < left.length && rightIndex < right.length) {
    if (left[leftIndex] < right[rightIndex]) {
      result.push(left[leftIndex]);
      leftIndex++;
    } else {
      result.push(right[rightIndex]);
      rightIndex++;
    }
  }

  return result.concat(left.slice(leftIndex)).concat(right.slice(rightIndex));
}

let testArray = [2, 4, 7, 1, 2, 4, 6, 9, 10, 8];
console.log(mergeSort(testArray));
