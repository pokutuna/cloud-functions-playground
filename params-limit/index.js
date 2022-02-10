const express = require('express');

const app = express();

app.use(express.urlencoded({parameterLimit: 2000, extended: true}));

app.all('*', (req, res) => {
  const obj = {
    body: req.body,
    parameters: Object.keys(req.body).length,
  };

  console.log(JSON.stringify(obj));

  res.json(obj);
});

exports['app'] = app;
