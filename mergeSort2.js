const mergeSort = (items) => {

  if (items.length === 1) {
    return items;
  } else {
    let firstHalf = items.slice(0, Math.floor(items.length / 2));
    let secondHalf = items.slice(Math.floor(items.length / 2));
    return merge(mergeSort(firstHalf), mergeSort(secondHalf));
  }
}

const merge = (arr1, arr2) => {
  let leftPointer = 0;
  let rightPointer = 0;

  let result = [];

  while (leftPointer < arr1.length && rightPointer < arr2.length) {
    if (arr1[leftPointer] <= arr2[rightPointer]) {
      result.push(arr1[leftPointer]);
      leftPointer++;
    } else {
      result.push(arr2[rightPointer]);
      rightPointer++;
    }
  }

  if (leftPointer < arr1.length) {
    result = result.concat(arr1.slice(leftPointer));
  } else {
    result = result.concat(arr2.slice(rightPointer));
  }

  return result;
}

console.log(mergeSort([1, 4, 7, 2, 3, 0]));