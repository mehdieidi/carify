CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    year INT,
    color INT,
    usage_km INT,
    body_status INT,
    cash_cost BIGINT,
    motor_status SMALLINT,
    front_chassis_status SMALLINT,
    rear_chassis_status SMALLINT,
    third_party_insurance_due INT,
    gearbox SMALLINT,
    car_token VARCHAR(20)
);
