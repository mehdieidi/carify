package setting

import (
	"back/protocol"
	"context"
	"encoding/json"
)

const domain = "site_setting"

type service struct {
	setting string
}

func NewService(setting string) protocol.SiteSettingService {
	return &service{
		setting: setting,
	}
}

func (s *service) Get(ctx context.Context) (protocol.SiteSetting, error) {
	var setting protocol.SiteSetting
	if err := json.Unmarshal([]byte(s.setting), &setting); err != nil {
		return protocol.SiteSetting{}, err
	}
	return setting, nil
}
