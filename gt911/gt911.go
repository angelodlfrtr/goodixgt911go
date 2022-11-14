package gt911

import (
	"encoding/binary"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	host "periph.io/x/host/v3"
)

// GT911 represent a goodix gt911 touch sensor device
type GT911 struct {
	bus i2c.BusCloser
	dev *i2c.Dev
}

func New(i2cBusID string, deviceAddr uint16) (*GT911, error) {
	// Init i2c lib
	if _, err := host.Init(); err != nil {
		return nil, err
	}
	if _, err := driverreg.Init(); err != nil {
		return nil, err
	}

	// Open bus
	bus, err := i2creg.Open(i2cBusID)
	if err != nil {
		return nil, err
	}

	// Open device
	dev := &i2c.Dev{Addr: deviceAddr, Bus: bus}

	return &GT911{bus, dev}, nil
}

// Apply configurations
func (gt911 *GT911) Apply() error {
	return gt911.configAfterWrite()
}

// Close device
func (gt911 *GT911) Close() error {
	return gt911.bus.Close()
}

// i2cRead on device
func (gt911 *GT911) i2cRead(reg uint16, len int) ([]byte, error) {
	write := make([]byte, 2)
	read := make([]byte, len)
	binary.BigEndian.PutUint16(write, uint16(reg))
	if err := gt911.dev.Tx(write, nil); err != nil {
		return nil, err
	}

	if err := gt911.dev.Tx(nil, read); err != nil {
		return nil, err
	}

	return read, nil
}

// i2cWrite on device
func (gt911 *GT911) i2cWrite(reg uint16, data []byte) error {
	write := make([]byte, 2)
	binary.BigEndian.PutUint16(write, uint16(reg))
	write = append(write, data...)
	return gt911.dev.Tx(write, nil)
}
