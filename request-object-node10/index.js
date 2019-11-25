const projectId = "pokutuna-dev";
const functionName = "request-object-node10";

const { Logging } = require("@google-cloud/logging");

const logging = new Logging({ projectId });
const logger = logging.log(functionName);

exports[functionName] = async (req, res) => {
  const prop = req.query.prop;
  const data = {
    result: req[prop]
  };
  logger.write(logger.entry(data));
  return res.status(200).json(data);
};
