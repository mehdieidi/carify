package divar

type Gearbox uint8

const (
	UnsetGearBox Gearbox = iota
	Dandeyi
	Automatic
)

func (g Gearbox) String() string {
	switch g {
	case UnsetGearBox:
		return "unset"
	case Dandeyi:
		return "دنده‌ای"
	case Automatic:
		return "اتوماتیک"
	default:
		return ""
	}
}

func (g Gearbox) MarshalText() ([]byte, error) {
	if s := g.String(); s != "" {
		return []byte(s), nil
	}
	return nil, ErrUnknownGearBox
}

func (g *Gearbox) UnmarshalText(p []byte) error {
	switch string(p) {
	case UnsetGearBox.String():
		*g = UnsetGearBox
	case Dandeyi.String():
		*g = Dandeyi
	case Automatic.String():
		*g = Automatic
	default:
		return ErrUnknownGearBox
	}
	return nil
}
