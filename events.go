package commons

import (
	"fmt"
)

type UserAcceptTermsEvent struct {
	UserId                     string `json:"userId"`
	SourceIp                   string `json:"sourceIp"`
	UnixTime                   int64  `json:"unixTime"`
	EventType                  string `json:"eventType"`
	DateTimeTermsAndConditions int64  `json:"dtTC"`
	DateTimePrivacyNotes       int64  `json:"dtPN"`
	DateTimeLegalAge           int64  `json:"dtLA"`
	CustomerId                 string `json:"customerId"`
	AndroidDeviceModel         string `json:"androidDeviceModel"`
	AndroidOsVersion           string `json:"androidOsVersion"`
	iOSDeviceModel             string `json:"iOsDeviceModel"`
	iOsVersion                 string `json:"iOsVersion"`
}

func (event UserAcceptTermsEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserAcceptTermsEvent(userId, customerId, sourceIp, deviceModel, osVersion string,
	dateTimeLegalAge, dateTimePrivacyNotes, dateTimeTermsAndConditions int64, isItAndroid bool) *UserAcceptTermsEvent {
	if isItAndroid {
		return &UserAcceptTermsEvent{
			UserId: userId,
			//gdpr?
			SourceIp:   sourceIp,
			CustomerId: customerId,

			UnixTime:                   UnixTimeInMillis(),
			DateTimeLegalAge:           dateTimeLegalAge,
			DateTimePrivacyNotes:       dateTimePrivacyNotes,
			DateTimeTermsAndConditions: dateTimeTermsAndConditions,
			AndroidDeviceModel:         deviceModel,
			AndroidOsVersion:           osVersion,
			EventType:                  "AUTH_USER_ACCEPT_TERMS",
		}
	}
	return &UserAcceptTermsEvent{
		UserId: userId,
		//gdpr?
		SourceIp:   sourceIp,
		CustomerId: customerId,

		UnixTime:                   UnixTimeInMillis(),
		DateTimeLegalAge:           dateTimeLegalAge,
		DateTimePrivacyNotes:       dateTimePrivacyNotes,
		DateTimeTermsAndConditions: dateTimeTermsAndConditions,
		iOSDeviceModel:             deviceModel,
		iOsVersion:                 osVersion,
		EventType:                  "AUTH_USER_ACCEPT_TERMS",
	}
}

type UserProfileCreatedEvent struct {
	UserId      string `json:"userId"`
	SourceIp    string `json:"sourceIp"`
	Sex         string `json:"sex"`
	YearOfBirth int    `json:"yearOfBirth"`
	ReferralId  string `json:"referralId"`
	PrivateKey  string `json:"privateKey"`
	UnixTime    int64  `json:"unixTime"`
	EventType   string `json:"eventType"`
}

func (event UserProfileCreatedEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserProfileCreatedEvent(userId, sex, sourceIp, referralId, privateKey string, yearOfBirth int) *UserProfileCreatedEvent {
	return &UserProfileCreatedEvent{
		UserId:      userId,
		SourceIp:    sourceIp,
		Sex:         sex,
		YearOfBirth: yearOfBirth,
		ReferralId:  referralId,
		PrivateKey:  privateKey,
		UnixTime:    UnixTimeInMillis(),
		EventType:   "AUTH_USER_PROFILE_CREATED",
	}
}

type UserClaimReferralCodeEvent struct {
	UserId     string `json:"userId"`
	SourceIp   string `json:"sourceIp"`
	ReferralId string `json:"referralId"`
	UnixTime   int64  `json:"unixTime"`
	EventType  string `json:"eventType"`
}

func (event UserClaimReferralCodeEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserClaimReferralCodeEvent(userId, sourceIp, referralId string) *UserClaimReferralCodeEvent {
	return &UserClaimReferralCodeEvent{
		UserId:     userId,
		SourceIp:   sourceIp,
		ReferralId: referralId,
		UnixTime:   UnixTimeInMillis(),
		EventType:  "AUTH_USER_CLAIM_REFERRAL_CODE",
	}
}

type UserSettingsUpdatedEvent struct {
	UserId             string `json:"userId"`
	SourceIp           string `json:"sourceIp"`
	Locale             string `json:"locale"`
	WasLocaleChanged   bool   `json:"wasLocaleChanged"`
	Push               bool   `json:"push"`
	WasPushChanged     bool   `json:"wasPushChanged"`
	TimeZone           int    `json:"timeZone"`
	WasTimeZoneChanged bool   `json:"wasTimeZoneChanged"`
	UnixTime           int64  `json:"unixTime"`
	EventType          string `json:"eventType"`
}

func (event UserSettingsUpdatedEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserSettingsUpdatedEvent(userId, sourceIp, locale string, wasLocaleChanged, push, wasPushChanged bool, timeZone int, wasTimeZoneChanged bool) *UserSettingsUpdatedEvent {
	return &UserSettingsUpdatedEvent{
		UserId:             userId,
		SourceIp:           sourceIp,
		Locale:             locale,
		WasLocaleChanged:   wasLocaleChanged,
		Push:               push,
		WasPushChanged:     wasPushChanged,
		TimeZone:           timeZone,
		WasTimeZoneChanged: wasTimeZoneChanged,
		UnixTime:           UnixTimeInMillis(),
		EventType:          "AUTH_USER_SETTINGS_UPDATED",
	}
}

//DEPRECATED
//todo:delete after 1.0.10 release
type GetUserSettingsEvent struct {
	UserId    string `json:"userId"`
	SourceIp  string `json:"sourceIp"`
	UnixTime  int64  `json:"unixTime"`
	EventType string `json:"eventType"`
}

func (event GetUserSettingsEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewGetUserSettingsEvent(userId, sourceIp string) *GetUserSettingsEvent {
	return &GetUserSettingsEvent{
		UserId:    userId,
		SourceIp:  sourceIp,
		UnixTime:  UnixTimeInMillis(),
		EventType: "AUTH_GET_USER_SETTINGS",
	}
}

type UserCallDeleteHimselfEvent struct {
	UserId           string `json:"userId"`
	SourceIp         string `json:"sourceIp"`
	UserReportStatus string `json:"userReportStatus"`
	UnixTime         int64  `json:"unixTime"`
	EventType        string `json:"eventType"`
}

func (event UserCallDeleteHimselfEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserCallDeleteHimselfEvent(userId, sourceIp string, userReportStatus string) *UserCallDeleteHimselfEvent {
	return &UserCallDeleteHimselfEvent{
		UserId:           userId,
		SourceIp:         sourceIp,
		UserReportStatus: userReportStatus,
		UnixTime:         UnixTimeInMillis(),
		EventType:        UserDeleteHimselfEvent,
	}
}

//it's not analytics event
type UserOnlineEvent struct {
	UserId    string `json:"userId"`
	UnixTime  int64  `json:"unixTime"`
	EventType string `json:"eventType"`
}

func (event UserOnlineEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserOnlineEvent(userId string) *UserOnlineEvent {
	return &UserOnlineEvent{
		UserId:    userId,
		UnixTime:  UnixTimeInMillis(),
		EventType: "AUTH_USER_ONLINE",
	}
}

//image service

type UserAskUploadPhotoLinkEvent struct {
	UserId    string `json:"userId"`
	SourceIp  string `json:"sourceIp"`
	Bucket    string `json:"bucket"`
	PhotoKey  string `json:"photoKey"`
	UnixTime  int64  `json:"unixTime"`
	EventType string `json:"eventType"`
}

func (event UserAskUploadPhotoLinkEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserAskUploadLinkEvent(bucket, photoKey, userId, sourceIp string) *UserAskUploadPhotoLinkEvent {
	return &UserAskUploadPhotoLinkEvent{
		UserId:    userId,
		SourceIp:  sourceIp,
		Bucket:    bucket,
		PhotoKey:  photoKey,
		UnixTime:  UnixTimeInMillis(),
		EventType: "IMAGE_USER_ASK_UPLOAD_PHOTO_LINK",
	}
}

type UserUploadedPhotoEvent struct {
	UserId    string `json:"userId"`
	Bucket    string `json:"bucket"`
	PhotoKey  string `json:"photoKey"`
	PhotoId   string `json:"photoId"`
	PhotoType string `json:"photoType"`
	Size      int64  `json:"size"`
	UnixTime  int64  `json:"unixTime"`
	EventType string `json:"eventType"`
}

func (event UserUploadedPhotoEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserUploadedPhotoEvent(userId, bucket, key, photoId, photoType string, size int64) *UserUploadedPhotoEvent {
	return &UserUploadedPhotoEvent{
		UserId:    userId,
		Bucket:    bucket,
		PhotoKey:  key,
		PhotoId:   photoId,
		PhotoType: photoType,
		Size:      size,
		UnixTime:  UnixTimeInMillis(),
		EventType: "IMAGE_USER_UPLOAD_PHOTO",
	}
}

type UserDeletePhotoEvent struct {
	UserId               string `json:"userId"`
	SourceIp             string `json:"sourceIp"`
	PhotoId              string `json:"photoId"`
	UserTakePartInReport bool   `json:"userTakePartInReport"`
	UnixTime             int64  `json:"unixTime"`
	EventType            string `json:"eventType"`
}

func (event UserDeletePhotoEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserDeletePhotoEvent(userId, photoId, sourceIp string, userTakePartInReport bool) *UserDeletePhotoEvent {
	return &UserDeletePhotoEvent{
		UserId:               userId,
		SourceIp:             sourceIp,
		PhotoId:              photoId,
		UserTakePartInReport: userTakePartInReport,
		UnixTime:             UnixTimeInMillis(),
		EventType:            "IMAGE_USER_DELETE_PHOTO",
	}
}

type RemoveTooLargeObjectEvent struct {
	UserId    string `json:"userId"`
	Bucket    string `json:"bucket"`
	Key       string `json:"key"`
	Size      int64  `json:"size"`
	UnixTime  int64  `json:"unixTime"`
	EventType string `json:"eventType"`
}

func (event RemoveTooLargeObjectEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewRemoveTooLargeObjectEvent(userId, bucket, key string, size int64) *RemoveTooLargeObjectEvent {
	return &RemoveTooLargeObjectEvent{
		UserId:    userId,
		Bucket:    bucket,
		Key:       key,
		Size:      size,
		UnixTime:  UnixTimeInMillis(),
		EventType: "IMAGE_REMOVE_TO_BIG_S3_OBJECT",
	}
}

type GetOwnPhotosEvent struct {
	UserId      string `json:"userId"`
	SourceIp    string `json:"sourceIp"`
	OwnPhotoNum int    `json:"ownPhotoNum"`
	UnixTime    int64  `json:"unixTime"`
	EventType   string `json:"eventType"`
}

func (event GetOwnPhotosEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewGetOwnPhotosEvent(userId, sourceIp string, ownPhotoNum int) *GetOwnPhotosEvent {
	return &GetOwnPhotosEvent{
		UserId:      userId,
		SourceIp:    sourceIp,
		OwnPhotoNum: ownPhotoNum,
		UnixTime:    UnixTimeInMillis(),
		EventType:   "IMAGE_GET_OWN_PHOTOS",
	}
}

type PhotoResizeEvent struct {
	UserId            string `json:"userId"`
	PhotoId           string `json:"photoId"`
	ResizedPhotoId    string `json:"resizedPhotoId"`
	ResizedResolution string `json:"resizedResolution"`
	ResizedPhotoLink  string `json:"resizedPhotoLink"`
	UnixTime          int64  `json:"unixTime"`
	EventType         string `json:"eventType"`
}

func (event PhotoResizeEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewPhotoResizeEvent(userId, photoId, resizedPhotoId, resizedResolution, resizedPhotoLink string) *PhotoResizeEvent {
	return &PhotoResizeEvent{
		UserId:            userId,
		PhotoId:           photoId,
		ResizedPhotoId:    resizedPhotoId,
		ResizedResolution: resizedResolution,
		ResizedPhotoLink:  resizedPhotoLink,
		UnixTime:          UnixTimeInMillis(),
		EventType:         ResizePhotoInternalEvent,
	}
}

//Internal events in kinesis stream
const (
	LikePhotoInternalEvent = "INTERNAL_PHOTO_LIKE_EVENT"
	UserDeleteHimselfEvent = "AUTH_USER_CALL_DELETE_HIMSELF"
	UserBlockEvent         = "ACTION_USER_BLOCK_OTHER"
	UserMessageEvent       = "INTERNAL_USER_SEND_MESSAGE_EVENT"

	DeleteUserConversationInternalEvent = "INTERNAL_DELETE_USER_CONVERSATION_EVENT"
	HideUserPhotoInternalEvent          = "INTERNAL_HIDE_PHOTO_EVENT"
	ResizePhotoInternalEvent            = "INTERNAL_RESIZE_PHOTO_EVENT"
)

type BaseInternalEvent struct {
	EventType string `json:"eventType"`
}

func (event BaseInternalEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

type PhotoLikeInternalEvent struct {
	EventType       string `json:"eventType"`
	UserId          string `json:"userId"`
	OriginalPhotoId string `json:"originPhotoId"`
}

func (event PhotoLikeInternalEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

type HidePhotoInternalEvent struct {
	EventType       string `json:"eventType"`
	UserId          string `json:"userId"`
	OriginalPhotoId string `json:"originPhotoId"`
}

func (event HidePhotoInternalEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

//actions service

type UserLikePhotoEvent struct {
	UserId                string `json:"userId"`
	SourceIp              string `json:"sourceIp"`
	PhotoId               string `json:"photoId"`
	OriginPhotoId         string `json:"originPhotoId"`
	TargetUserId          string `json:"targetUserId"`
	LikeCount             int    `json:"likeCount"`
	Source                string `json:"source"`
	LikedAt               int64  `json:"likedAt"`
	UnixTime              int64  `json:"unixTime"`
	EventType             string `json:"eventType"`
	InternalServiceSource string `json:"internalServiceSource"`
}

func (event UserLikePhotoEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserLikePhotoEvent(userId, photoId, originPhotoId, targetUserId, source, sourceIp string, likeCount int, likedAt int64, serviceName string) *UserLikePhotoEvent {
	return &UserLikePhotoEvent{
		UserId:                userId,
		SourceIp:              sourceIp,
		PhotoId:               photoId,
		OriginPhotoId:         originPhotoId,
		TargetUserId:          targetUserId,
		LikeCount:             likeCount,
		LikedAt:               likedAt,
		Source:                source,
		UnixTime:              UnixTimeInMillis(),
		EventType:             "ACTION_USER_LIKE_PHOTO",
		InternalServiceSource: serviceName,
	}
}

type UserViewPhotoEvent struct {
	UserId                string `json:"userId"`
	SourceIp              string `json:"sourceIp"`
	PhotoId               string `json:"photoId"`
	OriginPhotoId         string `json:"originPhotoId"`
	TargetUserId          string `json:"targetUserId"`
	ViewCount             int    `json:"viewCount"`
	ViewTimeMillis        int64  `json:"viewTimeMillis"`
	ViewAt                int64  `json:"viewAt"`
	Source                string `json:"source"`
	UnixTime              int64  `json:"unixTime"`
	EventType             string `json:"eventType"`
	InternalServiceSource string `json:"internalServiceSource"`
}

func (event UserViewPhotoEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserViewPhotoEvent(userId, photoId, originPhotoId, targetUserId, source, sourceIp string, viewCount int, viewTimeMillis int64, viewAt int64, serviceName string) *UserViewPhotoEvent {
	return &UserViewPhotoEvent{
		UserId:                userId,
		SourceIp:              sourceIp,
		PhotoId:               photoId,
		OriginPhotoId:         originPhotoId,
		TargetUserId:          targetUserId,
		ViewCount:             viewCount,
		ViewTimeMillis:        viewTimeMillis,
		ViewAt:                viewAt,
		Source:                source,
		UnixTime:              UnixTimeInMillis(),
		EventType:             "ACTION_USER_VIEW_PHOTO",
		InternalServiceSource: serviceName,
	}
}

type UserBlockOtherEvent struct {
	UserId                string `json:"userId"`
	SourceIp              string `json:"sourceIp"`
	TargetUserId          string `json:"targetUserId"`
	TargetPhotoId         string `json:"targetPhotoId"`
	OriginPhotoId         string `json:"originPhotoId"`
	BlockReasonNum        int    `json:"blockReasonNum"`
	BlockedAt             int64  `json:"blockedAt"`
	Source                string `json:"source"`
	UnixTime              int64  `json:"unixTime"`
	EventType             string `json:"eventType"`
	InternalServiceSource string `json:"internalServiceSource"`
}

func (event UserBlockOtherEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserBlockOtherEvent(userId, targetUserId, targetPhotoId, originPhotoId, source, sourceIp string, blockedAt int64, blockReasonNum int, serviceName string) *UserBlockOtherEvent {
	return &UserBlockOtherEvent{
		UserId:                userId,
		SourceIp:              sourceIp,
		TargetUserId:          targetUserId,
		TargetPhotoId:         targetPhotoId,
		OriginPhotoId:         originPhotoId,
		BlockedAt:             blockedAt,
		BlockReasonNum:        blockReasonNum,
		Source:                source,
		UnixTime:              UnixTimeInMillis(),
		EventType:             UserBlockEvent,
		InternalServiceSource: serviceName,
	}
}

type UserUnLikePhotoEvent struct {
	UserId                string `json:"userId"`
	SourceIp              string `json:"sourceIp"`
	PhotoId               string `json:"photoId"`
	OriginPhotoId         string `json:"originPhotoId"`
	TargetUserId          string `json:"targetUserId"`
	Source                string `json:"source"`
	UnLikedAt             int64  `json:"unLikedAt"`
	UnixTime              int64  `json:"unixTime"`
	EventType             string `json:"eventType"`
	InternalServiceSource string `json:"internalServiceSource"`
}

func (event UserUnLikePhotoEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserUnLikePhotoEvent(userId, photoId, originPhotoId, targetUserId, source, sourceIp string, unLikedAt int64, serviceName string) *UserUnLikePhotoEvent {
	return &UserUnLikePhotoEvent{
		UserId:                userId,
		SourceIp:              sourceIp,
		PhotoId:               photoId,
		OriginPhotoId:         originPhotoId,
		TargetUserId:          targetUserId,
		UnLikedAt:             unLikedAt,
		Source:                source,
		UnixTime:              UnixTimeInMillis(),
		EventType:             "ACTION_USER_UNLIKE_PHOTO",
		InternalServiceSource: serviceName,
	}
}

type UserMsgEvent struct {
	MessageId      string `json:"messageId"`
	ConversationId string `json:"conversationId"`
	UserId         string `json:"userId"`
	SourceIp       string `json:"sourceIp"`
	TargetUserId   string `json:"targetUserId"`
	PhotoId        string `json:"photoId"`
	OriginPhotoId  string `json:"originPhotoId"`
	Text           string `json:"text"`
	Source         string `json:"source"`
	MessageAt      int64  `json:"messageAt"`
	UnixTime       int64  `json:"unixTime"`
	EventType      string `json:"eventType"`
}

func (event UserMsgEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserMsgEvent(messageId, conversationId, userId, photoId, originPhotoId, targetUserId, source, sourceIp, text string, messageAt int64) *UserMsgEvent {
	return &UserMsgEvent{
		MessageId:      messageId,
		ConversationId: conversationId,
		UserId:         userId,
		SourceIp:       sourceIp,
		TargetUserId:   targetUserId,
		PhotoId:        photoId,
		OriginPhotoId:  originPhotoId,
		Text:           text,
		Source:         source,
		MessageAt:      messageAt,
		UnixTime:       UnixTimeInMillis(),
		EventType:      "ACTION_USER_MESSAGE",
	}
}

type UserViewChantEvent struct {
	UserId         string `json:"userId"`
	SourceIp       string `json:"sourceIp"`
	TargetUserId   string `json:"targetUserId"`
	PhotoId        string `json:"photoId"`
	OriginPhotoId  string `json:"originPhotoId"`
	Source         string `json:"source"`
	ViewTimeMillis int64  `json:"viewTimeMillis"`
	ViewAt         int64  `json:"viewAt"`
	UnixTime       int64  `json:"unixTime"`
	EventType      string `json:"eventType"`
}

func (event UserViewChantEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserViewChantEvent(userId, photoId, originPhotoId, targetUserId, source, sourceIp string, viewAt, viewTimeMillis int64) *UserViewChantEvent {
	return &UserViewChantEvent{
		UserId:         userId,
		SourceIp:       sourceIp,
		TargetUserId:   targetUserId,
		PhotoId:        photoId,
		OriginPhotoId:  originPhotoId,
		Source:         source,
		ViewTimeMillis: viewTimeMillis,
		ViewAt:         viewAt,
		UnixTime:       UnixTimeInMillis(),
		EventType:      "ACTION_USER_VIEW_CHAT",
	}
}

//feeds service

type ProfileWasReturnToNewFacesEvent struct {
	UserId             string   `json:"userId"`
	SourceIp           string   `json:"sourceIp"`
	TargetUserIds      []string `json:"targetUserIds"`
	NewFaceProfilesNum int      `json:"newFaceProfilesNum"`
	RepeatRequestAfter int64    `json:"repeatRequestAfter"`
	UnixTime           int64    `json:"unixTime"`
	EventType          string   `json:"eventType"`
}

func (event ProfileWasReturnToNewFacesEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewProfileWasReturnToNewFacesEvent(userId, sourceIp string, targetIds []string, repeatRequestAfter int64) ProfileWasReturnToNewFacesEvent {
	return ProfileWasReturnToNewFacesEvent{
		UserId:             userId,
		SourceIp:           sourceIp,
		TargetUserIds:      targetIds,
		NewFaceProfilesNum: len(targetIds),
		RepeatRequestAfter: repeatRequestAfter,
		UnixTime:           UnixTimeInMillis(),
		EventType:          "FEEDS_NEW_FACES_SEEN_PROFILES",
	}
}

type ProfileWasReturnToLMMEvent struct {
	UserId             string `json:"userId"`
	SourceIp           string `json:"sourceIp"`
	SourceFeed         string `json:"sourceFeed"`
	LikeYouProfilesNum int    `json:"likeYouProfilesNum"`
	MatchProfilesNum   int    `json:"matchProfilesNum"`
	MessageProfilesNum int    `json:"messageProfilesNum"`
	RepeatRequestAfter int64  `json:"repeatRequestAfter"`
	UnixTime           int64  `json:"unixTime"`
	EventType          string `json:"eventType"`
}

func (event ProfileWasReturnToLMMEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewProfileWasReturnToLMMEvent(userId, sourceIp, sourceFeed string, likesNum, matchNum, messageNum int, repeatRequestAfter int64) ProfileWasReturnToLMMEvent {
	return ProfileWasReturnToLMMEvent{
		UserId:             userId,
		SourceIp:           sourceIp,
		SourceFeed:         sourceFeed,
		LikeYouProfilesNum: likesNum,
		MatchProfilesNum:   matchNum,
		MessageProfilesNum: messageNum,
		RepeatRequestAfter: repeatRequestAfter,
		UnixTime:           UnixTimeInMillis(),
		EventType:          "FEEDS_LLM_PROFILES",
	}
}

type UserSendMessageEvent struct {
	UserId       string `json:"userId"`
	TargetUserId string `json:"targetUserId"`
	Text         string `json:"text"`
	UnixTime     int64  `json:"unixTime"`
	MessageAt    int64  `json:"messageAt"`
	EventType    string `json:"eventType"`
}

func (event UserSendMessageEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserSendMessageEvent(userId, targetUserId, text string, messageAt int64) *UserSendMessageEvent {
	return &UserSendMessageEvent{
		UserId:       userId,
		TargetUserId: targetUserId,
		Text:         text,
		UnixTime:     UnixTimeInMillis(),
		MessageAt:    messageAt,
		EventType:    UserMessageEvent,
	}
}

type DeleteUserConversationEvent struct {
	UserId       string `json:"userId"`
	TargetUserId string `json:"targetUserId"`
	EventType    string `json:"eventType"`
}

func (event DeleteUserConversationEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewDeleteUserConversationEvent(userId, targetUserId string) *DeleteUserConversationEvent {
	return &DeleteUserConversationEvent{
		UserId:       userId,
		TargetUserId: targetUserId,
		EventType:    DeleteUserConversationInternalEvent,
	}
}

type DeviceTokenRegisteredEvent struct {
	UserId      string `json:"userId"`
	DeviceToken string `json:"deviceToken"`
	IsItAndroid bool   `json:"isItAndroid"`
	SourceIp    string `json:"sourceIp"`
	UnixTime    int64  `json:"unixTime"`
	EventType   string `json:"eventType"`
}

func (event DeviceTokenRegisteredEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewDeviceTokenRegisteredEvent(userId, deviceId, sourceIp string, isItAndroid bool) *DeviceTokenRegisteredEvent {
	return &DeviceTokenRegisteredEvent{
		UserId:      userId,
		DeviceToken: deviceId,
		IsItAndroid: isItAndroid,
		SourceIp:    sourceIp,
		UnixTime:    UnixTimeInMillis(),
		EventType:   "PUSH_REGISTER_DEVICE_TOKEN",
	}
}
