
var reverse = function(x) {
  let limit = Math.pow(2, 31) - 1;
  console.log(limit);
  if (x >= 0) {
    let answer = Number(x.toString().split('').reverse().join(''));
    if (answer > limit) {
      return 0;
    } else {
      return answer;
    }
  } else {
    x = x.toString().split('').slice(1);
    x = x.reverse().join('');
    let answer = Number(x);
    if (answer > limit) {
      return 0;
    } else {
      return answer * -1;
    }
  }
};