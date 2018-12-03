package commons

import (
	"time"
	"fmt"
)

type UserAcceptTermsEvent struct {
	UserId                     string `json:"userId"`
	Locale                     string `json:"locale"`
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

func NewUserAcceptTermsEvent(userId, customerId, locale, sourceIp, deviceModel, osVersion string,
	dateTimeLegalAge, dateTimePrivacyNotes, dateTimeTermsAndConditions int64, isItAndroid bool) *UserAcceptTermsEvent {
	if isItAndroid {
		return &UserAcceptTermsEvent{
			UserId: userId,
			Locale: locale,
			//gdpr?
			SourceIp:   sourceIp,
			CustomerId: customerId,

			UnixTime:                   time.Now().Unix(),
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
		Locale: locale,
		//gdpr?
		SourceIp:   sourceIp,
		CustomerId: customerId,

		UnixTime:                   time.Now().Unix(),
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
	Sex         string `json:"sex"`
	YearOfBirth int    `json:"yearOfBirth"`
	UnixTime    int64  `json:"unixTime"`
	EventType   string `json:"eventType"`
}

func (event UserProfileCreatedEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserProfileCreatedEvent(userId, sex string, yearOfBirth int) *UserProfileCreatedEvent {
	return &UserProfileCreatedEvent{
		UserId:      userId,
		Sex:         sex,
		YearOfBirth: yearOfBirth,
		UnixTime:    time.Now().Unix(),
		EventType:   "AUTH_USER_PROFILE_CREATED",
	}
}

type UserSettingsUpdatedEvent struct {
	UserId              string `json:"userId"`
	SafeDistanceInMeter int    `json:"safeDistanceInMeter"` // 0 (default for men) || 10 (default for women)
	PushMessages        bool   `json:"pushMessages"`        // true (default for men) || false (default for women)
	PushMatches         bool   `json:"pushMatches"`         // true (default)
	PushLikes           string `json:"pushLikes"`           //EVERY (default for men) || 10_NEW (default for women) || 100_NEW
	UnixTime            int64  `json:"unixTime"`
	EventType           string `json:"eventType"`
}

func (event UserSettingsUpdatedEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserSettingsUpdatedEvent(userId string, safeDistanceInMeter int, pushMessages, pushMatches bool, pushLikes string, ) *UserSettingsUpdatedEvent {
	return &UserSettingsUpdatedEvent{
		UserId:              userId,
		SafeDistanceInMeter: safeDistanceInMeter,
		PushMessages:        pushMessages,
		PushMatches:         pushMatches,
		PushLikes:           pushLikes,
		UnixTime:            time.Now().Unix(),
		EventType:           "AUTH_USER_SETTINGS_UPDATED",
	}
}

type UserCallDeleteHimselfEvent struct {
	UserId           string `json:"userId"`
	UserReportStatus string `json:"userReportStatus"`
	UnixTime         int64  `json:"unixTime"`
	EventType        string `json:"eventType"`
}

func (event UserCallDeleteHimselfEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserCallDeleteHimselfEvent(userId string, userReportStatus string) *UserCallDeleteHimselfEvent {
	return &UserCallDeleteHimselfEvent{
		UserId:           userId,
		UserReportStatus: userReportStatus,
		UnixTime:         time.Now().Unix(),
		EventType:        "AUTH_USER_CALL_DELETE_HIMSELF",
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
		UnixTime:  time.Now().Unix(),
		EventType: "AUTH_USER_ONLINE",
	}
}

//image service

type UserAskUploadPhotoLinkEvent struct {
	UserId    string `json:"userId"`
	Bucket    string `json:"bucket"`
	PhotoKey  string `json:"photoKey"`
	UnixTime  int64  `json:"unixTime"`
	EventType string `json:"eventType"`
}

func (event UserAskUploadPhotoLinkEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserAskUploadLinkEvent(bucket, photoKey, userId string) *UserAskUploadPhotoLinkEvent {
	return &UserAskUploadPhotoLinkEvent{
		UserId:    userId,
		Bucket:    bucket,
		PhotoKey:  photoKey,
		UnixTime:  time.Now().Unix(),
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
		UnixTime:  time.Now().Unix(),
		EventType: "IMAGE_USER_UPLOAD_PHOTO",
	}
}

type UserDeletePhotoEvent struct {
	UserId    string `json:"userId"`
	PhotoId   string `json:"photoId"`
	UnixTime  int64  `json:"unixTime"`
	EventType string `json:"eventType"`
}

func (event UserDeletePhotoEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserDeletePhotoEvent(userId, photoId string) *UserDeletePhotoEvent {
	return &UserDeletePhotoEvent{
		UserId:    userId,
		PhotoId:   photoId,
		UnixTime:  time.Now().Unix(),
		EventType: "IMAGE_USER_DELETE_PHOTO",
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
		UnixTime:  time.Now().Unix(),
		EventType: "IMAGE_REMOVE_TO_BIG_S3_OBJECT",
	}
}

//Internal events in kinesis stream
const (
	LikePhotoInternalEvent = "INTERNAL_PHOTO_LIKE_EVENT"
	UserDeleteHimselfEvent = "AUTH_USER_CALL_DELETE_HIMSELF"
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

//actions service

type UserLikePhotoEvent struct {
	UserId                string `json:"userId"`
	PhotoId               string `json:"photoId"`
	OriginPhotoId         string `json:"originPhotoId"`
	TargetUserId          string `json:"targetUserId"`
	LikeCount             int    `json:"likeCount"`
	Source                string `json:"source"`
	LikedAt               int    `json:"likedAt"`
	UnixTime              int64  `json:"unixTime"`
	EventType             string `json:"eventType"`
	InternalServiceSource string `json:"internalServiceSource"`
}

func (event UserLikePhotoEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserLikePhotoEvent(userId, photoId, originPhotoId, targetUserId, source string, likeCount, likedAt int, serviceName string) *UserLikePhotoEvent {
	return &UserLikePhotoEvent{
		UserId:                userId,
		PhotoId:               photoId,
		OriginPhotoId:         originPhotoId,
		TargetUserId:          targetUserId,
		LikeCount:             likeCount,
		LikedAt:               likedAt,
		Source:                source,
		UnixTime:              time.Now().Unix(),
		EventType:             "ACTION_USER_LIKE_PHOTO",
		InternalServiceSource: serviceName,
	}
}

type UserViewPhotoEvent struct {
	UserId                string `json:"userId"`
	PhotoId               string `json:"photoId"`
	OriginPhotoId         string `json:"originPhotoId"`
	TargetUserId          string `json:"targetUserId"`
	ViewCount             int    `json:"viewCount"`
	ViewTimeSec           int    `json:"viewTimeSec"`
	ViewAt                int    `json:"viewAt"`
	Source                string `json:"source"`
	UnixTime              int64  `json:"unixTime"`
	EventType             string `json:"eventType"`
	InternalServiceSource string `json:"internalServiceSource"`
}

func (event UserViewPhotoEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserViewPhotoEvent(userId, photoId, originPhotoId, targetUserId, source string, viewCount, viewTimeSec, viewAt int, serviceName string) *UserViewPhotoEvent {
	return &UserViewPhotoEvent{
		UserId:                userId,
		PhotoId:               photoId,
		OriginPhotoId:         originPhotoId,
		TargetUserId:          targetUserId,
		ViewCount:             viewCount,
		ViewTimeSec:           viewTimeSec,
		ViewAt:                viewAt,
		Source:                source,
		UnixTime:              time.Now().Unix(),
		EventType:             "ACTION_USER_VIEW_PHOTO",
		InternalServiceSource: serviceName,
	}
}

type UserBlockOtherEvent struct {
	UserId                string `json:"userId"`
	TargetUserId          string `json:"targetUserId"`
	BlockedAt             int    `json:"blockedAt"`
	Source                string `json:"source"`
	UnixTime              int64  `json:"unixTime"`
	EventType             string `json:"eventType"`
	InternalServiceSource string `json:"internalServiceSource"`
}

func (event UserBlockOtherEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserBlockOtherEvent(userId, targetUserId, source string, blockedAt int, serviceName string) *UserBlockOtherEvent {
	return &UserBlockOtherEvent{
		UserId:                userId,
		TargetUserId:          targetUserId,
		BlockedAt:             blockedAt,
		Source:                source,
		UnixTime:              time.Now().Unix(),
		EventType:             "ACTION_USER_BLOCK_OTHER",
		InternalServiceSource: serviceName,
	}
}

type UserUnLikePhotoEvent struct {
	UserId                string `json:"userId"`
	PhotoId               string `json:"photoId"`
	OriginPhotoId         string `json:"originPhotoId"`
	TargetUserId          string `json:"targetUserId"`
	Source                string `json:"source"`
	UnLikedAt             int    `json:"unLikedAt"`
	UnixTime              int64  `json:"unixTime"`
	EventType             string `json:"eventType"`
	InternalServiceSource string `json:"internalServiceSource"`
}

func (event UserUnLikePhotoEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewUserUnLikePhotoEvent(userId, photoId, originPhotoId, targetUserId, source string, unLikedAt int, serviceName string) *UserUnLikePhotoEvent {
	return &UserUnLikePhotoEvent{
		UserId:                userId,
		PhotoId:               photoId,
		OriginPhotoId:         originPhotoId,
		TargetUserId:          targetUserId,
		UnLikedAt:             unLikedAt,
		Source:                source,
		UnixTime:              time.Now().Unix(),
		EventType:             "ACTION_USER_UNLIKE_PHOTO",
		InternalServiceSource: serviceName,
	}
}

//feeds service

type ProfileWasReturnToNewFacesEvent struct {
	UserId              string   `json:"userId"`
	TargetUserIds       []string `json:"targetUserIds"`
	TimeToDeleteViewRel int64    `json:"timeToDelete"`
	SeenProfilesNum     int      `json:"seenProfilesNum"`
	UnixTime            int64    `json:"unixTime"`
	EventType           string   `json:"eventType"`
}

func (event ProfileWasReturnToNewFacesEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewProfileWasReturnToNewFacesEvent(userId string, timeToDeleteViewRel int64, targetIds []string) ProfileWasReturnToNewFacesEvent {
	return ProfileWasReturnToNewFacesEvent{
		UserId:              userId,
		TargetUserIds:       targetIds,
		TimeToDeleteViewRel: timeToDeleteViewRel,
		SeenProfilesNum:     len(targetIds),
		UnixTime:            time.Now().Unix(),
		EventType:           "FEEDS_NEW_FACES_SEEN_PROFILES",
	}
}

type ProfileWasReturnToLMMEvent struct {
	UserId        string   `json:"userId"`
	TargetUserIds []string `json:"targetUserIds"`
	UnixTime      int64    `json:"unixTime"`
	EventType     string   `json:"eventType"`
}

func (event ProfileWasReturnToLMMEvent) String() string {
	return fmt.Sprintf("%#v", event)
}

func NewProfileWasReturnToLMMEvent(userId string, targetIds []string) ProfileWasReturnToLMMEvent {
	return ProfileWasReturnToLMMEvent{
		UserId:        userId,
		TargetUserIds: targetIds,
		UnixTime:      time.Now().Unix(),
		EventType:     "FEEDS_LLM_PROFILES",
	}
}
