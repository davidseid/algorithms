
// // linear
// const getCounts = (str) => {
//   const counts = {};

//   for (let i = 0; i < str.length; i++) {
//     if (counts[str[i]]) {
//       counts[str[i]]++;
//     } else {
//       counts[str[i]] = 1;
//     }
//   }

//   return counts;
// }

// // linear
// const isAnagram = (str1, str2) => {
//   if (str1.length === str2.length) {
//     let counts1 = getCounts(str1);
//     let counts2 = getCounts(str2);

//     for (let key in counts1) {
//       if (counts1[key] !== counts2[key]) {
//         return false;
//       }
//     }

//     return true;
//   }
//   return false;
// }

// // O(n2), perhaps we can make this nlogn by sorting first 
// const groupAnagrams = (strs) => {
//   const result = [];

//   const used = [];

//   strs.forEach((str) => {
//     used.push(false);
//   });

//   for (let i = 0; i < strs.length; i++) {
    
//     if (!used[i]) {
//       let strToCompare = strs[i];
//       let group = [];
//       used[i] = true;
//       group.push(strToCompare);
      
//       for (let j = 1; j < strs.length; j++) {
//         let str2 = strs[j];
        
//         if (!used[j]) {
//           if (isAnagram(strToCompare, str2)) {
//             group.push(str2);
//             used[j] = true;
//           }
//         }
//       }

//       result.push(group);

//     }
//   }

//   return result;
// }

// let testInput = ["nozzle","punjabi","waterlogged","imprison","crux","numismatists","sultans","rambles","deprecating","aware","outfield","marlborough","guardrooms","roast","wattage","shortcuts","confidential","reprint","foxtrot","dispossession","floodgate","unfriendliest","semimonthlies","dwellers","walkways","wastrels","dippers","engrossing","undertakings","unforeseen","oscilloscopes","pioneers","geller","neglects","cultivates","mantegna","elicit","couriered","shielded","shrew","heartening","lucks","teammates","jewishness","documentaries","subliming","sultan","redo","recopy","flippancy","rothko","conductor","e","carolingian"];

// console.log(groupAnagrams(testInput));

// go through creating a mapped array of objects with sortedStrings, word
// sort them by length
// go through, creating groups of like charCodeSums in batches

const groupAnagrams = (strs) => {
  const store = {};

  for (let i = 0; i < strs.length; i++) {
    let word = strs[i];
    let sortedWord = word.split('').sort((a, b) => a.charCodeAt(0) - b.charCodeAt(0)).join('');

    if (store[sortedWord]) {
      store[sortedWord].push(word);
    } else {
      store[sortedWord] = [word];
    }
  }

  
  let result = [];
  for (let key in store) {
    result.push(store[key]);
  }
  return result;
}



let testInput = ["eat", "tea", "tan", "ate", "nat", "bat"];

console.log(groupAnagrams(testInput));

// go through creating a hash table with keys as sorted str, pushing the word to that key each time
// when done, push each value to an array