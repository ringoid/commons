package commons

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type BaseResponse struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (resp BaseResponse) String() string {
	return fmt.Sprintf("%#v", resp)
}

type WarmUpRequest struct {
	WarmUpRequest bool `json:"warmUpRequest"`
}

func (req WarmUpRequest) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalGetUserIdReq struct {
	WarmUpRequest bool   `json:"warmUpRequest"`
	AccessToken   string `json:"accessToken"`
	BuildNum      int    `json:"buildNum"`
	IsItAndroid   bool   `json:"isItAndroid"`
}

func (req InternalGetUserIdReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalGetUserIdResp struct {
	BaseResponse
	UserId               string `json:"userId"`
	UserTakePartInReport bool   `json:"userTakePartInReport"`
}

func (resp InternalGetUserIdResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

//Feeds - Internal communication with Relationships service

type InternalGetNewFacesReq struct {
	WarmUpRequest  bool   `json:"warmUpRequest"`
	UserId         string `json:"userId"`
	Limit          int    `json:"limit"`
	LastActionTime int64  `json:"requestedLastActionTime"`
	Resolution     string `json:"resolution"`
}

func (req InternalGetNewFacesReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalGetNewFacesResp struct {
	NewFaces       []InternalProfiles `json:"newFaces"`
	LastActionTime int64              `json:"lastActionTime"`
}

func (resp InternalGetNewFacesResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type InternalLMMReq struct {
	WarmUpRequest           bool   `json:"warmUpRequest"`
	UserId                  string `json:"userId"`
	RequestNewPart          bool   `json:"requestNewPart"`
	RequestedLastActionTime int64  `json:"requestedLastActionTime"`
	Resolution              string `json:"resolution"`
}

func (req InternalLMMReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalLMHISReq struct {
	WarmUpRequest           bool   `json:"warmUpRequest"`
	UserId                  string `json:"userId"`
	RequestNewPart          bool   `json:"requestNewPart"`
	RequestedLastActionTime int64  `json:"requestedLastActionTime"`
	Resolution              string `json:"resolution"`
	LMHISPart               string `json:"lmhisPart"` //hellos | inbox | sent
}

func (req InternalLMHISReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalLMMResp struct {
	Profiles       []InternalProfiles `json:"profiles"`
	LastActionTime int64              `json:"lastActionTime"`
}

type InternalProfiles struct {
	UserId   string          `json:"userId"`
	Photos   []InternalPhoto `json:"photos"`
	Messages []Message       `json:"messages"`
}

func (resp InternalLMMResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type InternalLMHISResp struct {
	Profiles       []InternalProfiles `json:"profiles"`
	LastActionTime int64              `json:"lastActionTime"`
}

func (resp InternalLMHISResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

//
//private String originPhotoId;
//private String resizedPhotoId;
//private String link;
//private String resolution;

type InternalPhoto struct {
	OriginPhotoId  string `json:"originPhotoId"`
	ResizedPhotoId string `json:"resizedPhotoId"`
	Link           string `json:"link"`
	Resolution     string `json:"resolution"`
	ThumbnailLink  string `json:"thumbnailLink"`
}

func (resp InternalPhoto) String() string {
	return fmt.Sprintf("%#v", resp)
}

//Feeds - Images communications

type ProfilesResp struct {
	BaseResponse
	WarmUpRequest bool      `json:"warmUpRequest"`
	Profiles      []Profile `json:"profiles"`
}

type Profile struct {
	UserId                      string    `json:"userId"`
	DefaultSortingOrderPosition int       `json:"defaultSortingOrderPosition"`
	Unseen                      bool      `json:"notSeen"`
	Photos                      []Photo   `json:"photos"`
	Messages                    []Message `json:"messages"`
}

type Photo struct {
	PhotoId           string `json:"photoId"`
	PhotoUri          string `json:"photoUri"`
	ThumbnailPhotoUri string `json:"thumbnailPhotoUri"`
}

func (resp ProfilesResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type FacesWithUrlResp struct {
	//contains userId_photoId like a key and photoUrl like a value
	UserIdPhotoIdKeyUrlMap map[string]string `json:"urlPhotoMap"`
}

func (resp FacesWithUrlResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

// Feeds - Messages

type InternalGetMessagesReq struct {
	WarmUpRequest bool     `json:"warmUpRequest"`
	SourceUserId  string   `json:"sourceUserId"`
	TargetUserIds []string `json:"targetUserIds"`
}

func (req InternalGetMessagesReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalGetMessagesResp struct {
	//key = userId on the other side of conversation
	ConversationsMap map[string][]Message `json:"conversationsMap"`
}

func (resp InternalGetMessagesResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type Message struct {
	WasYouSender bool   `json:"wasYouSender"`
	Text         string `json:"text"`
}

func NewServiceResponse(body string) events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode:        200,
		StatusDescription: "HTTP OK",
		Headers:           map[string]string{"Content-Type": "application/json"},
		Body:              body}
}

func NewWrongHttpMethodServiceResponse() events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode:        400,
		StatusDescription: "Wrong HTTP method used",
		Headers:           map[string]string{"Content-Type": "application/json"},
		Body:              "{}"}
}

type PushRequest struct {
	Skip                int64 `json:"skip"`
	Limit               int64 `json:"limit"`
	MaxPeriod           int64 `json:"maxPeriod"`
	OfflinePeriod       int64 `json:"offlinePeriod"`
	MinProfilesForMen   int64 `json:"minForMen"`
	MinProfilesForWomen int64 `json:"minForWomen"`
	MinH                int64 `json:"minH"`
	MaxH                int64 `json:"maxH"`
}

func (resp PushRequest) String() string {
	return fmt.Sprintf("%#v", resp)
}

type PushResponse struct {
	Users       []PushObject `json:"users"`
	ResultCount int64        `json:"resultCount"`
}

func (resp PushResponse) String() string {
	return fmt.Sprintf("%#v", resp)
}

const (
	OnceDayPushType    = "ONCE_A_DAY_PUSH_TYPE"
	NewLikePushType    = "NEW_LIKE_PUSH_TYPE"
	NewMatchPushType   = "NEW_MATCH_PUSH_TYPE"
	NewMessagePushType = "NEW_MESSAGE_PUSH_TYPE"
)

type PushObject struct {
	UserId            string `json:"userId"`
	Sex               string `json:"sex"`
	Locale            string `json:"locale"`
	LastOnlineTime    int64  `json:"lastOnlineTime"`
	NewMessageCounter int64  `json:"newMessageCount"`
	NewMatchCounter   int64  `json:"newMatchCount"`
	NewLikeCounter    int64  `json:"newLikeCount"`
	NewProfileCounter int64  `json:"newProfiles"`
	PushType          string `json:"pushType"`
}

func (resp PushObject) String() string {
	return fmt.Sprintf("%#v", resp)
}
