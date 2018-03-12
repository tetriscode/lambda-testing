package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	log.Printf("Received body:%s", request.Body)
	pp, err := json.Marshal(request.PathParameters)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error1:" + err.Error(),
			StatusCode: 200,
		}, nil
	}
	qp, err := json.Marshal(request.QueryStringParameters)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error3:" + err.Error(),
			StatusCode: 200,
		}, nil
	}
	sv, err := json.Marshal(request.StageVariables)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error4:" + err.Error(),
			StatusCode: 200,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body: "Body:" + request.Body +
			" PathParams:" + string(pp) +
			" Path:" + request.Path +
			" ResourcePath:" + request.RequestContext.ResourcePath +
			" QueryPath:" + string(qp) +
			" StageVars:" + string(sv),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
