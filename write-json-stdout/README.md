## Payload

### textPayload
```json
{
  "textPayload": "{\"runtime\":\"nodejs8\",\"key\":\"value\",\"array\":[1,2,3]}",
}
```

## jsonPayload & Google.Protobuf.WellKnownTypes.Value

https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/value

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

## 2020-03-03

- nodejs8 (GA)
  - textPayload
- nodejs10 (Beta)
  - jsonPayload
- go111 (GA)
  - jsonPayload
- go113 (Beta)
  - jsonPayload
- python37 (GA)
  - textPayload
