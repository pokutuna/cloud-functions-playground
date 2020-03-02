const express = require("express");

const { logHttpRequest } = require("@pokutuna/requestlog-cloudfunctions");

const winston = require("winston");
const { Logging } = require("@google-cloud/logging");
const LoggingWinston = require("@google-cloud/logging-winston");

const projectId = "pokutuna-dev";

async function init() {
  const app = express();

  // httpRequest ログを記録する
  // https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#HttpRequest
  app.use(logHttpRequest({ projectId }));

  // winston logging middleware
  const logger = winston.createLogger();
  const logMiddleware = await LoggingWinston.express.makeMiddleware(logger);
  app.use(logMiddleware);

  app.get("/", (req, res) => {
    req.log.info("hello");
    return res.status(200).send("hello");
  });

  let counter = 0;
  app.get("/counter", (req, res) => {
    counter += 1;
    req.log.info("counter", { count: counter });
    return res.status(200).send(`${counter}`);
  });

  // 自前で trace ログを紐付ける
  const logging = new Logging({ projectId });
  const log = logging.log("raw");
  app.get("/trace", (req, res) => {
    const [traceId] = req.header("x-cloud-trace-context").split("/");
    const entry = log.entry(
      { trace: `projects/${projectId}/traces/${traceId}` },
      "trace manually"
    );
    log.write(entry);
    return res.status(200).send("ok");
  });

  app.post("/post", (req, res) => {
    return res.status(200).send(req.body);
  });

  return app;
}

const handler = init();
exports.app = (req, res) => handler.then(h => h(req, res));
