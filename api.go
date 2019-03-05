package commons

import "fmt"

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
}

func (req InternalGetNewFacesReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalGetNewFacesResp struct {
	NewFaces       []InternalNewFace `json:"newFaces"`
	LastActionTime int64             `json:"lastActionTime"`
}

type InternalNewFace struct {
	UserId   string   `json:"userId"`
	PhotoIds []string `json:"photoIds"`
}

func (resp InternalGetNewFacesResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type InternalLMMReq struct {
	WarmUpRequest           bool   `json:"warmUpRequest"`
	UserId                  string `json:"userId"`
	RequestNewPart          bool   `json:"requestNewPart"`
	RequestedLastActionTime int64  `json:"requestedLastActionTime"`
}

func (req InternalLMMReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalLMMResp struct {
	Profiles       []InternalProfiles `json:"profiles"`
	LastActionTime int64              `json:"lastActionTime"`
}

type InternalProfiles struct {
	UserId   string    `json:"userId"`
	PhotoIds []string  `json:"photoIds"`
	Messages []Message `json:"messages"`
}

func (resp InternalLMMResp) String() string {
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
	PhotoId  string `json:"photoId"`
	PhotoUri string `json:"photoUri"`
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
