exports.app = (req, res) => {
  const obj = { key: "value", array: [1, 2, 3] }
  console.log(JSON.stringify(obj));
  res.json(obj);
};
