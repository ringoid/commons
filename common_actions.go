package commons

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
	"strconv"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"fmt"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/s3"
)

//return userId, sessionToken, ok, error string
func DecodeToken(token, secretWord string, anlogger *Logger, lc *lambdacontext.LambdaContext) (string, string, bool, string) {

	receiveToken, err := jwt.Parse(token, func(rt *jwt.Token) (interface{}, error) {
		if _, ok := rt.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("common_action.go : unexpected signing method: %v", rt.Header["alg"])
		}
		return []byte(secretWord), nil
	})
	if err != nil {
		anlogger.Warnf(lc, "common_action.go : error parse access token [%s] : %v", token, err)
		return "", "", false, InternalServerError
	}

	if claims, ok := receiveToken.Claims.(jwt.MapClaims); ok && receiveToken.Valid {
		userId := fmt.Sprintf("%v", claims[AccessTokenUserIdClaim])
		sessionToken := fmt.Sprintf("%v", claims[AccessTokenSessionTokenClaim])
		anlogger.Debugf(lc, "common_action.go : successfully parse access token, userId [%s], sessionToken [%s]", userId, sessionToken)
		return userId, sessionToken, true, ""
	} else {
		anlogger.Warnf(lc, "common_action.go : access token [%s] is not valid", token)
		return "", "", false, InternalServerError
	}
}

//return is session valid, ok, error string
func IsSessionValid(userId, sessionToken, userProfileTableName string, awsDbClient *dynamodb.DynamoDB,
	anlogger *Logger, lc *lambdacontext.LambdaContext) (bool, bool, string) {

	anlogger.Debugf(lc, "common_action.go : check that sessionToken [%s] is valid for userId [%s]", sessionToken, userId)
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			UserIdColumnName: {
				S: aws.String(userId),
			},
		},
		ConsistentRead: aws.Bool(true),
		TableName:      aws.String(userProfileTableName),
	}

	result, err := awsDbClient.GetItem(input)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error getting userInfo for userId [%s] : %v", userId, err)
		return false, false, InternalServerError
	}

	if len(result.Item) == 0 {
		anlogger.Warnf(lc, "common_action.go : there is no user with such userId [%s], sessionToken [%s]", userId, sessionToken)
		return false, true, ""
	}

	lastSessionToken := *result.Item[SessionTokenColumnName].S
	if sessionToken != lastSessionToken {
		anlogger.Warnf(lc, "common_action.go : sessionToken [%s] expired for userId [%s]", sessionToken, userId)
		return false, true, ""
	}

	anlogger.Debugf(lc, "common_action.go : session token is valid for userId [%s]", userId)
	return true, true, ""
}

func SendAnalyticEvent(event interface{}, userId, deliveryStreamName string, awsDeliveryStreamClient *firehose.Firehose,
	anlogger *Logger, lc *lambdacontext.LambdaContext) {
	anlogger.Debugf(lc, "common_action.go : send analytics event [%v] for userId [%s]", event, userId)
	data, err := json.Marshal(event)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error marshaling analytics event [%v] for userId [%s] : %v", event, userId, err)
		return
	}
	newLine := "\n"
	data = append(data, newLine...)
	_, err = awsDeliveryStreamClient.PutRecord(&firehose.PutRecordInput{
		DeliveryStreamName: aws.String(deliveryStreamName),
		Record: &firehose.Record{
			Data: data,
		},
	})

	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error sending analytics event [%v] for userId [%s] : %v", event, userId, err)
	}

	anlogger.Debugf(lc, "common_action.go : successfully send analytics event [%v] for userId [%s]", event, userId)
}

//ok and error string
func SendCommonEvent(event interface{}, userId, commonStreamName, partitionKey string, awsKinesisClient *kinesis.Kinesis,
	anlogger *Logger, lc *lambdacontext.LambdaContext) (bool, string) {
	anlogger.Debugf(lc, "common_action.go : send common event [%v] for userId [%s]", event, userId)
	data, err := json.Marshal(event)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error marshaling common event [%v] for userId [%s] : %v", event, userId, err)
		return false, InternalServerError
	}
	if len(partitionKey) == 0 {
		partitionKey = userId
	}
	input := &kinesis.PutRecordInput{
		StreamName:   aws.String(commonStreamName),
		PartitionKey: aws.String(partitionKey),
		Data:         []byte(data),
	}
	_, err = awsKinesisClient.PutRecord(input)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error putting common event into stream, event [%v] for userId [%s] : %v", event, userId, err)
		return false, InternalServerError
	}
	anlogger.Debugf(lc, "common_action.go : successfully send common event [%v] for userId [%s]", event, userId)
	return true, ""
}

func GetSecret(secretBase, secretKeyName string, awsSession *session.Session, anlogger *Logger, lc *lambdacontext.LambdaContext) string {
	anlogger.Debugf(lc, "lambda-initialization : common_action.go : read secret with secretBase [%s], secretKeyName [%s]", secretBase, secretKeyName)
	svc := secretsmanager.New(awsSession)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretBase),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		anlogger.Fatalf(lc, "lambda-initialization : common_action.go : error reading %s secret from Secret Manager : %v", secretBase, err)
	}
	var secretMap map[string]string
	decoder := json.NewDecoder(strings.NewReader(*result.SecretString))
	err = decoder.Decode(&secretMap)
	if err != nil {
		anlogger.Fatalf(lc, "lambda-initialization : common_action.go : error decode %s secret from Secret Manager : %v", secretBase, err)
	}
	secret, ok := secretMap[secretKeyName]
	if !ok {
		anlogger.Fatalf(lc, "lambda-initialization : common_action.go : secret %s is empty", secretBase)
	}
	anlogger.Debugf(lc, "lambda-initialization : common_action.go : secret %s was successfully initialized", secretBase)

	return secret
}

func WarmUpLambda(functionName string, clientLambda *lambda.Lambda, anlogger *Logger, lc *lambdacontext.LambdaContext) {
	anlogger.Debugf(lc, "common_action.go : warmup lambda [%s]", functionName)
	req := WarmUpRequest{
		WarmUpRequest: true,
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error marshaling req %v into json : %v", req, err)
		return
	}

	apiReq := events.APIGatewayProxyRequest{
		Body: string(jsonBody),
	}

	apiJsonBody, err := json.Marshal(apiReq)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error marshaling req %v into json : %v", apiReq, err)
		return
	}

	payload := apiJsonBody
	if strings.Contains(functionName, "internal") {
		payload = jsonBody
	}
	_, err = clientLambda.Invoke(&lambda.InvokeInput{FunctionName: aws.String(functionName), Payload: payload, InvocationType: aws.String("Event")})

	if err != nil {
		anlogger.Errorf(lc, "warm_up.go : error invoke function [%s] with body %s : %v", functionName, string(payload), err)
		return
	}

	anlogger.Debugf(lc, "common_action.go : successfully warmup lambda [%s]", functionName)
	return
}

func IsItWarmUpRequest(body string, anlogger *Logger, lc *lambdacontext.LambdaContext) bool {
	anlogger.Debugf(lc, "common_action.go : is it warm up request, body [%s]", body)
	if len(body) == 0 {
		anlogger.Debugf(lc, "common_action.go : empty request body, it's no warm up request")
		return false
	}
	var req WarmUpRequest
	err := json.Unmarshal([]byte(body), &req)

	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error unmarshal required params from the string [%s] : %v", body, err)
		return false
	}
	result := req.WarmUpRequest
	anlogger.Debugf(lc, "common_action.go : successfully check that it's warm up request, body [%s], result [%v]", body, result)
	return result
}

//return ok and error string
func UpdateLastOnlineTimeAndBuildNum(userId, userProfileTableName string, buildNum int, isItAndroid bool,
	awsDbClient *dynamodb.DynamoDB, anlogger *Logger, lc *lambdacontext.LambdaContext) (bool, string) {
	anlogger.Debugf(lc, "common_action.go : update last online time and build num [%v] (is it andoid [%v]) for userId [%s]",
		buildNum, isItAndroid, userId)

	columnName := CurrentAndroidBuildNum
	if !isItAndroid {
		columnName = CurrentiOSBuildNum
	}

	input :=
		&dynamodb.UpdateItemInput{
			ExpressionAttributeNames: map[string]*string{
				"#onlineTime": aws.String(LastOnlineTimeColumnName),
				"#buildNum":   aws.String(columnName),
			},
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":onlineTimeV": {
					N: aws.String(fmt.Sprintf("%v", time.Now().Unix())),
				},
				":buildNumV": {
					N: aws.String(strconv.Itoa(buildNum)),
				},
			},
			Key: map[string]*dynamodb.AttributeValue{
				UserIdColumnName: {
					S: aws.String(userId),
				},
			},
			TableName:        aws.String(userProfileTableName),
			UpdateExpression: aws.String("SET #onlineTime = :onlineTimeV, #buildNum = :buildNumV"),
		}

	_, err := awsDbClient.UpdateItem(input)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error while update last online time and build num [%v] (is it android [%v]) for userId [%s] : %v",
			userId, buildNum, isItAndroid, err)
		return false, InternalServerError
	}

	anlogger.Debugf(lc, "common_action.go : successfully update last online time and build num [%v] (is it android [%v]) for userId [%s]",
		buildNum, isItAndroid, userId)

	return true, ""
}

//return userId, ok and error string
func Login(appVersion int, isItAndroid bool, token, secretWord, userProfileTable, commonStreamName string, awsDbClient *dynamodb.DynamoDB, awsKinesisClient *kinesis.Kinesis,
	anlogger *Logger, lc *lambdacontext.LambdaContext) (string, bool, string) {

	anlogger.Debugf(lc, "common_action.go : login for token [%s] with app version [%d] and isItAndroid [%v]", token, appVersion, isItAndroid)

	switch isItAndroid {
	case true:
		if appVersion < MinimalAndroidBuildNum {
			anlogger.Infof(lc, "common_action.go : too old Android version [%d] when min version is [%d]", appVersion, MinimalAndroidBuildNum)
			return "", false, TooOldAppVersionClientError
		}
	default:
		if appVersion < MinimaliOSBuildNum {
			anlogger.Infof(lc, "common_action.go : too old iOS version [%d] when min version is [%d]", appVersion, MinimaliOSBuildNum)
			return "", false, TooOldAppVersionClientError
		}
	}

	userId, sessionToken, ok, errStr := DecodeToken(token, secretWord, anlogger, lc)
	if !ok {
		return "", ok, errStr
	}

	valid, ok, errStr := IsSessionValid(userId, sessionToken, userProfileTable, awsDbClient, anlogger, lc)
	if !ok {
		return "", ok, errStr
	}

	if !valid {
		return "", false, InvalidAccessTokenClientError
	}

	ok, errStr = UpdateLastOnlineTimeAndBuildNum(userId, userProfileTable, appVersion, isItAndroid, awsDbClient, anlogger, lc)
	if !ok {
		return "", ok, errStr
	}

	event := NewUserOnlineEvent(userId)
	partitionKey := userId
	ok, errStr = SendCommonEvent(event, userId, commonStreamName, partitionKey, awsKinesisClient, anlogger, lc)
	if !ok {
		return "", ok, errStr
	}

	anlogger.Debugf(lc, "common_action.go : successfully login for token [%s] with app version [%d]", token, appVersion)
	return userId, true, ""
}

//return ok and error string
func SendAsyncTask(task interface{}, asyncTaskQueue, userId string, messageSecDelay int64,
	awsSqsClient *sqs.SQS, anlogger *Logger, lc *lambdacontext.LambdaContext) (bool, string) {
	anlogger.Debugf(lc, "common_action.go : send async task %v for userId [%s] with delay in sec [%v]", task, userId, messageSecDelay)
	body, err := json.Marshal(task)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error marshal task %v for userId [%s] with delay in sec [%v] : %v", task, userId, messageSecDelay, err)
		return false, InternalServerError
	}
	input := &sqs.SendMessageInput{
		DelaySeconds: aws.Int64(messageSecDelay),
		QueueUrl:     aws.String(asyncTaskQueue),
		MessageBody:  aws.String(string(body)),
	}
	_, err = awsSqsClient.SendMessage(input)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error sending async task %v to the queue for userId [%s] with delay in sec [%v] : %v", task, userId, messageSecDelay, err)
		return false, InternalServerError
	}
	anlogger.Debugf(lc, "common_action.go : successfully send async task %v for userId [%s] with delay in sec [%v]", task, userId, messageSecDelay)
	return true, ""
}

//return ok and error string
func SendCloudWatchMetric(baseCloudWatchNamespace, metricName string, value int, cwClient *cloudwatch.CloudWatch, anlogger *Logger, lc *lambdacontext.LambdaContext) (bool, string) {
	anlogger.Debugf(lc, "common_action.go : send value [%d] for namespace [%s] and metric name [%s]", value, baseCloudWatchNamespace, metricName)

	currentTime := time.Now().UTC()

	peD := cloudwatch.MetricDatum{
		MetricName: aws.String(metricName),
		Timestamp:  &currentTime,
		Value:      aws.Float64(float64(value)),
	}

	metricdatas := []*cloudwatch.MetricDatum{&peD}

	_, err := cwClient.PutMetricData(&cloudwatch.PutMetricDataInput{
		MetricData: metricdatas,
		Namespace:  aws.String(baseCloudWatchNamespace),
	})
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error sending cloudwatch metric with value [%d] for namespace [%s] and metric name [%s] : %v", value, baseCloudWatchNamespace, metricName, err)
		return false, InternalServerError
	}

	anlogger.Debugf(lc, "common_action.go : successfully send value [%d] for namespace [%s] and metric name [%s]", value, baseCloudWatchNamespace, metricName)
	return true, ""
}

//return buildnum, is it android, ok and error string
func ParseAppVersionFromHeaders(headers map[string]string, anlogger *Logger, lc *lambdacontext.LambdaContext) (int, bool, bool, string) {
	anlogger.Debugf(lc, "common_action.go : parse build num from the headers %v", headers)
	var appVersionInt int
	var err error

	if appVersionStr, ok := headers[AndroidBuildNum]; ok {
		appVersionInt, err = strconv.Atoi(appVersionStr)
		if err != nil {
			anlogger.Errorf(lc, "common_action.go : error converting header [%s] with value [%s] to int : %v", AndroidBuildNum, appVersionStr, err)
			return 0, false, false, WrongRequestParamsClientError
		}
		anlogger.Debugf(lc, "common_action.go : successfully parse Android build num [%d] from the headers %v", appVersionInt, headers)
		return appVersionInt, true, true, ""

	} else if appVersionStr, ok = headers[iOSdBuildNum]; ok {
		appVersionInt, err = strconv.Atoi(appVersionStr)
		if err != nil {
			anlogger.Errorf(lc, "common_action.go : error converting header [%s] with value [%s] to int : %v", iOSdBuildNum, appVersionStr, err)
			return 0, false, false, WrongRequestParamsClientError
		}
		anlogger.Debugf(lc, "common_action.go : successfully parse iOS build num [%d] from the headers %v", appVersionInt, headers)
		return appVersionInt, false, true, ""
	} else {
		anlogger.Errorf(lc, "common_action.go : error header [%s] is empty", AndroidBuildNum)
		return 0, false, false, WrongRequestParamsClientError
	}
}

//return userId, ok, was user reported, error string
func CallVerifyAccessToken(buildNum int, isItAndroid bool, accessToken, functionName string, clientLambda *lambda.Lambda, anlogger *Logger, lc *lambdacontext.LambdaContext) (string, bool, bool, string) {
	req := InternalGetUserIdReq{
		AccessToken: accessToken,
		BuildNum:    buildNum,
		IsItAndroid: isItAndroid,
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error marshaling req %s into json : %v", req, err)
		return "", false, false, InternalServerError
	}

	resp, err := clientLambda.Invoke(&lambda.InvokeInput{FunctionName: aws.String(functionName), Payload: jsonBody})
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error invoke function [%s] with body %s : %v", functionName, jsonBody, err)
		return "", false, false, InternalServerError
	}

	if *resp.StatusCode != 200 {
		anlogger.Errorf(lc, "common_action.go : status code = %d, response body %s for request %s", *resp.StatusCode, string(resp.Payload), jsonBody)
		return "", false, false, InternalServerError
	}

	var response InternalGetUserIdResp
	err = json.Unmarshal(resp.Payload, &response)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error unmarshaling response %s into json : %v", string(resp.Payload), err)
		return "", false, false, InternalServerError
	}

	if response.ErrorCode != "" {
		anlogger.Errorf(lc, "common_action.go : error response from function [%s], response=%v", functionName, response)
		switch response.ErrorCode {
		case "InvalidAccessTokenClientError":
			return "", false, false, InvalidAccessTokenClientError
		case "TooOldAppVersionClientError":
			return "", false, false, TooOldAppVersionClientError
		default:
			return "", false, false, InternalServerError
		}
	}

	anlogger.Debugf(lc, "common_action.go : successfully validate accessToken, userId [%s]", response.UserId)
	return response.UserId, true, response.IsUserReported, ""
}

//return ok and error string
func DeleteFromS3(bucket, key, userId string, awsS3Client *s3.S3, lc *lambdacontext.LambdaContext, anlogger *Logger) (bool, string) {
	anlogger.Debugf(lc, "common_action.go : delete from s3 bucket [%s] with key [%s] for userId [%s]",
		bucket, key, userId)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err := awsS3Client.DeleteObject(input)
	if err != nil {
		anlogger.Errorf(lc, "common_action.go : error delete from s3 bucket [%s] with key [%s] for userId [%s] : %v",
			bucket, key, userId, err)
		return false, InternalServerError
	}

	anlogger.Debugf(lc, "common_action.go : successfully delete from s3 bucket [%s] with key [%s] for userId [%s]",
		bucket, key, userId)
	return true, ""
}

func GetOriginPhotoId(userId, sourcePhotoId string, anlogger *Logger, lc *lambdacontext.LambdaContext) (string, bool) {
	anlogger.Debugf(lc, "common_action.go : get origin photo id based on source photo id [%s] for userId [%s]", sourcePhotoId, userId)
	if len(sourcePhotoId) == 0 {
		anlogger.Warnf(lc, "common_action.go : empty source photo id for userId [%s]", userId)
		return "", false
	}
	arr := strings.Split(sourcePhotoId, "_")
	if len(arr) != 2 {
		anlogger.Warnf(lc, "common_action.go : wrong source photo id [%s] for userId [%s]", sourcePhotoId, userId)
		return "", false
	}
	baseId := arr[1]
	originPhotoId := "origin_" + baseId
	anlogger.Debugf(lc, "common_action.go : successfully get origin photo id [%s] for source photo id [%s] for userId [%s]",
		originPhotoId, sourcePhotoId, userId)
	return originPhotoId, true
}

func GetResolutionPhotoId(userId, originPhotoId, resolution string, anlogger *Logger, lc *lambdacontext.LambdaContext) (string, bool) {
	anlogger.Debugf(lc, "common_action.go : get resolution [%s] photo id based on origin photo id [%s] for userId [%s]", resolution, originPhotoId, userId)
	if len(originPhotoId) == 0 {
		anlogger.Warnf(lc, "common_action.go : empty origin photo id for userId [%s]", userId)
		return "", false
	}
	arr := strings.Split(originPhotoId, "_")
	if len(arr) != 2 {
		anlogger.Warnf(lc, "common_action.go : wrong origin photo id [%s] for userId [%s]", originPhotoId, userId)
		return "", false
	}
	baseId := arr[1]
	resolutionPhotoId := resolution + "_" + baseId
	anlogger.Debugf(lc, "common_action.go : successfully get resolution [%s] photo id [%s] for origin photo id [%s] for userId [%s]",
		resolution, resolutionPhotoId, originPhotoId, userId)
	return resolutionPhotoId, true
}
