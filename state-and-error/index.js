let count = 0;

exports.app = (req, res) => {
  if (req.query.throw) {
    throw new Exception("exception");
  }

  count += 1;
  return req.res.json({ count });
};
