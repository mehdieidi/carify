package divar

type Color uint

const (
	UnsetColor Color = iota
	Abi
	Albaluyi
	Atlasi
	Bademjani
	Boronz
	Bezh
	Banafsh
	PoostPiyazi
	Titanium
	Khakestari
	Khaki
	Dolfini
	Zoghali
	Zard
	Zereshki
	Zeytooni
	Sabz
	Sorbi
	Sormeyi
	Sefid
	SefidSadafi
	Talayi
	Toosi
	Adasi
	Annabi
	Ghermez
	Ghahveyi
	CarbonBlack
	Kerem
	Gilasi
	Mesi
	Meshki
	Moka
	Narenji
	NogrAbi
	Noghreyi
	NookMedadi
	Yashmi
)

func (c Color) String() string {
	switch c {
	case UnsetColor:
		return "unset"
	case Abi:
		return "آبی"
	case Albaluyi:
		return "آلبالویی"
	case Atlasi:
		return "اطلسی"
	case Bademjani:
		return "بادمجانی"
	case Boronz:
		return "برنز"
	case Bezh:
		return "بژ"
	case Banafsh:
		return "بنفش"
	case PoostPiyazi:
		return "پوست‌پیازی"
	case Titanium:
		return "تیتانیوم"
	case Khakestari:
		return "خاکستری"
	case Khaki:
		return "خاکی"
	case Dolfini:
		return "دلفینی"
	case Zoghali:
		return "ذغالی"
	case Zard:
		return "زرد"
	case Zereshki:
		return "زرشکی"
	case Zeytooni:
		return "زیتونی"
	case Sabz:
		return "سبز"
	case Sorbi:
		return "سربی"
	case Sormeyi:
		return "سرمه‌ای"
	case Sefid:
		return "سفید"
	case SefidSadafi:
		return "سفید صدفی"
	case Talayi:
		return "طلایی"
	case Toosi:
		return "طوسی"
	case Adasi:
		return "عدسی"
	case Annabi:
		return "عنابی"
	case Ghermez:
		return "قرمز"
	case Ghahveyi:
		return "قهوه‌ای"
	case CarbonBlack:
		return "کربن‌بلک"
	case Kerem:
		return "کرم"
	case Gilasi:
		return "گیلاسی"
	case Mesi:
		return "مسی"
	case Meshki:
		return "مشکی"
	case Moka:
		return "موکا"
	case Narenji:
		return "نارنجی"
	case NogrAbi:
		return "نقرآبی"
	case Noghreyi:
		return "نقره‌ای"
	case NookMedadi:
		return "نوک‌مدادی"
	case Yashmi:
		return "یشمی"
	default:
		return ""
	}
}

func (c Color) MarshalText() ([]byte, error) {
	if s := c.String(); s != "" {
		return []byte(s), nil
	}
	return nil, ErrUnknownColor
}

func (c *Color) UnmarshalText(p []byte) error {
	switch string(p) {
	case UnsetColor.String():
		*c = UnsetColor
	case Abi.String():
		*c = Abi
	case Albaluyi.String():
		*c = Albaluyi
	case Atlasi.String():
		*c = Atlasi
	case Bademjani.String():
		*c = Bademjani
	case Boronz.String():
		*c = Boronz
	case Bezh.String():
		*c = Bezh
	case Banafsh.String():
		*c = Banafsh
	case PoostPiyazi.String():
		*c = PoostPiyazi
	case Titanium.String():
		*c = Titanium
	case Khakestari.String():
		*c = Khakestari
	case Khaki.String():
		*c = Khaki
	case Dolfini.String():
		*c = Dolfini
	case Zoghali.String():
		*c = Zoghali
	case Zard.String():
		*c = Zard
	case Zereshki.String():
		*c = Zereshki
	case Zeytooni.String():
		*c = Zeytooni
	case Sabz.String():
		*c = Sabz
	case Sorbi.String():
		*c = Sorbi
	case Sormeyi.String():
		*c = Sormeyi
	case Sefid.String():
		*c = Sefid
	case SefidSadafi.String():
		*c = SefidSadafi
	case Talayi.String():
		*c = Talayi
	case Toosi.String():
		*c = Toosi
	case Adasi.String():
		*c = Adasi
	case Annabi.String():
		*c = Annabi
	case Ghermez.String():
		*c = Ghermez
	case Ghahveyi.String():
		*c = Ghahveyi
	case CarbonBlack.String():
		*c = CarbonBlack
	case Kerem.String():
		*c = Kerem
	case Gilasi.String():
		*c = Gilasi
	case Mesi.String():
		*c = Mesi
	case Meshki.String():
		*c = Meshki
	case Moka.String():
		*c = Moka
	case Narenji.String():
		*c = Narenji
	case NogrAbi.String():
		*c = NogrAbi
	case Noghreyi.String():
		*c = Noghreyi
	case NookMedadi.String():
		*c = NookMedadi
	case Yashmi.String():
		*c = Yashmi
	default:
		return ErrUnknownColor
	}
	return nil
}
