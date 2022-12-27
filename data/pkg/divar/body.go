package divar

type BodyStatus uint

const (
	UnsetBodyStatus BodyStatus = iota
	BadaneSalem
	KhatKhashJozi
	SafkariBirang
	RangShodegi
	DorRang
	TamamRang
	Tasadofi
	Oraghi
)

func (b BodyStatus) String() string {
	switch b {
	case UnsetBodyStatus:
		return "unset"
	case BadaneSalem:
		return "سالم و بی‌خط و خش"
	case KhatKhashJozi:
		return "خط و خش جزیی"
	case SafkariBirang:
		return "صافکاری بی‌رنگ"
	case RangShodegi:
		return "رنگ‌شدگی"
	case DorRang:
		return "دوررنگ"
	case TamamRang:
		return "تمام‌رنگ"
	case Tasadofi:
		return "تصادفی"
	case Oraghi:
		return "اوراقی"
	default:
		return ""
	}
}

func (b BodyStatus) MarshalText() ([]byte, error) {
	if s := b.String(); s != "" {
		return []byte(s), nil
	}
	return nil, ErrUnknownBodyStatus
}

func (b *BodyStatus) UnmarshalText(p []byte) error {
	switch string(p) {
	case UnsetBodyStatus.String():
		*b = UnsetBodyStatus
	case BadaneSalem.String():
		*b = BadaneSalem
	case KhatKhashJozi.String():
		*b = KhatKhashJozi
	case SafkariBirang.String():
		*b = SafkariBirang
	case RangShodegi.String():
		*b = RangShodegi
	case DorRang.String():
		*b = DorRang
	case TamamRang.String():
		*b = TamamRang
	case Tasadofi.String():
		*b = Tasadofi
	case Oraghi.String():
		*b = Oraghi
	default:
		return ErrUnknownBodyStatus
	}
	return nil
}
