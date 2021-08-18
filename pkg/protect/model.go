package protect

// Incomplete Model, generated from my nvr, maybe miss some objects

type Bootstrap struct {
	AuthUserID   string   `json:"authUserId"`
	AccessKey    string   `json:"accessKey"`
	Cameras      []Camera `json:"cameras"`
	Nvr          Nvr      `json:"nvr"`
	LastUpdateID string   `json:"lastUpdateId"`
}

func (b Bootstrap) GetCameraID(id string) *Camera {
	for _, camera := range b.Cameras {
		if camera.ID == id {
			return &camera
		}
	}
	return nil
}

type Motion struct {
	Today       int   `json:"today"`
	Average     int   `json:"average"`
	LastDays    []int `json:"lastDays"`
	RecentHours []int `json:"recentHours"`
}
type Smart struct {
	Today    int   `json:"today"`
	Average  int   `json:"average"`
	LastDays []int `json:"lastDays"`
}
type EventStats struct {
	Motion Motion `json:"motion"`
	Smart  Smart  `json:"smart"`
}
type WiredConnectionState struct {
	PhyRate interface{} `json:"phyRate"`
}
type Channels struct {
	ID                       int         `json:"id"`
	VideoID                  string      `json:"videoId"`
	Name                     string      `json:"name"`
	Enabled                  bool        `json:"enabled"`
	IsRtspEnabled            bool        `json:"isRtspEnabled"`
	RtspAlias                interface{} `json:"rtspAlias"`
	Width                    int         `json:"width"`
	Height                   int         `json:"height"`
	Fps                      int         `json:"fps"`
	Bitrate                  int         `json:"bitrate"`
	MinBitrate               int         `json:"minBitrate"`
	MaxBitrate               int         `json:"maxBitrate"`
	MinClientAdaptiveBitRate int         `json:"minClientAdaptiveBitRate"`
	MinMotionAdaptiveBitRate int         `json:"minMotionAdaptiveBitRate"`
	FpsValues                []int       `json:"fpsValues"`
	IdrInterval              int         `json:"idrInterval"`
}
type IspSettings struct {
	AeMode                         string `json:"aeMode"`
	IrLedMode                      string `json:"irLedMode"`
	IrLedLevel                     int    `json:"irLedLevel"`
	Wdr                            int    `json:"wdr"`
	IcrSensitivity                 int    `json:"icrSensitivity"`
	Brightness                     int    `json:"brightness"`
	Contrast                       int    `json:"contrast"`
	Hue                            int    `json:"hue"`
	Saturation                     int    `json:"saturation"`
	Sharpness                      int    `json:"sharpness"`
	Denoise                        int    `json:"denoise"`
	IsFlippedVertical              bool   `json:"isFlippedVertical"`
	IsFlippedHorizontal            bool   `json:"isFlippedHorizontal"`
	IsAutoRotateEnabled            bool   `json:"isAutoRotateEnabled"`
	IsLdcEnabled                   bool   `json:"isLdcEnabled"`
	Is3DnrEnabled                  bool   `json:"is3dnrEnabled"`
	IsExternalIrEnabled            bool   `json:"isExternalIrEnabled"`
	IsAggressiveAntiFlickerEnabled bool   `json:"isAggressiveAntiFlickerEnabled"`
	IsPauseMotionEnabled           bool   `json:"isPauseMotionEnabled"`
	DZoomCenterX                   int    `json:"dZoomCenterX"`
	DZoomCenterY                   int    `json:"dZoomCenterY"`
	DZoomScale                     int    `json:"dZoomScale"`
	DZoomStreamID                  int    `json:"dZoomStreamId"`
	FocusMode                      string `json:"focusMode"`
	FocusPosition                  int    `json:"focusPosition"`
	TouchFocusX                    int    `json:"touchFocusX"`
	TouchFocusY                    int    `json:"touchFocusY"`
	ZoomPosition                   int    `json:"zoomPosition"`
	MountPosition                  string `json:"mountPosition"`
}
type TalkbackSettings struct {
	TypeFmt       string `json:"typeFmt"`
	TypeIn        string `json:"typeIn"`
	BindAddr      string `json:"bindAddr"`
	BindPort      int    `json:"bindPort"`
	FilterAddr    string `json:"filterAddr"`
	FilterPort    int    `json:"filterPort"`
	Channels      int    `json:"channels"`
	SamplingRate  int    `json:"samplingRate"`
	BitsPerSample int    `json:"bitsPerSample"`
	Quality       int    `json:"quality"`
}
type OsdSettings struct {
	IsNameEnabled  bool `json:"isNameEnabled"`
	IsDateEnabled  bool `json:"isDateEnabled"`
	IsLogoEnabled  bool `json:"isLogoEnabled"`
	IsDebugEnabled bool `json:"isDebugEnabled"`
}
type LedSettings struct {
	IsEnabled bool `json:"isEnabled"`
	BlinkRate int  `json:"blinkRate"`
}
type SpeakerSettings struct {
	IsEnabled              bool `json:"isEnabled"`
	AreSystemSoundsEnabled bool `json:"areSystemSoundsEnabled"`
	Volume                 int  `json:"volume"`
}
type RecordingSettings struct {
	PrePaddingSecs            int    `json:"prePaddingSecs"`
	PostPaddingSecs           int    `json:"postPaddingSecs"`
	MinMotionEventTrigger     int    `json:"minMotionEventTrigger"`
	EndMotionEventDelay       int    `json:"endMotionEventDelay"`
	SuppressIlluminationSurge bool   `json:"suppressIlluminationSurge"`
	Mode                      string `json:"mode"`
	Geofencing                string `json:"geofencing"`
	MotionAlgorithm           string `json:"motionAlgorithm"`
	EnablePirTimelapse        bool   `json:"enablePirTimelapse"`
	UseNewMotionAlgorithm     bool   `json:"useNewMotionAlgorithm"`
}

type Wifi struct {
	Channel        int         `json:"channel"`
	Frequency      int         `json:"frequency"`
	LinkSpeedMbps  interface{} `json:"linkSpeedMbps"`
	SignalQuality  int         `json:"signalQuality"`
	SignalStrength int         `json:"signalStrength"`
}
type Battery struct {
	Percentage interface{} `json:"percentage"`
	IsCharging bool        `json:"isCharging"`
	SleepState string      `json:"sleepState"`
}
type Video struct {
	RecordingStart   int64 `json:"recordingStart"`
	RecordingEnd     int64 `json:"recordingEnd"`
	RecordingStartLQ int64 `json:"recordingStartLQ"`
	RecordingEndLQ   int64 `json:"recordingEndLQ"`
	TimelapseStart   int64 `json:"timelapseStart"`
	TimelapseEnd     int64 `json:"timelapseEnd"`
	TimelapseStartLQ int64 `json:"timelapseStartLQ"`
	TimelapseEndLQ   int64 `json:"timelapseEndLQ"`
}
type Storage1 struct {
	Used int64   `json:"used"`
	Rate float64 `json:"rate"`
}
type Stats struct {
	RxBytes      int      `json:"rxBytes"`
	TxBytes      int64    `json:"txBytes"`
	Wifi         Wifi     `json:"wifi"`
	Battery      Battery  `json:"battery"`
	Video        Video    `json:"video"`
	Storage      Storage1 `json:"storage"`
	WifiQuality  int      `json:"wifiQuality"`
	WifiStrength int      `json:"wifiStrength"`
}
type PrivacyMaskCapability struct {
	MaxMasks      int  `json:"maxMasks"`
	RectangleOnly bool `json:"rectangleOnly"`
}
type Steps struct {
	Max  interface{} `json:"max"`
	Min  interface{} `json:"min"`
	Step interface{} `json:"step"`
}
type Degrees struct {
	Max  interface{} `json:"max"`
	Min  interface{} `json:"min"`
	Step interface{} `json:"step"`
}
type Focus struct {
	Steps   Steps   `json:"steps"`
	Degrees Degrees `json:"degrees"`
}
type Pan struct {
	Steps   Steps   `json:"steps"`
	Degrees Degrees `json:"degrees"`
}
type Tilt struct {
	Steps   Steps   `json:"steps"`
	Degrees Degrees `json:"degrees"`
}
type Zoom struct {
	Steps   Steps   `json:"steps"`
	Degrees Degrees `json:"degrees"`
}
type FeatureFlags struct {
	CanAdjustIrLedLevel     bool                  `json:"canAdjustIrLedLevel"`
	CanMagicZoom            bool                  `json:"canMagicZoom"`
	CanOpticalZoom          bool                  `json:"canOpticalZoom"`
	CanTouchFocus           bool                  `json:"canTouchFocus"`
	HasAccelerometer        bool                  `json:"hasAccelerometer"`
	HasAec                  bool                  `json:"hasAec"`
	HasBattery              bool                  `json:"hasBattery"`
	HasBluetooth            bool                  `json:"hasBluetooth"`
	HasChime                bool                  `json:"hasChime"`
	HasExternalIr           bool                  `json:"hasExternalIr"`
	HasIcrSensitivity       bool                  `json:"hasIcrSensitivity"`
	HasLdc                  bool                  `json:"hasLdc"`
	HasLedIr                bool                  `json:"hasLedIr"`
	HasLedStatus            bool                  `json:"hasLedStatus"`
	HasLineIn               bool                  `json:"hasLineIn"`
	HasMic                  bool                  `json:"hasMic"`
	HasPrivacyMask          bool                  `json:"hasPrivacyMask"`
	HasRtc                  bool                  `json:"hasRtc"`
	HasSdCard               bool                  `json:"hasSdCard"`
	HasSpeaker              bool                  `json:"hasSpeaker"`
	HasWifi                 bool                  `json:"hasWifi"`
	HasHdr                  bool                  `json:"hasHdr"`
	HasAutoICROnly          bool                  `json:"hasAutoICROnly"`
	VideoModes              []string              `json:"videoModes"`
	VideoModeMaxFps         []interface{}         `json:"videoModeMaxFps"`
	HasMotionZones          bool                  `json:"hasMotionZones"`
	HasLcdScreen            bool                  `json:"hasLcdScreen"`
	MountPositions          []interface{}         `json:"mountPositions"`
	SmartDetectTypes        []interface{}         `json:"smartDetectTypes"`
	MotionAlgorithms        []string              `json:"motionAlgorithms"`
	HasSquareEventThumbnail bool                  `json:"hasSquareEventThumbnail"`
	PrivacyMaskCapability   PrivacyMaskCapability `json:"privacyMaskCapability"`
	Focus                   Focus                 `json:"focus"`
	Pan                     Pan                   `json:"pan"`
	Tilt                    Tilt                  `json:"tilt"`
	Zoom                    Zoom                  `json:"zoom"`
	HasSmartDetect          bool                  `json:"hasSmartDetect"`
}
type PirSettings struct {
	PirSensitivity            int `json:"pirSensitivity"`
	PirMotionClipLength       int `json:"pirMotionClipLength"`
	TimelapseFrameInterval    int `json:"timelapseFrameInterval"`
	TimelapseTransferInterval int `json:"timelapseTransferInterval"`
}
type WifiConnectionState struct {
	Channel        int `json:"channel"`
	Frequency      int `json:"frequency"`
	PhyRate        int `json:"phyRate"`
	SignalQuality  int `json:"signalQuality"`
	SignalStrength int `json:"signalStrength"`
}
type LcdMessage struct {
	Type    string `json:"type"`
	Text    string `json:"text"`
	ResetAt int    `json:"resetAt"`
}
type Camera struct {
	IsDeleting                bool                 `json:"isDeleting"`
	Mac                       string               `json:"mac"`
	Host                      string               `json:"host"`
	ConnectionHost            string               `json:"connectionHost"`
	Type                      string               `json:"type"`
	Name                      string               `json:"name"`
	UpSince                   int64                `json:"upSince"`
	Uptime                    int                  `json:"uptime"`
	LastSeen                  int64                `json:"lastSeen"`
	ConnectedSince            int64                `json:"connectedSince"`
	State                     string               `json:"state"`
	HardwareRevision          string               `json:"hardwareRevision"`
	FirmwareVersion           string               `json:"firmwareVersion"`
	LatestFirmwareVersion     string               `json:"latestFirmwareVersion"`
	FirmwareBuild             string               `json:"firmwareBuild"`
	IsUpdating                bool                 `json:"isUpdating"`
	IsAdopting                bool                 `json:"isAdopting"`
	IsAdopted                 bool                 `json:"isAdopted"`
	IsAdoptedByOther          bool                 `json:"isAdoptedByOther"`
	IsProvisioned             bool                 `json:"isProvisioned"`
	IsRebooting               bool                 `json:"isRebooting"`
	IsSSHEnabled              bool                 `json:"isSshEnabled"`
	CanAdopt                  bool                 `json:"canAdopt"`
	IsAttemptingToConnect     bool                 `json:"isAttemptingToConnect"`
	LastMotion                int64                `json:"lastMotion"`
	MicVolume                 int                  `json:"micVolume"`
	IsMicEnabled              bool                 `json:"isMicEnabled"`
	IsRecording               bool                 `json:"isRecording"`
	IsMotionDetected          bool                 `json:"isMotionDetected"`
	PhyRate                   int                  `json:"phyRate"`
	HdrMode                   bool                 `json:"hdrMode"`
	VideoMode                 string               `json:"videoMode"`
	IsProbingForWifi          bool                 `json:"isProbingForWifi"`
	ApMac                     string               `json:"apMac"`
	ApRssi                    int                  `json:"apRssi"`
	ElementInfo               string               `json:"elementInfo"`
	ChimeDuration             int                  `json:"chimeDuration"`
	IsDark                    bool                 `json:"isDark"`
	LastPrivacyZonePositionID interface{}          `json:"lastPrivacyZonePositionId"`
	LastRing                  interface{}          `json:"lastRing"`
	IsLiveHeatmapEnabled      bool                 `json:"isLiveHeatmapEnabled"`
	AnonymousDeviceID         string               `json:"anonymousDeviceId"`
	EventStats                EventStats           `json:"eventStats"`
	WiredConnectionState      WiredConnectionState `json:"wiredConnectionState"`
	Channels                  []Channels           `json:"channels"`
	IspSettings               IspSettings          `json:"ispSettings"`
	TalkbackSettings          TalkbackSettings     `json:"talkbackSettings"`
	OsdSettings               OsdSettings          `json:"osdSettings"`
	LedSettings               LedSettings          `json:"ledSettings"`
	SpeakerSettings           SpeakerSettings      `json:"speakerSettings"`
	RecordingSettings         RecordingSettings    `json:"recordingSettings"`
	Stats                     Stats                `json:"stats"`
	FeatureFlags              FeatureFlags         `json:"featureFlags"`
	PirSettings               PirSettings          `json:"pirSettings"`
	WifiConnectionState       WifiConnectionState  `json:"wifiConnectionState"`
	ID                        string               `json:"id"`
	IsConnected               bool                 `json:"isConnected"`
	Platform                  string               `json:"platform"`
	HasSpeaker                bool                 `json:"hasSpeaker"`
	HasWifi                   bool                 `json:"hasWifi"`
	AudioBitrate              int                  `json:"audioBitrate"`
	CanManage                 bool                 `json:"canManage"`
	IsManaged                 bool                 `json:"isManaged"`
	ModelKey                  string               `json:"modelKey"`
	LcdMessage                LcdMessage           `json:"lcdMessage,omitempty"`
}
type Flags struct {
}
type Web struct {
	LiveviewIncludeGlobal bool `json:"liveview.includeGlobal"`
}
type Settings struct {
	Flags Flags `json:"flags"`
	Web   Web   `json:"web"`
}
type Schedule struct {
	Items []interface{} `json:"items"`
}
type System struct {
}

type Ports struct {
	Ump             int `json:"ump"`
	HTTP            int `json:"http"`
	HTTPS           int `json:"https"`
	Rtsp            int `json:"rtsp"`
	Rtsps           int `json:"rtsps"`
	Rtmp            int `json:"rtmp"`
	DevicesWss      int `json:"devicesWss"`
	CameraHTTPS     int `json:"cameraHttps"`
	CameraTCP       int `json:"cameraTcp"`
	LiveWs          int `json:"liveWs"`
	LiveWss         int `json:"liveWss"`
	TCPStreams      int `json:"tcpStreams"`
	Playback        int `json:"playback"`
	EmsCLI          int `json:"emsCLI"`
	EmsLiveFLV      int `json:"emsLiveFLV"`
	CameraEvents    int `json:"cameraEvents"`
	TCPBridge       int `json:"tcpBridge"`
	Ucore           int `json:"ucore"`
	DiscoveryClient int `json:"discoveryClient"`
}
type WifiSettings struct {
	UseThirdPartyWifi bool        `json:"useThirdPartyWifi"`
	Ssid              interface{} `json:"ssid"`
	Password          interface{} `json:"password"`
}
type LocationSettings struct {
	IsAway              bool    `json:"isAway"`
	IsGeofencingEnabled bool    `json:"isGeofencingEnabled"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Radius              int     `json:"radius"`
}
type NVRFeatureFlags struct {
	Beta bool `json:"beta"`
	Dev  bool `json:"dev"`
}
type CPU struct {
	AverageLoad float64 `json:"averageLoad"`
	Temperature float64 `json:"temperature"`
}
type Memory struct {
	Available int `json:"available"`
	Free      int `json:"free"`
	Total     int `json:"total"`
}
type Devices struct {
	Model   interface{} `json:"model"`
	Size    interface{} `json:"size"`
	Healthy bool        `json:"healthy"`
}
type Storage struct {
	Available   int64     `json:"available"`
	IsRecycling bool      `json:"isRecycling"`
	Size        int64     `json:"size"`
	Type        string    `json:"type"`
	Used        int64     `json:"used"`
	Devices     []Devices `json:"devices"`
}
type Tmpfs struct {
	Available int    `json:"available"`
	Total     int    `json:"total"`
	Used      int    `json:"used"`
	Path      string `json:"path"`
}
type SystemInfo struct {
	CPU     CPU     `json:"cpu"`
	Memory  Memory  `json:"memory"`
	Storage Storage `json:"storage"`
	Tmpfs   Tmpfs   `json:"tmpfs"`
}
type AllMessages struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type DoorbellSettings struct {
	DefaultMessageText           string        `json:"defaultMessageText"`
	DefaultMessageResetTimeoutMs int           `json:"defaultMessageResetTimeoutMs"`
	CustomMessages               []string      `json:"customMessages"`
	AllMessages                  []AllMessages `json:"allMessages"`
}
type SmartDetectAgreement struct {
	Status       string `json:"status"`
	LastUpdateAt int64  `json:"lastUpdateAt"`
}
type RecordingSpace struct {
	Total     int64 `json:"total"`
	Used      int64 `json:"used"`
	Available int64 `json:"available"`
}
type StorageStats struct {
	Utilization       float64        `json:"utilization"`
	Capacity          int64          `json:"capacity"`
	RemainingCapacity int64          `json:"remainingCapacity"`
	RecordingSpace    RecordingSpace `json:"recordingSpace"`
}
type MaxCameraCapacity struct {
	FourK int `json:"4K"`
	HD    int `json:"HD"`
}
type Nvr struct {
	Mac                          string               `json:"mac"`
	Host                         string               `json:"host"`
	Name                         string               `json:"name"`
	CanAutoUpdate                bool                 `json:"canAutoUpdate"`
	IsStatsGatheringEnabled      bool                 `json:"isStatsGatheringEnabled"`
	Timezone                     string               `json:"timezone"`
	Version                      string               `json:"version"`
	UcoreVersion                 string               `json:"ucoreVersion"`
	FirmwareVersion              string               `json:"firmwareVersion"`
	UIVersion                    interface{}          `json:"uiVersion"`
	HardwarePlatform             string               `json:"hardwarePlatform"`
	Ports                        Ports                `json:"ports"`
	Uptime                       int                  `json:"uptime"`
	LastSeen                     int64                `json:"lastSeen"`
	IsUpdating                   bool                 `json:"isUpdating"`
	LastUpdateAt                 interface{}          `json:"lastUpdateAt"`
	IsStation                    bool                 `json:"isStation"`
	EnableAutomaticBackups       bool                 `json:"enableAutomaticBackups"`
	EnableStatsReporting         bool                 `json:"enableStatsReporting"`
	IsSSHEnabled                 bool                 `json:"isSshEnabled"`
	ErrorCode                    interface{}          `json:"errorCode"`
	ReleaseChannel               string               `json:"releaseChannel"`
	Hosts                        []string             `json:"hosts"`
	EnableBridgeAutoAdoption     bool                 `json:"enableBridgeAutoAdoption"`
	HardwareID                   string               `json:"hardwareId"`
	HardwareRevision             string               `json:"hardwareRevision"`
	HostType                     int                  `json:"hostType"`
	HostShortname                string               `json:"hostShortname"`
	IsHardware                   bool                 `json:"isHardware"`
	IsWirelessUplinkEnabled      bool                 `json:"isWirelessUplinkEnabled"`
	TimeFormat                   string               `json:"timeFormat"`
	TemperatureUnit              string               `json:"temperatureUnit"`
	RecordingRetentionDurationMs interface{}          `json:"recordingRetentionDurationMs"`
	EnableCrashReporting         bool                 `json:"enableCrashReporting"`
	DisableAudio                 bool                 `json:"disableAudio"`
	AnalyticsData                string               `json:"analyticsData"`
	AnonymousDeviceID            string               `json:"anonymousDeviceId"`
	CameraUtilization            int                  `json:"cameraUtilization"`
	IsRecycling                  bool                 `json:"isRecycling"`
	AvgMotions                   []float64            `json:"avgMotions"`
	WifiSettings                 WifiSettings         `json:"wifiSettings"`
	LocationSettings             LocationSettings     `json:"locationSettings"`
	FeatureFlags                 NVRFeatureFlags      `json:"featureFlags"`
	SystemInfo                   SystemInfo           `json:"systemInfo"`
	DoorbellSettings             DoorbellSettings     `json:"doorbellSettings"`
	SmartDetectAgreement         SmartDetectAgreement `json:"smartDetectAgreement"`
	StorageStats                 StorageStats         `json:"storageStats"`
	ID                           string               `json:"id"`
	IsAway                       bool                 `json:"isAway"`
	IsSetup                      bool                 `json:"isSetup"`
	Network                      string               `json:"network"`
	Type                         string               `json:"type"`
	UpSince                      int64                `json:"upSince"`
	IsRecordingDisabled          bool                 `json:"isRecordingDisabled"`
	IsRecordingMotionOnly        bool                 `json:"isRecordingMotionOnly"`
	MaxCameraCapacity            MaxCameraCapacity    `json:"maxCameraCapacity"`
	ModelKey                     string               `json:"modelKey"`
}

type ProtectNvrUpdatePayloadCameraUpdate struct {
	IsMotionDetected bool       `json:"isMotionDetected"`
	LastMotion       int        `json:"lastMotion"`
	LastRing         int        `json:"lastRing"`
	LcdMessage       LcdMessage `json:"lcdMessage"`
}

type URL struct {
	URL string `json:"url"`
}
