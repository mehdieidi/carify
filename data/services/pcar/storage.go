package pcar

import (
	"context"

	"github.com/mehdieidi/carify/data/pkg/log"
	"github.com/mehdieidi/carify/data/pkg/xsql"
	"github.com/mehdieidi/carify/data/protocol"
	"github.com/mehdieidi/carify/data/protocol/derror"
)

const domain = "p_car"

type storage struct {
	db     xsql.TableOperator
	logger log.Logger
}

func NewStorage(db xsql.TableOperator, logger log.Logger) protocol.PCarStorage {
	return &storage{
		db:     db,
		logger: logger,
	}
}

func (s *storage) Store(ctx context.Context, pcar protocol.PCar) (protocol.PCarID, error) {
	const query = `
		INSERT INTO pcars (
			Year,
			Abi,
			Albaluyi,
			Atlasi,
			Bademjani,
			Boronz,
			Bezh,
			Banafsh,
			PoostPiyazi,
			Titanium,
			Khakestari,
			Khaki,
			Dolfini,
			Zoghali,
			Zard,
			Zereshki,
			Zeytooni,
			Sabz,
			Sorbi,
			Sormeyi,
			Sefid,
			SefidSadafi,
			Talayi,
			Toosi,
			Adasi,
			Annabi,
			Ghermez,
			Ghahveyi,
			CarbonBlack,
			Kerem,
			Gilasi,
			Mesi,
			Meshki,
			Moka,
			Narenji,
			NogrAbi,
			Noghreyi,
			NookMedadi,
			Yashmi,	
			usage_km,
			BadaneSalem,
			KhatKhashJozi,
			SafkariBirang,
			RangShodegi,
			DorRang,
			TamamRang,
			Tasadofi,
			Oraghi,
			cash_cost,
			MotorSalem,
			NiyazBeTaamir,
			TaavizShode,
			RearChasisSalem,
			RearZarbeKhorde,
			RearChasisRangShode,
			FrontChasisSalem,
			FrontZarbeKhorde,
			FrontChasisRangShode,
			third_party_insurance_due,
			Dandeyi,
			Automatic,
			car_token
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,
			$20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36,
			$37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53,
			$54, $55, $56, $57, $58, $59, $60, $61, $62  
		) RETURNING id
	`

	row := s.db.QueryRowContext(ctx, query,
		pcar.Year,
		pcar.Abi,
		pcar.Albaluyi,
		pcar.Atlasi,
		pcar.Bademjani,
		pcar.Boronz,
		pcar.Bezh,
		pcar.Banafsh,
		pcar.PoostPiyazi,
		pcar.Titanium,
		pcar.Khakestari,
		pcar.Khaki,
		pcar.Dolfini,
		pcar.Zoghali,
		pcar.Zard,
		pcar.Zereshki,
		pcar.Zeytooni,
		pcar.Sabz,
		pcar.Sorbi,
		pcar.Sormeyi,
		pcar.Sefid,
		pcar.SefidSadafi,
		pcar.Talayi,
		pcar.Toosi,
		pcar.Adasi,
		pcar.Annabi,
		pcar.Ghermez,
		pcar.Ghahveyi,
		pcar.CarbonBlack,
		pcar.Kerem,
		pcar.Gilasi,
		pcar.Mesi,
		pcar.Meshki,
		pcar.Moka,
		pcar.Narenji,
		pcar.NogrAbi,
		pcar.Noghreyi,
		pcar.NookMedadi,
		pcar.Yashmi,
		pcar.UsageKM,
		pcar.BadaneSalem,
		pcar.KhatKhashJozi,
		pcar.SafkariBirang,
		pcar.RangShodegi,
		pcar.DorRang,
		pcar.TamamRang,
		pcar.Tasadofi,
		pcar.Oraghi,
		pcar.CashCost,
		pcar.MotorSalem,
		pcar.NiyazBeTaamir,
		pcar.TaavizShode,
		pcar.RearChasisSalem,
		pcar.RearZarbeKhorde,
		pcar.RearChasisRangShode,
		pcar.FrontChasisSalem,
		pcar.FrontZarbeKhorde,
		pcar.FrontChasisRangShode,
		pcar.ThirdPartyInsuranceDue,
		pcar.Dandeyi,
		pcar.Automatic,
		pcar.Token,
	)

	err := row.Scan(&pcar.ID)
	if err != nil {
		s.logger.Error(domain, log.StorageLayer, "Store", log.Args{log.LogErrKey: err})
		return 0, derror.ErrUnexpected
	}

	return pcar.ID, nil
}
