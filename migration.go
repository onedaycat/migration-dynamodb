package migration

import (
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gopkg.in/yaml.v2"
)

type Migration struct {
	Config  *Config
	schemas []*Schema
	dynamo  *dynamodb.DynamoDB
}

type Config struct {
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsRegion          string
	DynamoDBEndpoint   string
	Filename           string
}

func New(config *Config) *Migration {
	m := &Migration{Config: config}

	schemas, err := m.createSchemas(m.Config.Filename)
	if err != nil {
		log.Fatal(err)
	}
	m.schemas = schemas

	dy, err := m.createDynamoDB()
	if err != nil {
		log.Fatal(err)
	}
	m.dynamo = dy

	return m
}

func (m *Migration) createDynamoDB() (*dynamodb.DynamoDB, error) {
	configDynamodb := aws.NewConfig()
	cred := credentials.NewStaticCredentials(m.Config.AwsAccessKeyID, m.Config.AwsSecretAccessKey, "")
	configDynamodb.WithCredentials(cred)
	configDynamodb.WithRegion(m.Config.AwsRegion)
	configDynamodb.WithEndpoint(m.Config.DynamoDBEndpoint)
	sess, err := session.NewSession(configDynamodb)
	if err != nil {
		return nil, err
	}

	return dynamodb.New(sess), nil
}

func (m *Migration) createSchemas(filename string) ([]*Schema, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	schemas := []*Schema{}
	if err := yaml.Unmarshal(data, &schemas); err != nil {
		return nil, err
	}

	return schemas, nil
}

func (m *Migration) makeCreateTableInput(schema *Schema) *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName:            aws.String(schema.Table),
		AttributeDefinitions: toAttributeDefinitions(schema.Attr),
		KeySchema:            toKeySchema(schema.Key),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(schema.Provisioned.Read),
			WriteCapacityUnits: aws.Int64(schema.Provisioned.Write),
		},
		GlobalSecondaryIndexes: toGlobalSecondaryIndex(schema.GlobalIndex),
		LocalSecondaryIndexes:  toLocalSecondaryIndex(schema.LocalIndex),
	}
}

func (m *Migration) makeDeleteTableInput(schema *Schema) *dynamodb.DeleteTableInput {
	return &dynamodb.DeleteTableInput{
		TableName: aws.String(schema.Table),
	}
}

func (m *Migration) Up() error {
	for _, schema := range m.schemas {
		input := m.makeCreateTableInput(schema)
		_, err := m.dynamo.CreateTable(input)
		if err != nil {
			aerr := err.(awserr.Error)
			if aerr.Code() != dynamodb.ErrCodeResourceInUseException {
				return err
			}
		}
	}

	return nil
}

func (m *Migration) Down() error {
	for _, schema := range m.schemas {
		input := m.makeDeleteTableInput(schema)
		_, err := m.dynamo.DeleteTable(input)
		if err != nil {
			aerr := err.(awserr.Error)
			if aerr.Code() != dynamodb.ErrCodeResourceNotFoundException {
				return err
			}
		}
	}

	return nil
}
