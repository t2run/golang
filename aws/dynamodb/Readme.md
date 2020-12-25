

#### Query with GSI and Filter

```Go

	filter := expression.Name("nameofcolumn/field").Equal("value")
  //Output fields required
	proj := expression.NamesList(expression.Name(""), expression.Name(""), expression.Name(""),
		expression.Name(""), expression.Name(""), expression.Name(""))
	expr, errExpression := expression.NewBuilder().
		WithFilter(filter).
		WithProjection(proj).
		Build()
	if errExpression != nil {
		errorString := "Query Expression Error" + errExpression.Error()
		logger.Error(errorString)
		return ncrStores, errExpression
	}

	var queryInput = &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String("tablename"),
		IndexName:                 aws.String("GSI Name"),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsipartitionkey": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(brand),
					},
				},
			},
			"sortkey(optional)": {  //to use incase additional filter param is unable to add becasue of sort key
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("true"),
					},
				},
			},
		},
		FilterExpression:     expr.Filter(),
		ProjectionExpression: expr.Projection(),
	}
 ```
 
  
  
