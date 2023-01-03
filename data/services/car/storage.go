package car

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mehdieidi/carify/data/pkg/log"
	"github.com/mehdieidi/carify/data/pkg/xsql"
	"github.com/mehdieidi/carify/data/protocol"
	"github.com/mehdieidi/carify/data/protocol/derror"
)

const domain = "car"

type storage struct {
	db     xsql.TableOperator
	logger log.Logger
}

func NewStorage(db xsql.TableOperator, logger log.Logger) protocol.CarStorage {
	return &storage{
		db:     db,
		logger: logger,
	}
}

func (s *storage) Store(ctx context.Context, car protocol.Car) (protocol.CarID, error) {
	const query = `
		INSERT INTO cars (
			year,
			color,
			usage_km,
			body_status,
			cash_cost,
			motor_status,
			front_chassis_status,
			rear_chassis_status,
			third_party_insurance_due,
			gearbox,
			car_token
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		) RETURNING id
	`

	row := s.db.QueryRowContext(ctx, query,
		car.Year,
		car.Color,
		car.UsageKM,
		car.BodyStatus,
		car.CashCost,
		car.MotorStatus,
		car.FrontChassisStatus,
		car.RearChassisStatus,
		car.ThirdPartyInsuranceDue,
		car.Gearbox,
		car.Token,
	)

	err := row.Scan(&car.ID)
	if err != nil {
		s.logger.Error(domain, log.StorageLayer, "Store", log.Args{log.LogErrKey: err})
		return 0, derror.ErrUnexpected
	}

	return car.ID, nil
}

func (s *storage) FindByToken(ctx context.Context, token protocol.CarToken) (protocol.Car, error) {
	const query = `
		SELECT
			id,
			year,
			color,
			usage_km,
			body_status,
			cash_cost,
			motor_status,
			front_chassis_status,
			rear_chassis_status,
			third_party_insurance_due,
			gearbox,
			car_token
		FROM cars 
		WHERE car_token = $1 
		LIMIT 1
	`

	row := s.db.QueryRowContext(ctx, query, token)

	var c protocol.Car
	err := row.Scan(
		&c.ID,
		&c.Year,
		&c.Color,
		&c.UsageKM,
		&c.BodyStatus,
		&c.CashCost,
		&c.MotorStatus,
		&c.FrontChassisStatus,
		&c.RearChassisStatus,
		&c.ThirdPartyInsuranceDue,
		&c.Gearbox,
		&c.Token,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return protocol.Car{}, derror.ErrUnknownCar
		}

		s.logger.Error(domain, log.StorageLayer, "FindByToken", log.Args{log.LogErrKey: err})

		return protocol.Car{}, derror.ErrUnexpected
	}

	return c, nil
}

func (s *storage) FindByID(ctx context.Context, id protocol.CarID) (protocol.Car, error) {
	const query = `
		SELECT
			id,
			year,
			color,
			usage_km,
			body_status,
			cash_cost,
			motor_status,
			front_chassis_status,
			rear_chassis_status,
			third_party_insurance_due,
			gearbox,
			car_token
		FROM cars 
		WHERE id = $1 
		LIMIT 1
	`

	row := s.db.QueryRowContext(ctx, query, id)

	var c protocol.Car
	err := row.Scan(
		&c.ID,
		&c.Year,
		&c.Color,
		&c.UsageKM,
		&c.BodyStatus,
		&c.CashCost,
		&c.MotorStatus,
		&c.FrontChassisStatus,
		&c.RearChassisStatus,
		&c.ThirdPartyInsuranceDue,
		&c.Gearbox,
		&c.Token,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return protocol.Car{}, derror.ErrUnknownCar
		}

		s.logger.Error(domain, log.StorageLayer, "FindByID", log.Args{log.LogErrKey: err})

		return protocol.Car{}, derror.ErrUnexpected
	}

	return c, nil
}

func (s *storage) Update(ctx context.Context, id protocol.CarID, car protocol.Car) error {
	const query = `
		UPDATE cars
		SET 
			year = $2,
			color = $3,
			usage_km = $4,
			body_status = $5,
			cash_cost = $6,
			motor_status = $7,
			front_chassis_status = $8,
			rear_chassis_status = $9,
			third_party_insurance_due = $10,
			gearbox = $11,
			car_token = $12
		WHERE id = $1
	`

	_, err := s.db.Exec(query, car.ID,
		car.Year,
		car.Color,
		car.UsageKM,
		car.BodyStatus,
		car.CashCost,
		car.MotorStatus,
		car.FrontChassisStatus,
		car.RearChassisStatus,
		car.ThirdPartyInsuranceDue,
		car.Gearbox,
		car.Token,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) Delete(ctx context.Context, id protocol.CarID) error {
	const query = `DELETE FROM cars WHERE id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
