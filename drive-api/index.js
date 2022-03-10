const { google } = require("googleapis");

exports["drive-api"] = async (req, res) => {
  const auth = new google.auth.GoogleAuth({
    projectId: "pokutuna-dev",
    scopes: [
      "https://www.googleapis.com/auth/cloud-platform",
      "https://www.googleapis.com/auth/drive.readonly",
    ],
  });

  const drive = google.drive({ version: "v3", auth });
  await drive.files
    .list({
      q: { pageSize: 3 },
    })
    .then((res) => console.log(JSON.stringify(res.data)))
    .catch(console.error);

  res.send("ok");
};
