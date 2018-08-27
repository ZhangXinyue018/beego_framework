package controllers

type VersionController struct {
	MainController
}

type Version struct {
	VersionTime    string `json:"version_time"`
	VersionNumber  string `json:"version_number"`
	VersionMessage string `json:"version_message"`
}

// @Title Get
// @Description get version
// @Success 200 {object} controllers.Version
// @router / [get]
func (u *VersionController) Get() {
	version := &Version{
		VersionTime:    "2018-07-27 --:-- PM",
		VersionNumber:  "v1.0.0",
		VersionMessage: "Hello from UOTC eco wallet service",
	}

	u.Data["json"] = version
	u.ServeJSON()
}
