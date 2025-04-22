package dynadump

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"os"
)

func DumpTable(tableName string) error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return fmt.Errorf("unable to load SDK config: %w", err)
	}
	dynamo := dynamodb.NewFromConfig(cfg)
	scanPaginator := dynamodb.NewScanPaginator(dynamo, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	jsonEncoder := json.NewEncoder(os.Stdout)
	for scanPaginator.HasMorePages() {
		page, err := scanPaginator.NextPage(context.Background())
		if err != nil {
			return fmt.Errorf("unable to scan next page: %w", err)
		}
		for _, item := range page.Items {
			var row map[string]any
			err := attributevalue.UnmarshalMap(item, &row)
			if err != nil {
				return fmt.Errorf("unable to decode dynamodb item: %w", err)
			}
			err = jsonEncoder.Encode(row)
			if err != nil {
				return fmt.Errorf("unable to encode dynamodb item as json: %w", err)
			}
		}
	}
	return nil
}
