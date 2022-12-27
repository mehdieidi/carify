package divar

import "errors"

var (
	ErrUnknownColor         = errors.New("unknown color")
	ErrUnknownBodyStatus    = errors.New("unknown body status")
	ErrUnknownMotorStatus   = errors.New("unknown motor status")
	ErrUnknownChassisStatus = errors.New("unknown chassis status")
	ErrUnknownGearBox       = errors.New("unknown gearbox")
)
