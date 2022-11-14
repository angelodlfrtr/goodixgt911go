package gt911

// GetTouchLevel value
func (gt911 *GT911) GetScreenTouchLevel() (int, error) {
	res, err := gt911.i2cRead(RegScreenTouchLevel, 1)
	if err != nil {
		return 0, err
	}

	return int(res[0]), nil
}

// SetScreenTouchLevel value. Value seems to go from 0 to ~130
func (gt911 *GT911) SetScreenTouchLevel(lvl int) error {
	return gt911.i2cWrite(RegScreenTouchLevel, []byte{byte(lvl)})
}
