/*
  
*/




class Interval {
  constructor(start, end) {
    this.start = start;
    this.end = end;
  }
}

const merge = (intervals) => {
  
  let i = 0;

  while (i < intervals.length - 1) {
    let currInterval = intervals[i];
    let nextInterval = intervals[i + 1];

    if (currInterval.end >= nextInterval.start) {
      currInterval.end = nextInterval.end;
      intervals.splice(i + 1, 1);
    } else {
      i++;
    }
  }

  return intervals;
}