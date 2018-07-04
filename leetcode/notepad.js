const https = require('https');

const getTitles = async () => {

  const titles = [];

  const fetch = (page, results) => {
    return new Promise((resolve) => {
      https.get(`https://jsonmock.hackerrank.com/api/movies/search/?Title=spiderman&page=${page}`, (res) => {
        let rawData = '';
        res.on('data', (chunk) => { rawData += chunk; });
        res.on('end', () => {
          const parsedData = JSON.parse(rawData);
          for (let i = 0; i < parsedData.data.length; i++) {
            results.push(parsedData.data[i].Title);
          }
          resolve(titles);
        });
      });
    });
  };

  const fetchAll = async (pages, titles) => {
    return new Promise((resolve) => {
      for (let i = 0; i < pages; i++) {
        fetch(i, titles);
      }
      resolve(titles);
    });
  }

  let promises = await fetchAll(2, []);
  return promises;


}

console.log(getTitles());

/*

  What I want to do:

  have a function which returns an array of the asynchronously compiled information
  must be done with a callback
*/



