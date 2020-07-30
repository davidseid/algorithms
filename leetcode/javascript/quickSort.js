

/*
  Quicksort is a sorting algorithm that uses a pivot (i believe selected at random) 
  and makes sure to put everything in the array smaller than the pivot in front, bigger in bacck
  it repeats until sorted

  
*/

const swapNodes = (arr, left, right) => {
  let temp = arr[left];
  arr[left] = arr[right];
  arr[right] = temp;
}

const partition = (arr, pivot, left, right) => {
  let partitionIndex = left;

  for (let i = left; i < right; i++) {
    if (arr[i] < arr[pivot]) {
      swapNodes(arr, i, partitionIndex);
      partitionIndex++;
    }
  }

  swapNodes(arr, partitionIndex, right);
  return partitionIndex;
}

const quickSort = (arr, left, right) => {
  
  if (left < right) {
    let pivot = right;
    let partitionIndex = partition(arr, pivot, left, right);

    quickSort(arr, left, partitionIndex - 1);
    quickSort(arr, partitionIndex + 1, right);
  }

  return arr;
}

let testInput = [2, 1, 9, 6, 5, 11, 3];

console.log(quickSort(testInput, 0, testInput.length - 1));