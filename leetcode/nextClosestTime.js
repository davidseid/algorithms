
const nextClosestTime = (time) => {

    let digits = time.split('').filter(digit => digit !== ':');
    const sortedDigits = digits.slice().sort();
    const lowest = sortedDigits[0];
    
    firstAttempt = tryReplacingFirstDigit(digits, sortedDigits, lowest);
    if (firstAttempt) return firstAttempt;

    secondAttempt = tryReplacingSecondDigit(digits, sortedDigits, lowest);
    if (secondAttempt) return secondAttempt;

    thirdAttempt = tryReplacingThirdDigit(digits, sortedDigits, lowest);
    if (thirdAttempt) return thirdAttempt;

    fourthAttempt = tryReplacingFourthDigit(digits, sortedDigits, lowest);
    if (fourthAttempt) return fourthAttempt;

    digits = digits.map((digit) => {
        return lowest;
    });

    digits.splice(2, 0, ':');
    console.log(digits, 'here')
    return digits.join('');
};

const tryReplacingFirstDigit = (digits, sortedDigits, lowest) => {
    for (let i = 0; i < sortedDigits.length; i++) {
        let tryDigit = sortedDigits[i];
        let currDigit = digits[3];

        if (tryDigit > currDigit) {
            digits.splice(3, 1, tryDigit); 
            digits.splice(2, 0, ':')
            return digits.join('');
        }
    }
    return false;
}

const tryReplacingSecondDigit = (digits, sortedDigits, lowest) => {
    for (let i = 0; i < sortedDigits.length; i++) {
        let tryDigit = sortedDigits[i];
        let currDigit = digits[2];

        if (tryDigit < 6 && tryDigit > currDigit) {
            digits.splice(2, 1, tryDigit); 
            digits.splice(2, 0, ':')

            digits = digits.map((digit, i) => {
                if (i > 3) {
                    return lowest;
                } else {
                    return digit;
                }
            });
            return digits.join('');
        }
    }
    return false;
}

const tryReplacingThirdDigit = (digits, sortedDigits, lowest) => {
    for (let i = 0; i < sortedDigits.length; i++) {
        let tryDigit = sortedDigits[i];
        let currDigit = digits[1];

        if (digits[0] === '2') {
            if (tryDigit < 5 && tryDigit > currDigit) {
                digits.splice(1, 1, tryDigit); 
                digits.splice(2, 0, ':')
    
                digits = digits.map((digit, i) => {
                    if (i > 2) {
                        return lowest;
                    } else {
                        return digit;
                    }
                });
                return digits.join('');
            }
        } else {
            if (tryDigit > currDigit) {
                digits.splice(1, 1, tryDigit); 
                digits.splice(2, 0, ':')
    
                digits = digits.map((digit, i) => {
                    if (i > 2) {
                        return lowest;
                    } else {
                        return digit;
                    }
                });
                return digits.join('');
            }
        }
    }
    return false;
}

const tryReplacingFourthDigit = (digits, sortedDigits, lowest) => {
    for (let i = 0; i < sortedDigits.length; i++) {
        let tryDigit = sortedDigits[i];
        let currDigit = digits[0];

        if (tryDigit < 3 && tryDigit > currDigit) {
            digits.splice(0, 1, tryDigit); 
            digits.splice(2, 0, ':')

            digits = digits.map((digit, i) => {
                if (i === 2) return digit;
                if (i > 1) {
                    return lowest;
                } else {
                    return digit;
                }
            });
            return digits.join('');
        }
    }
    return false;
}

console.log(nextClosestTime('20:48'));

/*
- HH:MM
- form the next closest time by reusing the digits
- no limit on how many times a digit can be used
- assume the input is always valid
- 01:34 is valid, 1:34 is not. 
- 19:34 -> 19:39
- 23:59 -> 22:22
*/