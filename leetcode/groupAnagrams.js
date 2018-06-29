
// linear
const getCounts = (str) => {
  const counts = {};

  for (let i = 0; i < str.length; i++) {
    if (counts[str[i]]) {
      counts[str[i]]++;
    } else {
      counts[str[i]] = 1;
    }
  }

  return counts;
}

// linear
const isAnagram = (str1, str2) => {
  if (str1.length === str2.length) {
    let counts1 = getCounts(str1);
    let counts2 = getCounts(str2);

    for (let key in counts1) {
      if (counts1[key] !== counts2[key]) {
        return false;
      }
    }

    return true;
  }
  return false;
}

// O(n2), perhaps we can make this nlogn by sorting first 
const groupAnagrams = (strs) => {
  const result = [];

  const used = [];

  strs.forEach((str) => {
    used.push(false);
  });

  for (let i = 0; i < strs.length; i++) {
    
    if (!used[i]) {
      let strToCompare = strs[i];
      let group = [];
      used[i] = true;
      group.push(strToCompare);
      
      for (let j = 1; j < strs.length; j++) {
        let str2 = strs[j];
        
        if (!used[j]) {
          if (isAnagram(strToCompare, str2)) {
            group.push(str2);
            used[j] = true;
          }
        }
      }

      result.push(group);

    }
  }

  return result;
}

let testInput = ["eat", "tea", "tan", "ate", "nat", "bat"];

console.log(groupAnagrams(testInput));