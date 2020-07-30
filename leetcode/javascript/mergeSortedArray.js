const shiftElements = (arr, startIndex, lastIndex) => {
  for (let i = lastIndex; i >= startIndex; i--) {
    arr[i + 1] = arr[i];
  }
}

const merge = (nums1, m, nums2, n) => {
  let p1 = 0;
  let p2 = 0;
  let lastIndex = m - 1;

  while (m > 0 && n > 0) {
    console.log(nums1);
    console.log(nums2);
    let num1 = nums1[p1];
    let num2 = nums2[p2];
    console.log(num1, num2)
    if (num1 <= num2) {
      p1++;
      m--;
    } else {
      shiftElements(nums1, p1, lastIndex); 
      lastIndex++;
      nums1[p1] = num2;
      p2++;
      m++;
      n--;
    }
  }

  if (n > 0) {
    for (let i = p2; i < nums2.length; i++) {
      nums1[p1] = nums2[i];
      p1++;
    }
  }
}


/* Tests */

let testNums1 = [1, 2, 3, 0, 0, 0];
let testNums2 = [2, 5, 6];

// result = [1, 2, 2, 3, 5, 6];

console.log(testNums1);
merge(testNums1, 3, testNums2, 3);
console.log(testNums1);