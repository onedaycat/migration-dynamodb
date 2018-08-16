# migration-dynamodb


## Installation

```sh
go get -u github.com/onedaycat/migration-dynamodb
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

