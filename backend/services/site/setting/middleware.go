package setting

import "back/protocol"

type Middleware func(protocol.SiteSettingService) protocol.SiteSettingService
