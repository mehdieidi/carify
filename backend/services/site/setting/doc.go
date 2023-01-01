package setting

// @Summary     get site settings
// @Description get site settings.
// @Tags        SiteSetting
// @Produce     json
// @Success     200 {object} transport.Response{Data=protocol.SiteSetting}
// @Failure     400 {object} map[string]string{error=string} "Invalid request"
// @Failure     500 {object} map[string]string{error=string} "Internal server error"
// @Failure     404 {object} map[string]string{error=string} "site setting not found"
// @Router      /site/settings/get [get]
func _() {}
