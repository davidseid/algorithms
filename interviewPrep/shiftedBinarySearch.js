
const testArray = [4, 5, 6, 7, 8, 9, 1, 2, 3];

const shiftedBinarySearch = (array, target, start = 0, end = array.length - 1) => {

    let mid = start + Math.floor((end - start) / 2);

    if (array[mid] === target) return mid;

    if (array[start] <= array[mid]) {

        // if the first half is sorted

        if (target >= array[start] && target <= array[mid]) {
            // if the target is in the first half, search it

            return shiftedBinarySearch(array, target, start, mid);

        } else {
            // if the target is in the second half, search it

            return shiftedBinarySearch(array, target, mid, end);
        }
    } else if (target >= array[mid] && target <= array[end]) {

        // if the target is greater than the middle and less than the end, search that side

        return shiftedBinarySearch(array, target, mid, end);
    } else {

        // the first half is unsorted && the target does not appear to be in the second half
        return shiftedBinarySearch(array, target, start, mid);
    }
}

console.log(shiftedBinarySearch(testArray, 9));