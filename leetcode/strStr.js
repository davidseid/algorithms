const strStr = (haystack, needle) => {

  if (needle === '' || haystack === needle) return 0;
  for (let i = 0; i <= haystack.length - needle.length; i++) {
let substring = haystack.substring(i, i + needle.length);
if (substring === needle) {
  return i;
}
}
return -1;
}

const test = 'aaaaa';
console.log(strStr(test, 'aab'));