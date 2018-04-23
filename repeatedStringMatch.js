var repeatedStringMatch = function(A, B) {
  let possibleCounterStarts = [];
  for (let i = 0; i < A.length; i++) {
    if (A[i] === B[0]) {
      possibleCounterStarts.push(i);
    } 
  }

  if (possibleCounterStarts.length === 0) {
    return -1;
  }

  let result;
  
  for (let k = 0; k < possibleCounterStarts.length; k++) {
    let numRepeats = 1;
    let counter = possibleCounterStarts[k];
    
    for (let i = 0; i < B.length; i++) {
      let bChar = B[i];
      let aChar = A[counter];
      if (aChar !== bChar) {
        return -1; 
      } else {
        if (A[counter + 1] === undefined) {
          counter = 0;
          numRepeats++;
        } else {
          counter++;
        }
      }

    }
    if (result === undefined) {
      result = numRepeats;
    }
    if (numRepeats < result) {
      result = numRepeats;
    }

  }
  
  return result;
};