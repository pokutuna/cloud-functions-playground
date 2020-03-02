exports.App = (req, res) => {
  const payload = {
    runtime: process.env.GCF_RUNTIME,
    key: "value",
    array: [1, 2, 3]
  };
  console.log(JSON.stringify(payload));
  res.json(payload);
};
