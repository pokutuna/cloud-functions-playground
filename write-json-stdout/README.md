## Results

### 2020-03-03

- nodejs8
  - textPayload
- nodejs10 (Beta)
  - jsonPayload(well-known-types.value)
- go111
  - jsonPayload(well-known-types.value)
- go113 (Beta)
  - jsonPayload(well-known-types.value)
- python37
  - textPayload

### 2020-08-20

[JSON strings logged to the console are being auto-parsed into the jsonPayload instead of the textPayload in Stackdriver Logging \[147978256\] - Visible to Public - Issue Tracker](https://issuetracker.google.com/issues/147978256)

The issue marked as fixed. I checked them.

- nodejs8 (deprecated)
  - textPayload
- nodejs10
  - jsonPayload 👍
- nodejs12 (Beta)
  - jsonPayload 👍
- go111
  - jsonPayload 👍
- go113
  - jsonPayload 👍
- python37
  - textPayload
- python38 (Beta)
  - jsonPayload 👍
- java11
  - jsonPayload 👍


## Payload

Written & retrievable formats in Cloud Logging when writing this JSON string to stdout.

```json
{
  "runtime": $GCF_RUNTIME,
  "key": "value",
  "array": [1, 2, 3]
}
```

### textPayload

serialized as a string and written as `textPayload`.

```json
{
  "textPayload": "{\"runtime\":\"nodejs8\",\"key\":\"value\",\"array\":[1,2,3]}",
}
```

### jsonPayload

written as `jsonPayload` and keeps its structure.

```json
{
  "jsonPayload": {
    "runtime": $GCF_RUNTIME,
    "key": "value",
    "array": [1, 2, 3]
  },
}
```

### jsonPayload(well-known-types.value)

written as `jsonPayload` and parsed it like [Google.Protobuf.WellKnownTypes.Value](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/value).

```json
{
  "jsonPayload": {
    "fields": {
      "array": {
        "Kind": {
          "ListValue": {
            "values": [
              {
                "Kind": {
                  "NumberValue": 1
                }
              },
              {
                "Kind": {
                  "NumberValue": 2
                }
              },
              {
                "Kind": {
                  "NumberValue": 3
                }
              }
            ]
          }
        }
      },
      "key": {
        "Kind": {
          "StringValue": "value"
        }
      },
      "runtime": {
        "Kind": {
          "StringValue": "nodejs10"
        }
      }
    }
  }
}
```
