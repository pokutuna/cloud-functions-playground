associated-logs
===

- 親のログに `httpRequest` の構造がないとダメそう
- `execution_id` や GAE と同様に `trace` を引き回してもだめ
- JSON を書いても textPayload になるので logging client つかわないとだめそう
