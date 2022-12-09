package divar

type ChassisStatus uint8

const (
	UnsetChasisStatus ChassisStatus = iota
	ChasisSalem
	ZarbeKhorde
	ChasisRangShode
)

func (c ChassisStatus) String() string {
	switch c {
	case UnsetChasisStatus:
		return "unset"
	case ChasisSalem:
		return "سالم"
	case ZarbeKhorde:
		return "ضربه خورده"
	case ChasisRangShode:
		return "رنگ شده"
	default:
		return ""
	}
}

func (c ChassisStatus) MarshalText() ([]byte, error) {
	if s := c.String(); s != "" {
		return []byte(s), nil
	}
	return nil, ErrUnknownChassisStatus
}

func (c *ChassisStatus) UnmarshalText(p []byte) error {
	switch string(p) {
	case UnsetChasisStatus.String():
		*c = UnsetChasisStatus
	case ChasisSalem.String():
		*c = ChasisSalem
	case ZarbeKhorde.String():
		*c = ZarbeKhorde
	case ChasisRangShode.String():
		*c = ChasisRangShode
	default:
		return ErrUnknownChassisStatus
	}
	return nil
}
