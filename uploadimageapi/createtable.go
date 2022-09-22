package meathods

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//this is the function to create table
func CreateTablee() {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("AKIATICND5ZYGYCNGVXY", "cMPxwo9iryAst6V97gB3UgTcsdmoZeVMRrWrrDpe", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})
	tableName := "Movies"
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Movieid"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Movieid"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Title"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		log.Fatal("Got error calling CreateTable:", err)
	}
	fmt.Println("Created the table", tableName)
}
