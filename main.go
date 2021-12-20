package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/likexian/whois"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lookupInfo := ""
	if tmpName, ok := request.QueryStringParameters["name"]; ok {
		lookupInfo = tmpName
	} else {
		return events.APIGatewayProxyResponse{Body: "missing lookupInfo in path", StatusCode: http.StatusBadRequest}, nil
	}

	result, err := whois.Whois(lookupInfo)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("unable to get ip info: %s", err), StatusCode: http.StatusBadRequest}, nil
	}

	mapData := createMapData(result)
	obj := buildIPData(mapData)
	returnData, err := json.Marshal(obj)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("failed to parse ip info: %s", err), StatusCode: http.StatusNotAcceptable}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(returnData), StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}
