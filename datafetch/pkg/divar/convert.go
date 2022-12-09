package divar

import (
	"fmt"
)

func ToColor(color string) (Color, error) {
	if color == "" {
		return UnsetColor, nil
	}

	var c Color
	if err := c.UnmarshalText([]byte(color)); err != nil {
		return UnsetColor, err
	}

	return c, nil
}

func ToMotorStatus(motorStatus string) (MotorStatus, error) {
	if motorStatus == "" {
		return UnsetMotorStatus, nil
	}

	var m MotorStatus
	if err := m.UnmarshalText([]byte(motorStatus)); err != nil {
		return UnsetMotorStatus, err
	}

	return m, nil
}

func ToChassisStatus(chassisStatus string) (front ChassisStatus, rear ChassisStatus, err error) {
	switch chassisStatus {
	case "هر دو سالم و پلمپ":
		front = ChasisSalem
		rear = ChasisSalem
	case "عقب ضربه‌خورده":
		front = ChasisSalem
		rear = ZarbeKhorde
	case "عقب رنگ‌شده":
		front = ChasisSalem
		rear = ChasisRangShode
	case "جلو ضربه‌خورده":
		front = ZarbeKhorde
		rear = ChasisSalem
	case "جلو رنگ‌شده":
		front = ChasisRangShode
		rear = ChasisSalem
	case "عقب ضربه‌خورده، جلو رنگ‌شده":
		front = ChasisRangShode
		rear = ZarbeKhorde
	case "عقب رنگ‌شده، جلو ضربه‌خورده":
		front = ZarbeKhorde
		rear = ChasisRangShode
	case "هردو ضربه‌خورده":
		front = ZarbeKhorde
		rear = ZarbeKhorde
	case "هردو رنگ‌شده":
		front = ChasisRangShode
		rear = ChasisRangShode
	case "ضربه‌خورده":
		front = ZarbeKhorde
		rear = ZarbeKhorde
	case "رنگ‌شده":
		front = ChasisRangShode
		rear = ChasisRangShode
	case "سالم و پلمپ":
		front = ChasisSalem
		rear = ChasisSalem
	case "":
		front = UnsetChasisStatus
		rear = UnsetChasisStatus
	default:
		return UnsetChasisStatus, UnsetChasisStatus, fmt.Errorf("invalid chassis status: %+v", chassisStatus)
	}
	return
}

func ToBodyStatus(bodyStatus string) (BodyStatus, error) {
	if bodyStatus == "" {
		return UnsetBodyStatus, nil
	}

	var b BodyStatus
	if err := b.UnmarshalText([]byte(bodyStatus)); err != nil {
		return UnsetBodyStatus, err
	}

	return b, nil
}

func ToGearbox(gearbox string) (Gearbox, error) {
	if gearbox == "" {
		return UnsetGearBox, nil
	}

	var g Gearbox
	if err := g.UnmarshalText([]byte(gearbox)); err != nil {
		return UnsetGearBox, err
	}

	return g, nil
}
