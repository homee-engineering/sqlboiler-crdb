package main

import (
	"github.com/homee-engineering/sqlboiler-crdb/driver"
	"github.com/volatiletech/sqlboiler/drivers"
)

func main() {
	drivers.DriverMain(&driver.CockroachDBDriver{})
}
