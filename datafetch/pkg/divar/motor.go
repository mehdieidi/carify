package divar

type MotorStatus uint8

const (
	UnsetMotorStatus MotorStatus = iota
	MotorSalem
	NiyazBeTaamir
	TaavizShode
)

func (m MotorStatus) String() string {
	switch m {
	case UnsetMotorStatus:
		return "unset"
	case MotorSalem:
		return "سالم"
	case NiyazBeTaamir:
		return "نیاز به تعمیر"
	case TaavizShode:
		return "تعویض شده"
	default:
		return ""
	}
}

func (m MotorStatus) MarshalText() ([]byte, error) {
	if s := m.String(); s != "" {
		return []byte(s), nil
	}
	return nil, ErrUnknownMotorStatus
}

func (m *MotorStatus) UnmarshalText(p []byte) error {
	switch string(p) {
	case UnsetMotorStatus.String():
		*m = UnsetMotorStatus
	case MotorSalem.String():
		*m = MotorSalem
	case NiyazBeTaamir.String():
		*m = NiyazBeTaamir
	case TaavizShode.String():
		*m = TaavizShode
	default:
		return ErrUnknownMotorStatus
	}
	return nil
}
