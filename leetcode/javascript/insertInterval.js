

const insert = (intervals, newInterval) => {

    let result = [];
    let left = 0;
    let right = intervals.length - 1;

    while (intervals[left][1] < newInterval[0]) {
        result.push(intervals[left]);
        left++;
    }

    let rightSide = [];
    while (intervals[right][0] > newInterval[1]) {
        rightSide.unshift(intervals[right]);
        right--;
    }

    let start = Math.min(intervals[left][0], newInterval[0]);
    let end = Math.max(intervals[right][1], newInterval[1]);
    result.push([start, end]);

    result = result.concat(rightSide);
    return result;
};

let testIntervals = [[1, 2,], [3, 5], [6, 7], [8, 10], [12, 16]];
let testInterval = [4, 8];
let testIntervals2 = [[1, 3], [6, 9]];
let testInterval2 = [2, 5];
console.log(insert(testIntervals2, testInterval2));














// first attempt
    // let newStart = newInterval[0];
    // let newEnd = newInterval[1];

    // let result = [];
    // let merged = false;

    // let i = 0;
    // while (i < intervals.length) {
    //     let interval = intervals[i];

    //     if (!merged) {
    //         let currStart = interval[0];
    //         let currEnd = interval[1];
    
    //         if (newEnd < currStart) {
    //             result.push(interval);
    //             i++;
    //             merged = true;
    //         } else {
                
    //             if (newEnd === currStart) {
    //                 result.push([newStart, currEnd]);
    //                 i++;
    //                 merged = true;
    //             } else if (newEnd > currStart) {

    //                 if (newStart < currEnd) {
    //                     // find end spot
    //                     let resultInterval = [currStart, null];


    //                 } else if (newStart === currEnd) {
    //                     let 
    //                 } else if (newStart > currEnd) {

    //                 }

    //             }
    //         }





    //     } else {
    //         result.push(interval);
    //         i++;
    //     }
    // }