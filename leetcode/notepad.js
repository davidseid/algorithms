const https = require('https');

const getNumPages = (keyword) => {
  return new Promise((resolve) => {
    https.get(`https://jsonmock.hackerrank.com/api/movies/search/?Title=${keyword}`, (res) => {
      let rawData = '';
      res.on('data', (chunk) => {
        rawData += chunk;
      });
      res.on('end', () => {
        let parsedData = JSON.parse(rawData);
        resolve(parsedData.total_pages);
      })
    });
  });
};
  
const assembleTitles = (data) => {
  let titles = [];

  data.data.forEach((movie) => {
    titles.push(movie.Title);
  });

  return titles;
}

const fetchTitlesByPage = (keyword, page) => {
  return new Promise((resolve) => {
    https.get(`https://jsonmock.hackerrank.com/api/movies/search/?Title=${keyword}&page=${page}`, (res) => {
      let rawData = '';
      res.on('data', (chunk) => {
        rawData += chunk;
      });
      res.on('end', () => {
        let parsedData = JSON.parse(rawData);
        let titles = assembleTitles(parsedData);
        resolve(titles);
      })
    });
  })
}

const fetchAllTitles = (keyword, numPages) => {
  let titles = [];

  for (let i = 1; i <= numPages; i++) {
    titles = titles.concat(fetchTitlesByPage(keyword, i));
  }
  return titles;
}

const flatten = (nestedArr) => {
  let result = [];
  for (let i = 0; i < nestedArr.length; i++) {
    result = result.concat(nestedArr[i]);
  }
  return result;
}

const getTitles = (keyword) => {
  
  return new Promise((resolve) => {
    getNumPages(keyword)
      .then(res => {
        return Promise.all(fetchAllTitles(keyword, res));
      })
      .then(res => {
        let flattened = flatten(res);
        console.log(flattened);
        resolve(flattened);
      });
  });
};

console.log(getTitles('waterworld'));

























// const assembleTitlesFromData = (data) => {
//   let titles = [];

//   data.data.forEach((movie) => {
//     titles.push(movie.Title);
//   });
//   return titles;
// }

// const fetchTitlesByPage = (keyword, page) => {
//   return new Promise((resolve) => {
//     https.get(`https://jsonmock.hackerrank.com/api/movies/search/?Title=${keyword}&page=${page}`, (res) => {
//       let rawData = '';
//       res.on('data', (chunk) => {
//         rawData += chunk;
//       });

//       res.on('end', () => {
//         let parsedData = JSON.parse(rawData);
//         resolve(assembleTitlesFromData(parsedData));
//       });
//     });
//   });
// };

// const fetchAllTitles = (keyword, pages) => {
//   let titles = [];
//   for (let i = 1; i <= pages; i++) {
//     titles = titles.concat(fetchTitlesByPage(keyword, i));
//   }
//   return Promise.all(titles)
//     .then((res) => {
//       let result = [];

//       for (let i = 0; i < res.length; i++) {
//         result = result.concat(res[i]);
//       }

//       result = result.sort();
      
//       console.log(result);
//       return result;
//     });
// };

// const getTitles = (keyword) => {

//   let titles = fetchAllTitles(keyword, 2);
  
//   return titles;
  
// };

// console.log(getTitles('waterworld'));


