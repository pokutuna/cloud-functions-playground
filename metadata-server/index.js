// https://cloud.google.com/functions/docs/securing/authenticating
// https://cloud.google.com/functions/docs/securing/function-identity
const fetch = require("node-fetch");

const idTokenURL =
  "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity?audience=test&format=full";
const accessTokenURL =
  "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/token";

const headers = {
  "Metadata-Flavor": "Google"
};

exports.app = (req, res) => {
  fetch(idTokenURL, { headers })
    .then(res => res.text())
    .then(data => console.log("id_token", data));
  fetch(accessTokenURL, { headers })
    .then(res => res.json())
    .then(data => console.log("access_token", data));
  res.status(200);
  res.send("ok");
  return;
};
