const testArray = [3, 1, 7, 4, 0];

const mergeSort = (arr) => {
  if (arr.length <= 1) return arr;

  const mid = Math.floor(arr.length / 2);
  
  let left = arr.slice(0, mid);
  let right = arr.slice(mid);

  console.log(left, right);

  return merge(mergeSort(left), mergeSort(right));
};

// helper
const merge = (arr1, arr2) => {

  let sorted = [];
  let left = 0;
  let right = 0;

  while (left < arr1.length && right < arr2.length) {
    if (arr1[left] <= arr2[right]) {
      sorted.push(arr1[left]);
      left++;
    } else if (arr2[right] < arr1[left]) {
      sorted.push(arr2[right]);
      right++;
    }
  }

  console.log(sorted);

  while (left < arr1.length) {
    sorted.push(arr1[left]);
    left++;
  }

  while (right < arr2.length) {
    sorted.push(arr2[right]);
    right++;
  }

  return sorted;
};

console.log(mergeSort(testArray));