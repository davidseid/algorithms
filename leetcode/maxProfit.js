

const maxProfit = (prices) => {
    let max = 0;
    let low = prices[0];

    for (let i = 1; i < prices.length; i++) {
        let price = prices[i];

        if (price < low) {
            low = price;
        } else {
            let diff = price - low;
            if (diff > max) max = diff;
        }
    }
    return max;
};

/*
- given an array for which the ith element is the price of a given stock on day i
- 
*/