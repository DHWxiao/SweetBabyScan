package initializes

import (
	"github.com/inbug-team/SweetBabyScan/initializes/initialize_http_client"
	"github.com/inbug-team/SweetBabyScan/initializes/initialize_screenshot"
)

func InitAll() {
	initialize_screenshot.InitScreenShot()
	initialize_http_client.InitHttpClient()
}
