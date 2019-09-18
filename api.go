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

type InternalPrepareNewFacesReq struct {
	UserId string `json:"userId"`
}

func (req InternalPrepareNewFacesReq) String() string {
	return fmt.Sprintf("%#v", req)
}

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
	NewFaces        []InternalProfiles `json:"newFaces"`
	LastActionTime  int64              `json:"lastActionTime"`
	HowMuchPrepared int64              `json:"howMuchPrepared"`
}

func (resp InternalGetNewFacesResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type DiscoverRequest struct {
	UserId         *string `json:"userId"`
	AccessToken    *string `json:"accessToken"`
	Resolution     *string `json:"resolution"`
	LastActionTime *int64  `json:"lastActionTime"`
	Limit          *int    `json:"limit"`
	Filter         *Filter `json:"filter"`
}

func (req DiscoverRequest) String() string {
	u := "nil"
	if req.UserId != nil {
		u = *req.UserId
	}
	a := "nil"
	if req.AccessToken != nil {
		a = *req.AccessToken
	}
	r := "nil"
	if req.Resolution != nil {
		r = *req.Resolution
	}
	l := "nil"
	if req.LastActionTime != nil {
		l = fmt.Sprintf("%v", *req.LastActionTime)
	}
	li := "nil"
	if req.Limit != nil {
		li = fmt.Sprintf("%d", *req.Limit)
	}
	f := "nil"
	if req.Filter != nil {
		f = fmt.Sprintf("[%v]", *req.Filter)
	}
	return fmt.Sprintf("userId [%s], accessToken [%s], resolution [%s], lastActionTime [%s], limit [%s], filter [%s]",
		u, a, r, l, li, f)
}

type Filter struct {
	MinAge      *int `json:"minAge"`
	MaxAge      *int `json:"maxAge"`
	MaxDistance *int `json:"maxDistance"`
}

func (filter Filter) String() string {
	min := "nil"
	if filter.MinAge != nil {
		min = fmt.Sprintf("%d", *filter.MinAge)
	}
	max := "nil"
	if filter.MaxAge != nil {
		max = fmt.Sprintf("%d", *filter.MaxAge)
	}
	maxD := "nil"
	if filter.MaxDistance != nil {
		maxD = fmt.Sprintf("%d", *filter.MaxDistance)
	}
	return fmt.Sprintf("minAge [%s], maxAge [%s], maxDistance [%s]", min, max, maxD)
}

type GetLCRequest struct {
	AccessToken    *string `json:"accessToken"`
	Resolution     *string `json:"resolution"`
	LastActionTime *int64  `json:"lastActionTime"`
	Limit          *int    `json:"limit"`
	Source         *string `json:"source"`
	Filter         *Filter `json:"filter"`

	//inner fields (set on server side)
	UserId *string `json:"userId"`
}

func (req GetLCRequest) String() string {
	u := "nil"
	if req.UserId != nil {
		u = *req.UserId
	}
	a := "nil"
	if req.AccessToken != nil {
		a = *req.AccessToken
	}
	r := "nil"
	if req.Resolution != nil {
		r = *req.Resolution
	}
	l := "nil"
	if req.LastActionTime != nil {
		l = fmt.Sprintf("%v", *req.LastActionTime)
	}
	li := "nil"
	if req.Limit != nil {
		li = fmt.Sprintf("%d", *req.Limit)
	}
	s := "nil"
	if req.Source != nil {
		s = fmt.Sprintf("%s", *req.Source)
	}
	f := "nil"
	if req.Filter != nil {
		f = fmt.Sprintf("[%v]", *req.Filter)
	}
	return fmt.Sprintf("userId [%s], accessToken [%s], resolution [%s], lastActionTime [%s], limit [%s], source [%s], filter [%s]",
		u, a, r, l, li, s, f)
}

type InternalGetLCResp struct {
	Profiles       []InternalProfiles `json:"profiles"`
	AllProfilesNum int                `json:"allProfilesNum"`
	LastActionTime int64              `json:"lastActionTime"`
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
	UserId         string          `json:"userId"`
	Photos         []InternalPhoto `json:"photos"`
	Messages       []Message       `json:"messages"`
	LastOnlineTime int64           `json:"lastOnlineTime"`
	LocationExist  bool            `json:"locationExist"`
	Lat            float64         `json:"lat"`
	Lon            float64         `json:"lon"`
	SourceLat      float64         `json:"slat"`
	SourceLon      float64         `json:"slon"`
	SourceLocale   string          `json:"slocale"`
	Age            int             `json:"age"`
	Sex            string          `json:"sex"`
	Property       int             `json:"property"`
	Transport      int             `json:"transport"`
	Income         int             `json:"income"`
	Height         int             `json:"height"`
	EducationLevel int             `json:"educationLevel"`
	HairColor      int             `json:"hairColor"`
	Children       int             `json:"children"`
	Name           string          `json:"name"`
	JobTitle       string          `json:"jobTitle"`
	Company        string          `json:"company"`
	EducationText  string          `json:"education"`
	About          string          `json:"about"`
	Instagram      string          `json:"instagram"`
	TikTok         string          `json:"tikTok"`
	WhereLive      string          `json:"whereLive"`
	WhereFrom      string          `json:"whereFrom"`
	Unseen         bool            `json:"unseen"`
	StatusText     string          `json:"statusText"`
	TotalLikes     int64           `json:"totalLikes"`

	//These properties needs for debug mode
	TotalScores          int `json:"totalScores"`
	TotalChatCount       int `json:"totalChatCount"`
	TotalChatCountScores int `json:"totalChatCountScores"`

	TotalMatchesCount       int `json:"totalMatchesCount"`
	TotalMatchesCountScores int `json:"totalMatchesCountScores"`

	PhotosCount       int `json:"photosCount"`
	PhotosCountScores int `json:"photosCountScores"`

	IncomeScores   int `json:"incomeScores"`
	ChildrenScores int `json:"childrenScores"`
	EduScores      int `json:"eduScores"`
	CityScores     int `json:"cityScores"`
	JobTitleScore  int `json:"jobTitleScore"`
	CompanyScores  int `json:"companyScores"`
	StatusScores   int `json:"statusScores"`
	NameScores     int `json:"nameScores"`
}

func (resp InternalProfiles) String() string {
	return fmt.Sprintf("%#v", resp)
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
	LastOnlineText              string    `json:"lastOnlineText"`
	LastOnlineFlag              string    `json:"lastOnlineFlag"`
	DistanceText                string    `json:"distanceText"`
	Unseen                      bool      `json:"notSeen"`
	Photos                      []Photo   `json:"photos"`
	Messages                    []Message `json:"messages"`
	Age                         int       `json:"age"`
	Sex                         string    `json:"sex"`
	Property                    int       `json:"property"`
	Transport                   int       `json:"transport"`
	Income                      int       `json:"income"`
	Height                      int       `json:"height"`
	EducationLevel              int       `json:"educationLevel"`
	HairColor                   int       `json:"hairColor"`
	Children                    int       `json:"children"`
	Name                        string    `json:"name"`
	JobTitle                    string    `json:"jobTitle"`
	Company                     string    `json:"company"`
	EducationText               string    `json:"education"`
	About                       string    `json:"about"`
	Instagram                   string    `json:"instagram"`
	TikTok                      string    `json:"tikTok"`
	WhereLive                   string    `json:"whereLive"`
	WhereFrom                   string    `json:"whereFrom"`
	StatusText                  string    `json:"statusText"`
	TotalLikes                  int64     `json:"totalLikes"`
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
	HaveBeenRead    bool   `json:"haveBeenRead"`
	WasYouSender    bool   `json:"wasYouSender"`
	Text            string `json:"text"`
	MessageId       string `json:"msgId"`
	ClientMessageId string `json:"clientMsgId"`
	MessageAt       int64  `json:"msgAt"`
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
	NewLikeEnabled    bool   `json:"newLikeEnabled"`
	NewMatchEnabled   bool   `json:"newMatchEnabled"`
	NewMessageEnabled bool   `json:"newMessageEnabled"`
	OppositeUserId    string `json:"oppositeUserId"`
}

func (resp PushObject) String() string {
	return fmt.Sprintf("%#v", resp)
}

type InternalChatRequest struct {
	UserId         string `json:"userId"`
	OppositeUserId string `json:"oppositeUserId"`
	LastActionTime int64  `json:"requestedLastActionTime"`
	Resolution     string `json:"resolution"`
}

func (resp InternalChatRequest) String() string {
	return fmt.Sprintf("%#v", resp)
}

type InternalChatResponse struct {
	Profile        InternalProfiles `json:"profile"`
	IsChatExists   bool             `json:"chatExists"`
	LastActionTime int64            `json:"lastActionTime"`
}

func (resp InternalChatResponse) String() string {
	return fmt.Sprintf("%#v", resp)
}
