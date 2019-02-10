const testArray = [6, 3, 1, 5, 7, 9, 2, 4];

const quickSort = (arr, left = 0, right = arr.length - 1) => {
  let index = partition(arr, left, right);

  if (left < index - 1) {
    quickSort(arr, left, index - 1);
  }
  if (index < right) {
    quickSort(arr, index, right);
  }
};

const partition = (arr, left, right) => {
  const pivot = arr[Math.floor((left + right) / 2)];
  
  while (left <= right) {

    while (arr[left] < pivot) left++;

    while (arr[right] > pivot) right--;

    if (left <= right) {
      swap(arr, left, right);
      left++;
      right--;
      console.log(arr);
    }
  }
  return left;
};

const swap = (arr, left, right) => {
  let temp = arr[left];
  arr[left] = arr[right];
  arr[right] = temp;
};

console.log(quickSort(testArray));