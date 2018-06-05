const URLify = (str) => {
  // replace all spaces in str with %20
  return str.replace(' ', '%20');
}

console.log(URLify('blue diamond'));