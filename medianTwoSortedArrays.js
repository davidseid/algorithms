var findMedianSortedArrays = function(nums1, nums2) {
  let mergedArray = []
  let secondPointer = 0;
  for (let i = 0; i < nums1.length; i++) {
    let num1 = nums1[i];
    while (nums2[secondPointer] <= num1) {
      let num2 = nums2[secondPointer];
      mergedArray.push(num2);
      secondPointer++;
    }
    mergedArray.push(num1);
  }

  if (mergedArray.length < (nums1.length + nums2.length)) {
    for (let i = secondPointer; i < nums2.length; i++) {
      mergedArray.push(nums2[i]);
    }
  }

  if (mergedArray.length % 2 === 0) {
    let medianIndex1 = mergedArray.length / 2;
    let medianIndex2 = medianIndex1 - 1;
    return ((mergedArray[medianIndex1] + mergedArray[medianIndex2]) / 2);
  } else {
    let medianIndex = ((mergedArray.length - 1) / 2);
    return mergedArray[medianIndex];
  }
};