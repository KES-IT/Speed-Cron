package g_consts

var (
	PingUrl               = "https://www.ithome.com/"
	BaseUrl               = "unknown"
	SpeedBackendUrl       = "/UploadSpeedData"
	PingBackendUrl        = "/UploadPingData"
	AuthBackendUrl        = "/DeviceAuth"
	ConfigBackendUrl      = "/GetCronConfig"
	MonitorLogBackendUrl  = "/ReportMonitorData"
	MonitorListBackendUrl = "/GetAllTaskWebsite"
	StableBackendUrl      = "/GetLatestVersion"
	BetaBackendUrl        = "/GetBetaVersion"
	DownloadProxyUrl      = "https://gh.buycoffee.tech/"
	DownloadFileName      = "speed_cron_windows_amd64.exe"
)

// BackendBaseUrl 返回后端地址
func BackendBaseUrl() string {
	return BaseUrl
}
