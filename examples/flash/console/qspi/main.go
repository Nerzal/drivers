package main

import (
	"machine"

	"github.com/Nerzal/drivers/examples/flash/console"
	"github.com/Nerzal/drivers/flash"
)

func main() {
	console_example.RunFor(
		flash.NewQSPI(
			machine.QSPI_CS,
			machine.QSPI_SCK,
			machine.QSPI_DATA0,
			machine.QSPI_DATA1,
			machine.QSPI_DATA2,
			machine.QSPI_DATA3,
		),
	)
}
