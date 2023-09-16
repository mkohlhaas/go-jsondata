### Go and JSON

[How To Use JSON in Go](https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go)

```shell
$ echo '{"boolValue":true,"dateValue":"2022-03-02T09:10:00Z","intValue":1234,"now":"2023-09-17T18:31:57.472089891+02:00","nullIntValue":null,"nullStringValue":null,"objectValue":{"arrayValue":[1,2,3,4]},"stringValue":"hello!"}' | jq
```

```json
{
  "boolValue": true,
  "dateValue": "2022-03-02T09:10:00Z",
  "intValue": 1234,
  "now": "2023-09-17T18:31:57.472089891+02:00",
  "nullIntValue": null,
  "nullStringValue": null,
  "objectValue": {
    "arrayValue": [
      1,
      2,
      3,
      4
    ]
  },
  "stringValue": "hello!"
}
```

```shell
$ echo '{"int_value":1234,"bool_value":true,"string_value":"hello!","date_value":"2022-03-02T09:10:00Z","object_value":{"array_value":[1,2,3,4]},"nullInt_value":4321}' | jq
```

```json
{
  "int_value": 1234,
  "bool_value": true,
  "string_value": "hello!",
  "date_value": "2022-03-02T09:10:00Z",
  "object_value": {
    "array_value": [
      1,
      2,
      3,
      4
    ]
  },
  "nullInt_value": 4321
}
```
