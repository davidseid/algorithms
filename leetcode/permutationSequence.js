'use strict';

const getPermutation = (n, k) => {
    const options = [];

    for (let i = 1; i <= n; i++) {
        options.push(i);
    }
    let result = '';

    let offset = k;

    for (let i = n; i > 1; i--) {
        let bucketSize = getFactorial(i) / i;
        let bucket = Math.ceil(offset / bucketSize) - 1;
        console.log(`Digit is ${i}`);
        console.log(`Bucket Size is ${bucketSize}`);
        console.log(`Offset is ${offset}`);
        console.log(`Bucket is ${bucket}`);

        let value = options.splice(bucket, 1);
        result += value;

        if (i === 3) {
            if (offset === 2) bucket = 2;
            if (offset === 1) bucket = 1;
        } else {
            if (offset > bucketSize) {
                offset = offset % bucketSize;
            }
        }
    }
    
    result += options[0];
    console.log(result);
};

const getFactorial = (n) => {
    if (n === 1) return 1;
    return n * getFactorial(n - 1);
};

getPermutation(3, 3);