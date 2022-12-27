package protocol

import "context"

type PCarID uint

// PCar represents one-hot-encoded Car.
type PCar struct {
	ID                     PCarID
	Token                  string
	Year                   int
	UsageKM                int
	Abi                    int
	Albaluyi               int
	Atlasi                 int
	Bademjani              int
	Boronz                 int
	Bezh                   int
	Banafsh                int
	PoostPiyazi            int
	Titanium               int
	Khakestari             int
	Khaki                  int
	Dolfini                int
	Zoghali                int
	Zard                   int
	Zereshki               int
	Zeytooni               int
	Sabz                   int
	Sorbi                  int
	Sormeyi                int
	Sefid                  int
	SefidSadafi            int
	Talayi                 int
	Toosi                  int
	Adasi                  int
	Annabi                 int
	Ghermez                int
	Ghahveyi               int
	CarbonBlack            int
	Kerem                  int
	Gilasi                 int
	Mesi                   int
	Meshki                 int
	Moka                   int
	Narenji                int
	NogrAbi                int
	Noghreyi               int
	NookMedadi             int
	Yashmi                 int
	BadaneSalem            int
	KhatKhashJozi          int
	SafkariBirang          int
	RangShodegi            int
	DorRang                int
	TamamRang              int
	Tasadofi               int
	Oraghi                 int
	CashCost               int
	MotorSalem             int
	NiyazBeTaamir          int
	TaavizShode            int
	RearChasisSalem        int
	RearZarbeKhorde        int
	RearChasisRangShode    int
	FrontChasisSalem       int
	FrontZarbeKhorde       int
	FrontChasisRangShode   int
	ThirdPartyInsuranceDue int
	Dandeyi                int
	Automatic              int
}

type PCarStorage interface {
	Store(context.Context, PCar) (PCarID, error)
}

type PCarService interface {
	OneHotEncode(context.Context) error
}
