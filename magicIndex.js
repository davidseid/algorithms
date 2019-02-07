

// first attempt, no duplicates

const testArray1 = [-2, 1, 6, 9, 12];

const findMagicIndex = (arr, start = 0, end = arr.length - 1) => {

    
    let mid = Math.floor((end - start) / 2) + start;

    if (arr[mid] === mid) return mid;

    if (start === end) return -1;

    if (arr[mid] < mid) return findMagicIndex(arr, mid + 1, end);

    if (arr[mid] > mid) return findMagicIndex(arr, start, mid - 1);
};

console.log(findMagicIndex(testArray1));