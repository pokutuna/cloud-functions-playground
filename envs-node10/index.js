const projectId = "pokutuna-dev";
const { Logging } = require("@google-cloud/logging");

const logging = new Logging({ projectId });
const logger = logging.log("envs-node-8");

exports["envs-node10"] = async (req, res) => {
  const data = {
    envs: process.env,
    headers: req.headers
  };

  logger.write(logger.entry(data));
  return res.status(200).json(data);
};
