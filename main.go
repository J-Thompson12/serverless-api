package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/likexian/whois"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type IPData struct {
	Name      string `json:"Name:"`
	City      string `json:"City"`
	StateProv string `json:"StateProv"`
	Email     string `json:"Email"`
}

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

func createMapData(data string) map[string]string {
	temp := strings.Split(data, "\n")
	result := make(map[string]string)
	for _, line := range temp {
		if !strings.Contains(line, "#") && strings.Contains(line, ":") {
			mapValues := strings.Split(line, ":")
			key := strings.TrimSpace(mapValues[0])
			value := strings.TrimSpace(mapValues[1])
			result[key] = value
		}
	}

	return result
}

func buildIPData(data map[string]string) IPData {
	if isDomain(data) {
		return IPData{
			Name:      data["Domain Name"],
			City:      data["Registrant City"],
			StateProv: data["Admin State/Province"],
			Email:     data["Admin Email"],
		}
	}
	return IPData{
		Name:      data["OrgName"],
		City:      data["City"],
		StateProv: data["StateProv"],
		Email:     data["OrgAbuseEmail"],
	}
}

func isDomain(data map[string]string) bool {
	if _, ok := data["Domain Name"]; ok {
		return true
	}
	return false
}
