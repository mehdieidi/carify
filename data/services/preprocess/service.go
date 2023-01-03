package preprocess

import (
	"context"
	"fmt"

	"github.com/mehdieidi/carify/data/pkg/divar"
	"github.com/mehdieidi/carify/data/protocol"
)

type service struct {
	preprocessStorage protocol.PreProcessStorage

	carService protocol.CarService
}

func NewService(preprocessStorage protocol.PreProcessStorage, carService protocol.CarService) protocol.PreProcessService {
	return &service{
		preprocessStorage: preprocessStorage,
		carService:        carService,
	}
}

func (s *service) List(ctx context.Context) ([]protocol.Car, error) {
	return s.preprocessStorage.FindAll(ctx)
}

func (s *service) Year(ctx context.Context, minYear, maxYear int) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	carYearTypesCount := map[int]int{}

	for _, c := range cars {
		if c.Year >= minYear && c.Year <= maxYear {
			carYearTypesCount[c.Year]++
		}
	}

	mostCommon := 1390
	for year, count := range carYearTypesCount {
		if count > carYearTypesCount[mostCommon] {
			mostCommon = year
		}
	}

	fmt.Println("most common year is", mostCommon, "count =", carYearTypesCount[mostCommon])

	preProcessedCount := 0
	for _, c := range cars {
		if c.Year < minYear || c.Year > maxYear {
			fmt.Println("car", c.ID, "has wrong year:", c.Year)

			err := s.carService.Update(ctx, c.ID, protocol.Car{Year: mostCommon})
			if err != nil {
				return err
			}

			preProcessedCount++
		}
	}

	fmt.Println("fixed", preProcessedCount, "records year attribute")

	return nil
}

func (s *service) Color(ctx context.Context) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	carColorTypesCount := map[int]int{}

	for _, c := range cars {
		if c.Color != 0 {
			carColorTypesCount[int(c.Color)]++
		}
	}

	mostCommon := 1
	for color, count := range carColorTypesCount {
		if count > carColorTypesCount[mostCommon] {
			mostCommon = color
		}
	}

	fmt.Println("most common color is", mostCommon, "count =", carColorTypesCount[mostCommon])

	preProcessedCount := 0
	for _, c := range cars {
		if c.Color == 0 {
			fmt.Println("car", c.ID, "has wrong color:", c.Color)

			err := s.carService.Update(ctx, c.ID, protocol.Car{Color: divar.Color(mostCommon)})
			if err != nil {
				return err
			}

			preProcessedCount++
		}
	}

	fmt.Println("fixed", preProcessedCount, "records color attribute")

	return nil
}

func (s *service) UsageKM(ctx context.Context) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	const maxUsageKM = 600_000

	sumUsageKM := 0
	count := 0

	for _, c := range cars {
		if c.UsageKM <= maxUsageKM && c.UsageKM > 0 {
			sumUsageKM += c.UsageKM
			count++
		}
	}

	avgUsageKM := sumUsageKM / count

	fmt.Println("avg usage km is:", avgUsageKM)

	preProcessedCount := 0
	for _, c := range cars {
		if c.UsageKM <= 0 || c.UsageKM > maxUsageKM {
			fmt.Println("car", c.ID, "has wrong usageKM:", c.UsageKM)

			err := s.carService.Update(ctx, c.ID, protocol.Car{UsageKM: avgUsageKM})
			if err != nil {
				return err
			}

			preProcessedCount++
		}
	}

	fmt.Println("fixed", preProcessedCount, "records usageKM attribute")

	return nil
}

func (s *service) BodyStatus(ctx context.Context) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	carBodyStatusTypesCount := map[int]int{}

	for _, c := range cars {
		if c.BodyStatus != 0 {
			carBodyStatusTypesCount[int(c.BodyStatus)]++
		}
	}

	mostCommon := 1
	for bodyStatus, count := range carBodyStatusTypesCount {
		if count > carBodyStatusTypesCount[mostCommon] {
			mostCommon = bodyStatus
		}
	}

	fmt.Println("most common body status is", mostCommon, "count =", carBodyStatusTypesCount[mostCommon])

	preProcessedCount := 0
	for _, c := range cars {
		if c.BodyStatus == 0 {
			fmt.Println("car", c.ID, "has wrong body status:", c.BodyStatus)

			err := s.carService.Update(ctx, c.ID, protocol.Car{BodyStatus: divar.BodyStatus(mostCommon)})
			if err != nil {
				return err
			}

			preProcessedCount++
		}
	}

	fmt.Println("fixed", preProcessedCount, "records body status attribute")

	return nil
}

func (s *service) CashCost(ctx context.Context) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	const maxCost = 260_000_000
	const minCost = 100_000_000

	wrongCostCount := 0

	for _, c := range cars {
		if c.CashCost > maxCost || c.CashCost < minCost {
			err := s.carService.Delete(ctx, c.ID)
			if err != nil {
				return err
			}

			fmt.Println("wrong cost", c.CashCost)

			wrongCostCount++
		}
	}

	fmt.Println("fixed", wrongCostCount, "records cost attribute")

	return nil
}

func (s *service) MotorStatus(ctx context.Context) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	carMotorStatusTypesCount := map[int]int{}

	for _, c := range cars {
		if c.MotorStatus != 0 {
			carMotorStatusTypesCount[int(c.MotorStatus)]++
		}
	}

	mostCommon := 1
	for motorStatus, count := range carMotorStatusTypesCount {
		if count > carMotorStatusTypesCount[mostCommon] {
			mostCommon = motorStatus
		}
	}

	fmt.Println("most common motor status is", mostCommon, "count =", carMotorStatusTypesCount[mostCommon])

	preProcessedCount := 0
	for _, c := range cars {
		if c.MotorStatus == 0 {
			fmt.Println("car", c.ID, "has wrong motor status:", c.MotorStatus)

			err := s.carService.Update(ctx, c.ID, protocol.Car{MotorStatus: divar.MotorStatus(mostCommon)})
			if err != nil {
				return err
			}

			preProcessedCount++
		}
	}

	fmt.Println("fixed", preProcessedCount, "records motor status attribute")

	return nil
}

func (s *service) FrontChassisStatus(ctx context.Context) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	carFrontChassisStatusTypesCount := map[int]int{}

	for _, c := range cars {
		if c.FrontChassisStatus != 0 {
			carFrontChassisStatusTypesCount[int(c.FrontChassisStatus)]++
		}
	}

	mostCommon := 1
	for frontChassisStatus, count := range carFrontChassisStatusTypesCount {
		if count > carFrontChassisStatusTypesCount[mostCommon] {
			mostCommon = frontChassisStatus
		}
	}

	fmt.Println("most common front chassis status is", mostCommon, "count =", carFrontChassisStatusTypesCount[mostCommon])

	preProcessedCount := 0
	for _, c := range cars {
		if c.FrontChassisStatus == 0 {
			fmt.Println("car", c.ID, "has wrong front chassis status:", c.FrontChassisStatus)

			err := s.carService.Update(ctx, c.ID, protocol.Car{FrontChassisStatus: divar.ChassisStatus(mostCommon)})
			if err != nil {
				return err
			}

			preProcessedCount++
		}
	}

	fmt.Println("fixed", preProcessedCount, "records front chassis status attribute")

	return nil
}

func (s *service) RearChassisStatus(ctx context.Context) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	carRearChassisStatusTypesCount := map[int]int{}

	for _, c := range cars {
		if c.RearChassisStatus != 0 {
			carRearChassisStatusTypesCount[int(c.RearChassisStatus)]++
		}
	}

	mostCommon := 1
	for rearChassisStatus, count := range carRearChassisStatusTypesCount {
		if count > carRearChassisStatusTypesCount[mostCommon] {
			mostCommon = rearChassisStatus
		}
	}

	fmt.Println("most common rear chassis status is", mostCommon, "count =", carRearChassisStatusTypesCount[mostCommon])

	preProcessedCount := 0
	for _, c := range cars {
		if c.RearChassisStatus == 0 {
			fmt.Println("car", c.ID, "has wrong rear chassis status:", c.RearChassisStatus)

			err := s.carService.Update(ctx, c.ID, protocol.Car{RearChassisStatus: divar.ChassisStatus(mostCommon)})
			if err != nil {
				return err
			}

			preProcessedCount++
		}
	}

	fmt.Println("fixed", preProcessedCount, "records rear chassis status attribute")

	return nil
}

func (s *service) InsuranceDue(ctx context.Context) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	carInsuranceDueTypesCount := map[int]int{}

	for _, c := range cars {
		if c.ThirdPartyInsuranceDue != 0 {
			carInsuranceDueTypesCount[c.ThirdPartyInsuranceDue]++
		}
	}

	mostCommon := 1
	for insuranceDue, count := range carInsuranceDueTypesCount {
		if count > carInsuranceDueTypesCount[mostCommon] {
			mostCommon = insuranceDue
		}
	}

	fmt.Println("most common insurance due is", mostCommon, "count =", carInsuranceDueTypesCount[mostCommon])

	preProcessedCount := 0
	for _, c := range cars {
		if c.ThirdPartyInsuranceDue == 0 {
			fmt.Println("car", c.ID, "has wrong insurance due:", c.ThirdPartyInsuranceDue)

			err := s.carService.Update(ctx, c.ID, protocol.Car{ThirdPartyInsuranceDue: mostCommon})
			if err != nil {
				return err
			}

			preProcessedCount++
		}
	}

	fmt.Println("fixed", preProcessedCount, "records insurance due attribute")

	return nil
}

func (s *service) GearBox(ctx context.Context) error {
	cars, err := s.List(ctx)
	if err != nil {
		return err
	}

	carGearboxTypesCount := map[int]int{}

	for _, c := range cars {
		if c.Gearbox != 0 {
			carGearboxTypesCount[int(c.Gearbox)]++
		}
	}

	mostCommon := 1
	for gearbox, count := range carGearboxTypesCount {
		if count > carGearboxTypesCount[mostCommon] {
			mostCommon = gearbox
		}
	}

	fmt.Println("most common gearbox type is", mostCommon, "count =", carGearboxTypesCount[mostCommon])

	preProcessedCount := 0
	for _, c := range cars {
		if c.Gearbox == 0 {
			fmt.Println("car", c.ID, "has wrong gearbox type:", c.Gearbox)

			err := s.carService.Update(ctx, c.ID, protocol.Car{Gearbox: divar.Gearbox(mostCommon)})
			if err != nil {
				return err
			}

			preProcessedCount++
		}
	}

	fmt.Println("fixed", preProcessedCount, "records gearbox type attribute")

	return nil
}
