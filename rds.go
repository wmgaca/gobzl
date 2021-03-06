package main

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func ListRdsInstances(region string) {
	conn := rds.New(session.New(), &aws.Config{Region: aws.String(region)})

	resp, err := conn.DescribeDBInstances(nil)
	if err != nil {
		panic(err)
	}

	table := NewTable([]string{
		"ID", "NAME", "ENGINE", "CLASS", "ENDPOINT", "PORT",
	})

	for _, dbInstance := range resp.DBInstances {
		var address string

		if nil != dbInstance.Endpoint.Address {
			address = stringFromPointer(dbInstance.Endpoint.Address, "")
		} else {
			address = ""
		}

		table.Append([]string{
			stringFromPointer(dbInstance.DBInstanceIdentifier, ""),
			stringFromPointer(dbInstance.DBName, ""),
			stringFromPointer(dbInstance.Engine, ""),
			stringFromPointer(dbInstance.DBInstanceClass, ""),
			address,
			strconv.FormatInt(int64FromPointer(dbInstance.Endpoint.Port, 0), 10),
		})
	}

	table.Render()
}
