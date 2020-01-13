module.exports = function enableTrace(tracer, options) {
  return (req, res, next) => {
    tracer.runInRootSpan(options, span => {
      req.rootSpan = span;
      next();
      span.endSpan();
    });
  };
};
