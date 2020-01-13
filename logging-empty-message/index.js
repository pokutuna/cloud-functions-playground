const projectId = "pokutuna-dev";
const functionName = "logging-empty-message";

const { Logging } = require("@google-cloud/logging");

const logging = new Logging({ projectId });
const logger = logging.log(functionName);

exports[functionName] = async (req, res) => {
  const [traceId] = req.header("x-cloud-trace-context").split("/");
  const entry = logger.logging.entry(
    { trace: `projects/${projectId}/traces/${traceId}` },
  );
  logger.write(entry);
  return res.status(200).json(entry.toJSON());
};
