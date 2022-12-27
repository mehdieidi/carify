package pcar

import (
	"context"
	"errors"

	"github.com/mehdieidi/carify/data/pkg/divar"
	"github.com/mehdieidi/carify/data/protocol"
)

type service struct {
	pcarStorage protocol.PCarStorage

	preprocessService protocol.PreProcessService
}

func NewService(pcarStorage protocol.PCarStorage, preprocessService protocol.PreProcessService) protocol.PCarService {
	return &service{pcarStorage: pcarStorage, preprocessService: preprocessService}
}

func (s *service) OneHotEncode(ctx context.Context) error {
	cars, err := s.preprocessService.List(ctx)
	if err != nil {
		return err
	}

	for _, c := range cars {
		pcar := protocol.PCar{
			Token:                  c.Token,
			Year:                   c.Year,
			UsageKM:                c.UsageKM,
			CashCost:               c.CashCost,
			ThirdPartyInsuranceDue: c.ThirdPartyInsuranceDue,
		}

		pcar, err = oneHotColor(pcar, c.Color)
		if err != nil {
			return err
		}

		pcar, err = oneHotBodyStatus(pcar, c.BodyStatus)
		if err != nil {
			return err
		}

		pcar, err = oneHotMotorStatus(pcar, c.MotorStatus)
		if err != nil {
			return err
		}

		pcar, err = oneHotRearChassisStatus(pcar, c.RearChassisStatus)
		if err != nil {
			return err
		}

		pcar, err = oneHotFrontChassisStatus(pcar, c.FrontChassisStatus)
		if err != nil {
			return err
		}

		pcar, err = oneHotGearbox(pcar, c.Gearbox)
		if err != nil {
			return err
		}

		if _, err := s.pcarStorage.Store(ctx, pcar); err != nil {
			return err
		}
	}

	return nil
}

func oneHotColor(pcar protocol.PCar, color divar.Color) (protocol.PCar, error) {
	switch color {
	case divar.Abi:
		pcar.Abi = 1
	case divar.Albaluyi:
		pcar.Albaluyi = 1
	case divar.Atlasi:
		pcar.Atlasi = 1
	case divar.Bademjani:
		pcar.Bademjani = 1
	case divar.Boronz:
		pcar.Boronz = 1
	case divar.Bezh:
		pcar.Bezh = 1
	case divar.Banafsh:
		pcar.Banafsh = 1
	case divar.PoostPiyazi:
		pcar.PoostPiyazi = 1
	case divar.Titanium:
		pcar.Titanium = 1
	case divar.Khakestari:
		pcar.Khakestari = 1
	case divar.Khaki:
		pcar.Khaki = 1
	case divar.Dolfini:
		pcar.Dolfini = 1
	case divar.Zoghali:
		pcar.Zoghali = 1
	case divar.Zard:
		pcar.Zard = 1
	case divar.Zereshki:
		pcar.Zereshki = 1
	case divar.Zeytooni:
		pcar.Zeytooni = 1
	case divar.Sabz:
		pcar.Sabz = 1
	case divar.Sorbi:
		pcar.Sorbi = 1
	case divar.Sormeyi:
		pcar.Sormeyi = 1
	case divar.Sefid:
		pcar.Sefid = 1
	case divar.SefidSadafi:
		pcar.SefidSadafi = 1
	case divar.Talayi:
		pcar.Talayi = 1
	case divar.Toosi:
		pcar.Toosi = 1
	case divar.Adasi:
		pcar.Adasi = 1
	case divar.Annabi:
		pcar.Annabi = 1
	case divar.Ghermez:
		pcar.Ghermez = 1
	case divar.Ghahveyi:
		pcar.Ghahveyi = 1
	case divar.CarbonBlack:
		pcar.CarbonBlack = 1
	case divar.Kerem:
		pcar.Kerem = 1
	case divar.Gilasi:
		pcar.Gilasi = 1
	case divar.Mesi:
		pcar.Mesi = 1
	case divar.Meshki:
		pcar.Meshki = 1
	case divar.Moka:
		pcar.Moka = 1
	case divar.Narenji:
		pcar.Narenji = 1
	case divar.NogrAbi:
		pcar.NogrAbi = 1
	case divar.Noghreyi:
		pcar.Noghreyi = 1
	case divar.NookMedadi:
		pcar.NookMedadi = 1
	case divar.Yashmi:
		pcar.Yashmi = 1
	default:
		return protocol.PCar{}, errors.New("unknown color")
	}

	return pcar, nil
}

func oneHotBodyStatus(pcar protocol.PCar, bodyStatus divar.BodyStatus) (protocol.PCar, error) {
	switch bodyStatus {
	case divar.BadaneSalem:
		pcar.BadaneSalem = 1
	case divar.KhatKhashJozi:
		pcar.KhatKhashJozi = 1
	case divar.SafkariBirang:
		pcar.SafkariBirang = 1
	case divar.RangShodegi:
		pcar.RangShodegi = 1
	case divar.DorRang:
		pcar.DorRang = 1
	case divar.TamamRang:
		pcar.TamamRang = 1
	case divar.Tasadofi:
		pcar.Tasadofi = 1
	case divar.Oraghi:
		pcar.Oraghi = 1
	default:
		return protocol.PCar{}, errors.New("unknown body status")
	}

	return pcar, nil
}

func oneHotMotorStatus(pcar protocol.PCar, motorStatus divar.MotorStatus) (protocol.PCar, error) {
	switch motorStatus {
	case divar.MotorSalem:
		pcar.MotorSalem = 1
	case divar.NiyazBeTaamir:
		pcar.NiyazBeTaamir = 1
	case divar.TaavizShode:
		pcar.TaavizShode = 1
	default:
		return protocol.PCar{}, errors.New("unknown motor status")
	}

	return pcar, nil
}

func oneHotRearChassisStatus(pcar protocol.PCar, rearChassisStatus divar.ChassisStatus) (protocol.PCar, error) {
	switch rearChassisStatus {
	case divar.ChasisSalem:
		pcar.RearChasisSalem = 1
	case divar.ZarbeKhorde:
		pcar.RearZarbeKhorde = 1
	case divar.ChasisRangShode:
		pcar.RearChasisRangShode = 1
	default:
		return protocol.PCar{}, errors.New("unknown rear chassis status")
	}

	return pcar, nil
}

func oneHotFrontChassisStatus(pcar protocol.PCar, frontChassisStatus divar.ChassisStatus) (protocol.PCar, error) {
	switch frontChassisStatus {
	case divar.ChasisSalem:
		pcar.FrontChasisSalem = 1
	case divar.ZarbeKhorde:
		pcar.FrontZarbeKhorde = 1
	case divar.ChasisRangShode:
		pcar.FrontChasisRangShode = 1
	default:
		return protocol.PCar{}, errors.New("unknown front chassis status")
	}

	return pcar, nil
}

func oneHotGearbox(pcar protocol.PCar, gearbox divar.Gearbox) (protocol.PCar, error) {
	switch gearbox {
	case divar.Dandeyi:
		pcar.Dandeyi = 1
	case divar.Automatic:
		pcar.Automatic = 1
	default:
		return protocol.PCar{}, errors.New("unknown gearbox")
	}

	return pcar, nil
}
