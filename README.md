# migration-dynamodb


## Installation

```sh
go get -u github.com/onedaycat/migration-dynamodb
```

## Example Yaml

```yaml
- table: user
  attr:
    - name: id
      type: S
    - name: email
      type: S
    - name: createdAt
      type: 'N'
  key:
    - name: id
      type: HASH
    - name: createdAt
      type: RANGE
  provisioned:
    read: 5
    write: 5
  globalIndex:
    - name: emailIndex
      key:
        - name: email
          type: HASH
      projection:
        type: ALL #KEYS_ONLY,INCLUDE,ALL
      provisioned:
        read: 5
        write: 5

- table: profile
  attr:
    - name: id
      type: S
  key:
    - name: id
      type: HASH
  provisioned:
    read: 5
    write: 5
```

## Usage

```go

import "github.com/onedaycat/migration-dynamodb"

m := migration.New(&migration.Config{
  AwsAccessKeyID:     "key",
  AwsSecretAccessKey: "secret",
  AwsRegion:          "ap-southeast-1",
  DynamoDBEndpoint:   "http://localhost:8000",
  Filename:           "./example/schema.yml",
})

if err := m.Up(); err != nil {
  log.Fatal(err)
}
```

