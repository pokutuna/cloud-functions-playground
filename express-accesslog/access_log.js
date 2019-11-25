const { Logging, Log, middleware } = require("@google-cloud/logging");

exports.loggingAccessLog = options => {
  const projectId = options.projectId || process.env("GCLOUD_PROJECT");
  const requestLogName = Log.formatName_(
    projectId,
    options.logName || "request_log"
  );

  const logging = new Logging(options);
  const log = logging.log(requestLogName);

  const emitRequestLog = (httpRequest, trace) => {
    const entry = log.entry({ trace, httpRequest }, {});
    return log.write(entry);
  };

  return middleware.express.makeMiddleware(projectId, () => {}, emitRequestLog);
};
