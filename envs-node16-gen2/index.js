exports.gen2 = function(req, res) {
  const data = {
    envs: process.env,
    headers: req.headers,
  };

  console.log(JSON.stringify(data));
  return res.status(200).json(data);
}
