const https = require('https');

async function getTitles() {

  const titles = [];

  const fetch = (page) => {
    return new Promise((resolve) => {
      https.get(`https://jsonmock.hackerrank.com/api/movies/search/?Title=waterworld&page=${page}`, (res) => {
        let rawData = '';
        res.on('data', (chunk) => { rawData += chunk; });
        res.on('end', () => {
          const parsedData = JSON.parse(rawData);
          resolve(parsedData);
        });
      });
    });
  };
  
  let page = 1;

  while (page < 3) {
    await fetch(page)
  }

  console.log(titles)

}

getTitles();




