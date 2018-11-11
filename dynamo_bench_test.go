package dynamo_bench

import (
	"os"
	"testing"
)

func bench(b *testing.B, endpoint string) {
	b.Helper()

	db, err := ConnectDB(endpoint)
	if err != nil {
		b.Error(err)
	}

	CreateTable(db)

	err = CreateRecords(db, b.N)
	if err != nil {
		b.Error(err)
	}
}

func BenchmarkDynamoLocal(b *testing.B) {
	bench(b, os.Getenv("DYNAMOLOCAL"))
}

func BenchmarkLocalStack(b *testing.B) {
	bench(b, os.Getenv("LOCALSTACK"))
}