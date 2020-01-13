const projectId = "pokutuna-dev";

const tracer = require("@google-cloud/trace-agent").start({
  projectId,
  bufferSize: 1,
  keyFilename: './key.json'
});

const express = require("express");
const winston = require("winston");

const LoggingWinston = require("@google-cloud/logging-winston");
const { logHttpRequest } = require("requestlog-cloudfunctions");
const traceMiddleware = require("./middleware-trace");
const other = require('./other')

const handler = (async () => {
  const app = express();

  app.use(traceMiddleware(tracer, { name: "handler" }));
  app.use(logHttpRequest({ projectId }));

  // winston logging middleware
  const logger = winston.createLogger();
  const logMiddleware = await LoggingWinston.express.makeMiddleware(logger);
  app.use(logMiddleware);

  app.get("/", (req, res) => {
    req.log.info("hello", { ips: req.ips, ip: req.ip });
    other.hoge();
    other.fuga();
    return res.status(200).send("hello");
  });

  let counter = 0;
  app.get("/counter", (req, res) => {
    counter += 1;
    req.log.info("counter", { count: counter });
    other.hoge();
    other.fuga();
    return res.status(200).send(`${counter}`);
  });

  return app;
})();

exports.app = (req, res) => handler.then(h => h(req, res));
