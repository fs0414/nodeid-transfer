# nodeid-transfer

```sh
brew install fs0414/tap/nodeid-transfer
```

## Description
Here's a simple CLI tool that encodes and decodes GraphQL nodeIds using base64

## Use
### Encoder
```sh
Encode a node id using base64 encoding.
Usage:
  nodeid-transfer enc [flags]
Flags:
  -h, --help              help for enc
  -i, --id string         Your id
  -o, --operator string   Your operator
```

```sh
% nodeid-transfer enc -o Object -i 1234
2025/01/13 22:50:43 T2JqZWN0OjEyMzQ=
```

### Decoder
```sh
Decode a base64 encoded node id.
Usage:
  nodeid-transfer dec [flags]
Flags:
  -e, --encoded string   Your encoded node ID
  -h, --help             help for dec
```

```sh
% nodeid-transfer dec -e T2JqZWN0OjEyMzQ=
2025/01/13 22:51:37 Operator: Object
2025/01/13 22:51:37 ID: 1234
```

## Feature
TBD
