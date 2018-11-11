package dynamo_bench

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

type Table struct {
	PK string `dynamo:"PK,hash"`
	SK string `dynamo:"SK,range"`
}

// Config 接続設定
func Config(endpoint string) *aws.Config {
	return &aws.Config{
		Region: aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
		Endpoint: aws.String(endpoint),
	}
}

// ConnectDB DynamoDBに接続する
func ConnectDB(endpoint string) (*dynamo.DB, error) {
	sess, err := session.NewSession(Config(endpoint))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dynamo.New(sess), nil
}

// CreateTable DynamoDBにテーブルを作成する
func CreateTable(db *dynamo.DB) error {
	return db.CreateTable("Table", Table{}).Run()
}

// CreateRecords DynamoDBのテーブルにn回書き込む
func CreateRecords(db *dynamo.DB, n int) error {
	for i := 0; i < n; i++ {
		err := db.Table("Table").Put(&Table{
			PK: fmt.Sprintf("PK%d", i),
			SK: fmt.Sprintf("SK%d", i),
		}).Run()
		if err != nil {
			return err
		}
	}
	return nil
}