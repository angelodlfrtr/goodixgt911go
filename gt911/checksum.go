package gt911

func (gt911 *GT911) configAfterWrite() error {
	vals, err := gt911.readConfigValues()
	if err != nil {
		return err
	}

	check_sum := gt911.calcConfigValuesChecksum(vals)

	// Write checksum
	if err := gt911.i2cWrite(RegConfigChksum, []byte{check_sum}); err != nil {
		return err
	}

	// Write config frech
	if err := gt911.i2cWrite(RegConfigFresh, []byte{1}); err != nil {
		return err
	}

	return nil
}

// readConfigValues from registers (checksum value of the bytes from 0x8047 to 0x80FE)
func (gt911 *GT911) readConfigValues() ([]byte, error) {
	vals := []byte{}

	for i := 0x8047; i <= 0x80FE; i++ {
		val, err := gt911.i2cRead(uint16(i), 1)
		if err != nil {
			return nil, err
		}

		vals = append(vals, val[0])
	}

	return vals, nil
}

// calcConfigValuesChecksum from https://github.com/torvalds/linux/blob/master/drivers/input/touchscreen/goodix.c#L543
func (gt911 *GT911) calcConfigValuesChecksum(vals []byte) byte {
	var check_sum byte = 0

	for _, v := range vals {
		check_sum += v
	}

	check_sum = (^check_sum) + 1

	return check_sum
}
