const projectId = "pokutuna-dev";
const { Logging } = require("@google-cloud/logging");

const logging = new Logging({ projectId });
const logger = logging.log("associated");

exports["associated-logs"] = async (req, res) => {
  // 04c8c205b8c13947fcd1d7fd34c0b644/10837708167463222485;o=1
  const [context] = req.header("X-Cloud-Trace-Context").split("/");
  const trace = `projects/${projectId}/traces/${context}`;

  const entry = logger.entry(
    {
      "logging.googleapis.com/trace": trace,
      trace
    },
    {
      message: "trying to correlate",
      key: "value",
      obj: {
        a: 1,
        b: [1, 2, 3]
      }
      // "logging.googleapis.com/trace": trace,
      // trace
    }
  );

  await logger.info(entry, {
    labels: {
      execution_id: req.header("Function-Execution-ID")
    }
  });

  return res.status(200).send("ok");
};

// だめだー
// https://scrapbox.io/pokutuna/Stackdriver_に_log_書く_&_ログのグループ化(未完)
