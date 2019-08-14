package commons

const (
	Region     = "eu-west-1"
	MaxRetries = 3

	TwilioSecretKeyBase = "%s/Twilio/Api/Key"
	TwilioApiKeyName    = "twilio-api-key"

	NexmoSecretKeyBase = "%s/Nexmo/Api/Key"
	NexmoApiKeyName    = "nexmo-api-key"

	NexmoApiSecretKeyBase = "%s/Nexmo/Api/Secret"
	NexmoApiSecretKeyName = "nexmo-api-secret"

	SecretWordKeyName = "secret-word-key"
	SecretWordKeyBase = "%s/SecretWord"

	SessionGSIName = "sessionGSI"

	PhoneColumnName     = "phone"
	UserIdColumnName    = "user_id"
	SessionIdColumnName = "session_id"

	CountryCodeColumnName      = "country_code"
	PhoneNumberColumnName      = "phone_number"
	TokenUpdatedTimeColumnName = "token_updated_at"

	SessionTokenColumnName = "session_token"
	SexColumnName          = "sex"

	ReferralIdColumnName          = "referral_id"
	PrivateKeyColumnName          = "private_key"
	YearOfBirthColumnName         = "year_of_birth"
	ProfileCreatedAt              = "profile_created_at"
	CustomerIdColumnName          = "customer_id"
	VerifyProviderColumnName      = "verify_provider"
	VerifyRequestIdColumnName     = "verify_request_id"
	VerificationStatusColumnName  = "verify_status"
	VerificationStartAtColumnName = "verification_start_at"
	LocaleColumnName              = "locale"
	AndroidDeviceModelColumnName  = "android_device"
	AndroidOsVersionColumnName    = "android_os_version"
	IOSDeviceModelColumnName      = "ios_device"
	IOsVersionColumnName          = "ios_version"
	CurrentActiveDeviceIsAndroid  = "current_device_is_android"
	UserStatusColumnName          = "user_status"
	UserReportStatusColumnName    = "user_report_status"
	UserEmailColumnName           = "email"

	UserProfilePropertyColumnName       = "property"
	UserProfileTransportColumnName      = "transport"
	UserProfileIncomeColumnName         = "income"
	UserProfileHeightColumnName         = "height"
	UserProfileEducationLevelColumnName = "edu_level"
	UserProfileHairColorColumnName      = "hair_color"
	UserProfileChildrenColumnName       = "children"
	UserProfileNameColumnName           = "name"
	UserProfileJobTitleColumnName       = "job_title"
	UserProfileCompanyColumnName        = "company"
	UserProfileEducationTextColumnName  = "education_text"
	UserProfileAboutColumnName          = "about"
	UserProfileInstagramColumnName      = "instagram"
	UserProfileTikTokColumnName         = "tiktok"
	UserProfileWhereILiveColumnName     = "where_i_live"
	UserProfileWhereIFromColumnName     = "where_i_from"

	EmailAuthMailColumnName            = "email"
	EmailAuthStatusColumnName          = "auth_status"
	EmailAuthSessionIdColumnName       = "auth_session_id"
	EmailAuthUserIdColumnName          = "user_id"
	EmailAuthStatusStartedValue        = "AUTH_STARTED"
	EmailAuthStatusAccountCreatedValue = "ACCOUNT_CREATED"

	AuthConfirmMailColumnName      = "email"
	AuthConfirmPinColumnName       = "pin_code"
	AuthConfirmSessionIdColumnName = "auth_session_id"
	AuthConfirmUserIdColumnName    = "user_id"
	AuthConfirmStatusColumnName    = "confirmation_status"
	AuthConfirmStatusStartedValue  = "CONFIRMATION_STARTED"
	AuthConfirmStatusCompleteValue = "CONFIRMATION_COMPLETE"

	UpdatedTimeColumnName    = "updated_at"
	LastOnlineTimeColumnName = "last_online_time"
	CurrentAndroidBuildNum   = "current_android_buildnum"
	CurrentiOSBuildNum       = "current_ios_buildnum"

	PushColumnName                = "push"
	PushNewLikeColumnName         = "push_new_like"
	PushNewMatchColumnName        = "push_new_match"
	PushNewMessageColumnName      = "push_new_message"
	TimeZoneColumnName            = "time_zone"
	SafeDistanceInMeterColumnName = "safe_distance_in_meter"
	PushMessagesColumnName        = "push_messages"
	PushMatchesColumnName         = "push_matches"
	PushLikesColumnName           = "push_likes"

	AccessTokenUserIdClaim       = "userId"
	AccessTokenSessionTokenClaim = "sessionToken"

	AndroidBuildNum = "x-ringoid-android-buildnum"
	iOSdBuildNum    = "x-ringoid-ios-buildnum"

	PhotoIdColumnName        = "photo_id"
	PhotoSourceUriColumnName = "photo_uri"
	PhotoTypeColumnName      = "photo_type"
	PhotoBucketColumnName    = "photo_bucket"
	PhotoKeyColumnName       = "photo_key"
	PhotoSizeColumnName      = "photo_size"
	PhotoDeletedAtColumnName = "deleted_at"
	PhotoHiddenAtColumnName  = "hidden_at"
	PhotoLikesColumnName     = "likes"

	PhotoPrimaryKeyMetaPostfix = "_meta"

	MessagesConversationIdColumnName = "conversion_id"
	MessagesCreatedAtColumnName      = "message_created_at"
	MessagesSenderIdColumnName       = "sender_id"
	MessagesTextColumnName           = "message_text"

	DeviceTokenColumnName         = "device_token"
	PlatformEndpointArnColumnName = "platform_arn"
	OSColumnName                  = "os"

	InternalServerError           = `{"errorCode":"InternalServerError","errorMessage":"Internal Server Error"}`
	WrongRequestParamsClientError = `{"errorCode":"WrongRequestParamsClientError","errorMessage":"Wrong request params"}`
	PhoneNumberClientError        = `{"errorCode":"PhoneNumberClientError","errorMessage":"Phone number is invalid"}`
	CountryCallingCodeClientError = `{"errorCode":"CountryCallingCodeClientError","errorMessage":"Country code is invalid"}`

	WrongSessionIdClientError        = `{"errorCode":"WrongSessionIdClientError","errorMessage":"Session id is invalid"}`
	NoPendingVerificationClientError = `{"errorCode":"NoPendingVerificationClientError","errorMessage":"No pending verifications found"}`
	WrongVerificationCodeClientError = `{"errorCode":"WrongVerificationCodeClientError","errorMessage":"Wrong verification code"}`

	WrongYearOfBirthClientError   = `{"errorCode":"WrongYearOfBirthClientError","errorMessage":"Wrong year of birth"}`
	WrongSexClientError           = `{"errorCode":"WrongSexClientError","errorMessage":"Wrong sex"}`
	InvalidAccessTokenClientError = `{"errorCode":"InvalidAccessTokenClientError","errorMessage":"Invalid access token"}`

	TooOldAppVersionClientError = `{"errorCode":"TooOldAppVersionClientError","errorMessage":"Too old app version"}`

	EmailNotVerifiedClientError             = `{"errorCode":"EmailNotVerifiedClientError","errorMessage":"Email not verified"}`
	ErrorCodeEmailNotVerifiedClientError    = "EmailNotVerifiedClientError"
	ErrorMessageEmailNotVerifiedClientError = "Email not verified"

	EmailConcurrentUsageClientError     = `{"errorCode":"EmailConcurrentUsageClientError","errorMessage":"Email concurrent usage error"}`
	EmailInvalidVerificationClientError = `{"errorCode":"EmailInvalidVerificationClientError","errorMessage":"Email invalid verification"}`
	WrongPinCodeClientError             = `{"errorCode":"WrongPinCodeClientError","errorMessage":"Wrong pin code"}`
	EmailAlreadyInUseClientError        = `{"errorCode":"EmailAlreadyInUseClientError","errorMessage":"Email already in use"}`

	Twilio = "Twilio"
	Nexmo  = "Nexmo"

	LikeActionType     = "LIKE"
	ViewActionType     = "VIEW"
	BlockActionType    = "BLOCK"
	UnlikeActionType   = "UNLIKE"
	MessageActionType  = "MESSAGE"
	ViewChatActionType = "VIEW_CHAT"
	LocationActionType = "LOCATION"

	UserActiveStatus = "ACTIVE"
	UserHiddenStatus = "HIDDEN"

	UserCleanReportStatus = "CLEAN"
	UserTakePartInReport  = "TAKE_PART_IN_REPORT"

	FeedNameNewFaces        = "new_faces"
	FeedNameWhoLikedMe      = "who_liked_me"
	FeedNameMatches         = "matches"
	FeedNameMessages        = "messages"
	FeedNameMessagesInbox   = "messages_inbox"
	FeedNameMessagesStarred = "messages_starred"
	FeedNameMessagesSent    = "messages_sent"
	FeedNameChat            = "chat"
	FeedNameHellos          = "hellos"
	FeedNameInbox           = "inbox"
	FeedNameSent            = "sent"

	AndroidOperationalSystemName = "android"
	IOSOperationalSystemName     = "ios"

	PrepareLimitNum        = 120
	NewFacesHardcodedLimit = 50
)

var MinimalAndroidBuildNum = 211
var MinimaliOSBuildNum = 134

var AllowedPhotoResolution map[string]bool
var ResolutionValues map[string]int
var BiggestDefaultPhotoResolution = "1440x1920" //currently use if client ask not existing resolution

const DefaultJPEGQuality = 80
const DefaultMaxPhotoSize = 10000000 //10 Mb
const ThumbnailJPEGQuality = 30
const ThumbnailPhotoType = "thumbnail"

const ThumbnailPhotoWidth = 150
const ThumbnailPhotoHeight = 200

//const ThumbnailPhotoWidth = 480
//const ThumbnailPhotoHeight = 640

var FeedNames map[string]bool
var ActionNames map[string]bool

func init() {
	AllowedPhotoResolution = make(map[string]bool)
	AllowedPhotoResolution["480x640"] = true
	AllowedPhotoResolution["720x960"] = true
	AllowedPhotoResolution["1080x1440"] = true
	AllowedPhotoResolution["1440x1920"] = true

	AllowedPhotoResolution["640x852"] = true
	AllowedPhotoResolution["750x1000"] = true
	AllowedPhotoResolution["828x1104"] = true
	AllowedPhotoResolution["1125x1500"] = true
	AllowedPhotoResolution["1242x1656"] = true

	ResolutionValues = make(map[string]int)
	ResolutionValues["480x640_width"] = 480
	ResolutionValues["480x640_height"] = 640

	ResolutionValues["720x960_width"] = 720
	ResolutionValues["720x960_height"] = 960

	ResolutionValues["1080x1440_width"] = 1080
	ResolutionValues["1080x1440_height"] = 1440

	ResolutionValues["1440x1920_width"] = 1440
	ResolutionValues["1440x1920_height"] = 1920

	ResolutionValues["640x852_width"] = 640
	ResolutionValues["640x852_height"] = 852

	ResolutionValues["750x1000_width"] = 750
	ResolutionValues["750x1000_height"] = 1000

	ResolutionValues["828x1104_width"] = 828
	ResolutionValues["828x1104_height"] = 1104

	ResolutionValues["1125x1500_width"] = 1125
	ResolutionValues["1125x1500_height"] = 1500

	ResolutionValues["1242x1656_width"] = 1242
	ResolutionValues["1242x1656_height"] = 1656

	FeedNames = make(map[string]bool)
	FeedNames[FeedNameNewFaces] = true
	FeedNames[FeedNameWhoLikedMe] = true
	FeedNames[FeedNameMatches] = true
	FeedNames[FeedNameMessages] = true
	FeedNames[FeedNameMessagesInbox] = true
	FeedNames[FeedNameMessagesStarred] = true
	FeedNames[FeedNameMessagesSent] = true
	FeedNames[FeedNameChat] = true
	FeedNames[FeedNameHellos] = true
	FeedNames[FeedNameInbox] = true
	FeedNames[FeedNameSent] = true

	ActionNames = make(map[string]bool)
	ActionNames[LikeActionType] = true
	ActionNames[ViewActionType] = true
	ActionNames[BlockActionType] = true
	ActionNames[UnlikeActionType] = true
	ActionNames[MessageActionType] = true
	ActionNames[ViewChatActionType] = true
	ActionNames[LocationActionType] = true

}
