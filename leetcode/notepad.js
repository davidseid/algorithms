/*
let m=mult(2)(3)(6);

console.log(m);//36
*/

const mult = (multiplier) => {
  return (nextMultiplier) => {
    return (nextNextMultiplier) => {
      
    }
  }
}






/*
For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].

Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].

constraints: modify input
edge cases: input values can be positive or negative
  - if no overlaps return original

visualization:

[1, 3] = start 1, end 3
[2, 6] = start 2, end 6 -> start 2 < end3 -> combine -> [1, 6]

[1, 6] compare [8, 10]

-----------------------

[1,3],[2,6],[4,7],[8,10],[15,18]

*/



// const combineOverlaps = (arr) => {
  
//   let i = 0;

//   while (i < arr.length - 1) {
//     let tuple1 = arr[i]; // 1, 3
//     let tuple2 = arr[i + 1]; // 2, 6

//     let endOfFirst = tuple1[1]; // 3
//     let startOfSecond = tuple2[0]; // 2
    
//     if (endOfFirst >= startOfSecond) {
//       arr[i][1] = arr[i + 1][1]; // 1, 6
//       arr.splice(i + 1, 1);
//     } else {
//       i++;
//     }
//   }

//   return arr;
// }

// let testInput = [[1,3],[2,6],[4,7],[8,10],[15,18]];

// console.log(combineOverlaps(testInput));