package preprocess

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mehdieidi/carify/data/pkg/log"
	"github.com/mehdieidi/carify/data/pkg/xsql"
	"github.com/mehdieidi/carify/data/protocol"
	"github.com/mehdieidi/carify/data/protocol/derror"
)

const domain = "preprocess"

type storage struct {
	db     xsql.TableOperator
	logger log.Logger
}

func NewStorage(db xsql.TableOperator, logger log.Logger) protocol.PreProcessStorage {
	return &storage{
		db:     db,
		logger: logger,
	}
}

func (s *storage) FindAll(ctx context.Context) ([]protocol.Car, error) {
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
		ORDER BY id
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		s.logger.Error(domain, log.StorageLayer, "FindAll", log.Args{log.LogErrKey: err})
		return nil, derror.ErrUnexpected
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			s.logger.Error(domain, log.StorageLayer, "FindAll", log.Args{log.LogErrKey: err})
		}
	}(rows)

	cars := []protocol.Car{}
	for rows.Next() {
		var c protocol.Car
		if err := rows.Scan(
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
		); err != nil {
			s.logger.Error(domain, log.StorageLayer, "FindAll", log.Args{log.LogErrKey: err})
			if errors.Is(err, sql.ErrNoRows) {
				return nil, derror.ErrUnknownCar
			}
			return nil, derror.ErrUnexpected
		}

		cars = append(cars, c)
	}

	if err := rows.Err(); err != nil {
		s.logger.Error(domain, log.StorageLayer, "FindAll", log.Args{log.LogErrKey: err})
		return nil, derror.ErrUnexpected
	}

	return cars, nil
}
