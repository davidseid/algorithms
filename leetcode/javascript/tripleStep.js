


const tripleStep = (n) => {
    let result = 0;
    let memo = {};

    if (n === 0) return 1;
    if (n < 0) return 0;

    for (let i = 1; i <= 3; i++) {

        if (!memo[n - i]) {
            memo[n - i] = tripleStep(n - i);
        }
        result += memo[n - i];
    }
    return result;
};

console.log(tripleStep(3));