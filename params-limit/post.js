const fetch = require('node-fetch');
const { stringify } = require('qs');

const url = 'https://us-central1-pokutuna-dev.cloudfunctions.net/params-limit';

const size = parseInt(process.argv[2], 10);
const body = [...Array(size).keys()].reduce((obj, i) => {
  obj[i] = i;
  return obj;
}, {});

(async () => {
  fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
    },
    body: stringify(body),
  }).catch(console.err).then(res => {
    console.log(res.status);
    return res;
  }).then(res => res.json()).then(console.log);
})();
