exports["show-pubsub-data"] = async (req, res) => {
  let data = {};
  try {
    data = JSON.parse(Buffer.from(req.body.message.data, "base64").toString());
  } catch {
    data = {};
  }

  const request = {
    envs: process.env,
    headers: req.headers,
    body: req.body,
    decodedData: data,
  };

  console.log(JSON.stringify(request));

  return res.status(200).send("ok");
};
