const https = require('https');

const getTitles = () => {

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
    for (let i = 0; i < pages; i++) {
      await fetch(i, titles);
    }
    console.log(titles)
    return titles;
  }

  return fetchAll(2, titles);



  
  // fetch(1, titles)
  // .then((res) => fetch(2, res))
  // .then((res) => console.log(titles));
}

getTitles();




