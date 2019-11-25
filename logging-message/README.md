logging-message
====

`Log.entry` の第2引数の形式や書かれるログについて調べるための function

`$ curl -X POST -H 'Content-Type: application/json' 'https://us-central1-pokutuna-dev.cloudfunctions.net/logging-message' -d '{"message": <2ND_ARG> }'`

- `"hello"`
  - `textPayload: "hello"`
- `{}`
  - `jsonPayload: {}`
- `null`, `undefined`, `""`, `false`
  - `jsonPayload: { trace: ... }`
  - 第1引数の metadata が message として扱われる、そりゃそうか
- `true`
  - `jsonPayload: { trace: ... }`
  - 書かれるログは同上だが、`entry.toJSON` にはいずれの Payload も含まれない
