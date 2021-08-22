const { google } = require('googleapis');

exports['spreadsheet'] = async (_req, res) => {
  const auth = new google.auth.GoogleAuth({
    scopes: 'https://www.googleapis.com/auth/spreadsheets.readonly'
  });
  const sheets = google.sheets({version: 'v4', auth });
  const result = await sheets.spreadsheets.values.get({
    spreadsheetId: '1eccLY03DNCqEtJmgltI6eqGWX4kz20SBVQfOjh8JNf8',
    range: 'Sheet1!A1:D4',
  }).catch(console.error);
  console.log(result.data);
  return res.json(result.data.values);
};
