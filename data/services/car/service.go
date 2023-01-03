package car

import (
	"context"
	"time"

	"github.com/mehdieidi/carify/data/pkg/divar"
	"github.com/mehdieidi/carify/data/pkg/log"
	"github.com/mehdieidi/carify/data/pkg/pnum"
	"github.com/mehdieidi/carify/data/protocol"
)

type service struct {
	carStorage protocol.CarStorage

	logger log.Logger
}

func NewService(carStorage protocol.CarStorage, logger log.Logger) protocol.CarService {
	return &service{carStorage: carStorage, logger: logger}
}

func (s *service) Search(ctx context.Context, cities []string, brandModel string) ([]protocol.CarToken, error) {
	carTokens := []protocol.CarToken{}

	searchReq := divar.NewSearchRequest("light", brandModel, cities, 0)

	for {
		time.Sleep(1 * time.Second)

		searchResp, err := divar.Search(ctx, searchReq)
		if err != nil {
			s.logger.Error(domain, log.ServiceLayer, "Search", log.Args{log.LogErrKey: err})
			continue
		}

		// All posts are retrieved.
		if len(searchResp.WebWidgets.PostList) == 0 {
			break
		}

		searchReq.LastPostDate = searchResp.LastPostDate

		for _, t := range searchResp.WebWidgets.PostList {
			carTokens = append(carTokens, protocol.CarToken(t.Data.Action.Payload.Token))
		}
	}

	return carTokens, nil
}

func (s *service) Get(ctx context.Context, carToken protocol.CarToken) (protocol.Car, error) {
	getResp, err := divar.Get(ctx, string(carToken))
	if err != nil {
		return protocol.Car{}, err
	}

	var widgets []divar.Widget

	for _, sec := range getResp.Sections {
		if sec.SectionName == "BREADCRUMB" {
			continue
		}
		if sec.SectionName == "TITLE" {
			continue
		}
		if sec.SectionName == "DESCRIPTION" {
			continue
		}
		if sec.SectionName == "IMAGE" {
			continue
		}

		if sec.SectionName == "LIST_DATA" {
			widgets = sec.Widgets
			break
		}
	}

	var usageStr string
	var yearStr string
	var colorStr string
	var motorStatusStr string
	var chassisStatusStr string
	var frontChassisStatusStr string
	var rearChassisStatusStr string
	var bodyStatusStr string
	var insuranceDueStr string
	var gearboxStr string
	var costStr string

	for _, w := range widgets {
		if w.WidgetType == "GROUP_INFO_ROW" {
			usageStr = w.Data.Items[0].Value
			yearStr = w.Data.Items[1].Value
			colorStr = w.Data.Items[2].Value
		}

		switch w.Data.Title {
		case "وضعیت موتور":
			motorStatusStr = w.Data.Value
		case "وضعیت شاسی‌ها":
			chassisStatusStr = w.Data.Value
		case "وضعیت بدنه":
			bodyStatusStr = w.Data.Value
		case "مهلت بیمهٔ شخص ثالث":
			insuranceDueStr = w.Data.Value
		case "گیربکس":
			gearboxStr = w.Data.Value
		case "قیمت فروش نقدی":
			costStr = w.Data.Value
		case "شاسی جلو":
			frontChassisStatusStr = w.Data.Value
		case "شاسی عقب":
			rearChassisStatusStr = w.Data.Value
		}
	}

	usageKM, err := pnum.ToInt(usageStr)
	if err != nil {
		return protocol.Car{}, err
	}

	year, err := pnum.ToInt(yearStr)
	if err != nil {
		return protocol.Car{}, err
	}

	color, err := divar.ToColor(colorStr)
	if err != nil {
		return protocol.Car{}, err
	}

	motorStatus, err := divar.ToMotorStatus(motorStatusStr)
	if err != nil {
		return protocol.Car{}, err
	}

	frontChassisStatus, rearChassisStatus, err := divar.ToChassisStatus(chassisStatusStr)
	if err != nil {
		frontChassisStatus, _, err = divar.ToChassisStatus(frontChassisStatusStr)
		if err != nil {
			return protocol.Car{}, err
		}

		rearChassisStatus, _, err = divar.ToChassisStatus(rearChassisStatusStr)
		if err != nil {
			return protocol.Car{}, err
		}
	}

	bodyStatus, err := divar.ToBodyStatus(bodyStatusStr)
	if err != nil {
		return protocol.Car{}, err
	}

	insuranceDue, err := pnum.ToInt(insuranceDueStr)
	if err != nil {
		return protocol.Car{}, err
	}

	gearbox, err := divar.ToGearbox(gearboxStr)
	if err != nil {
		return protocol.Car{}, err
	}

	cost, err := pnum.ToInt(costStr)
	if err != nil {
		return protocol.Car{}, err
	}

	c := protocol.Car{
		Token:                  string(carToken),
		Year:                   year,
		Color:                  color,
		UsageKM:                usageKM,
		BodyStatus:             bodyStatus,
		CashCost:               cost,
		MotorStatus:            motorStatus,
		FrontChassisStatus:     frontChassisStatus,
		RearChassisStatus:      rearChassisStatus,
		ThirdPartyInsuranceDue: insuranceDue,
		Gearbox:                gearbox,
	}

	return c, nil
}

func (s *service) GetByID(ctx context.Context, id protocol.CarID) (protocol.Car, error) {
	return s.carStorage.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, id protocol.CarID, car protocol.Car) error {
	c, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if car.Token != "" {
		c.Token = car.Token
	}

	if car.Year != 0 {
		c.Year = car.Year
	}

	if car.Color != 0 {
		c.Color = car.Color
	}

	if car.UsageKM != 0 {
		c.UsageKM = car.UsageKM
	}

	if car.BodyStatus != 0 {
		c.BodyStatus = car.BodyStatus
	}

	if car.CashCost != 0 {
		c.CashCost = car.CashCost
	}

	if car.MotorStatus != 0 {
		c.MotorStatus = car.MotorStatus
	}

	if car.FrontChassisStatus != 0 {
		c.FrontChassisStatus = car.FrontChassisStatus
	}

	if car.RearChassisStatus != 0 {
		c.RearChassisStatus = car.RearChassisStatus
	}

	if car.ThirdPartyInsuranceDue != 0 {
		c.ThirdPartyInsuranceDue = car.ThirdPartyInsuranceDue
	}

	if car.Gearbox != 0 {
		c.Gearbox = car.Gearbox
	}

	if err := s.carStorage.Update(ctx, id, c); err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id protocol.CarID) error {
	return s.carStorage.Delete(ctx, id)
}
