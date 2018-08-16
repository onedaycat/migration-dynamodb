package migration

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func toAttributeDefinitions(nameTypes []NameType) []*dynamodb.AttributeDefinition {
	if len(nameTypes) == 0 {
		return nil
	}

	a := make([]*dynamodb.AttributeDefinition, len(nameTypes))
	for i, nameType := range nameTypes {
		a[i] = &dynamodb.AttributeDefinition{
			AttributeName: aws.String(nameType.Name),
			AttributeType: aws.String(nameType.Type),
		}
	}

	return a
}

func toKeySchema(nameTypes []NameType) []*dynamodb.KeySchemaElement {
	if len(nameTypes) == 0 {
		return nil
	}

	a := make([]*dynamodb.KeySchemaElement, len(nameTypes))
	for i, nameType := range nameTypes {
		a[i] = &dynamodb.KeySchemaElement{
			AttributeName: aws.String(nameType.Name),
			KeyType:       aws.String(nameType.Type),
		}
	}

	return a
}

func toGlobalSecondaryIndex(indexs []GlobalIndex) []*dynamodb.GlobalSecondaryIndex {
	if len(indexs) == 0 {
		return nil
	}

	a := make([]*dynamodb.GlobalSecondaryIndex, len(indexs))
	for i, index := range indexs {
		a[i] = &dynamodb.GlobalSecondaryIndex{
			IndexName: aws.String(index.Name),
			KeySchema: toKeySchema(index.Key),
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(index.Provisioned.Read),
				WriteCapacityUnits: aws.Int64(index.Provisioned.Write),
			},
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(index.Projection.Type),
			},
		}

		if len(index.Projection.Key) > 0 {
			a[i].Projection.NonKeyAttributes = aws.StringSlice(index.Projection.Key)
		}
	}

	return a
}

func toLocalSecondaryIndex(indexs []LocalIndex) []*dynamodb.LocalSecondaryIndex {
	if len(indexs) == 0 {
		return nil
	}

	a := make([]*dynamodb.LocalSecondaryIndex, len(indexs))
	for i, index := range indexs {
		a[i] = &dynamodb.LocalSecondaryIndex{
			IndexName: aws.String(index.Name),
			KeySchema: toKeySchema(index.Key),
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(index.Projection.Type),
			},
		}

		if len(index.Projection.Key) > 0 {
			a[i].Projection.NonKeyAttributes = aws.StringSlice(index.Projection.Key)
		}
	}

	return a
}
