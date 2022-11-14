package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/angelodlfrtr/goodixgt911go/gt911"
)

func Main() {
	args := os.Args

	if strings.Contains(args[1], "help") {
		usage()
		os.Exit(0)
	}

	if len(args) < 3 {
		usage()
		os.Exit(1)
	}

	// Get params
	busName := args[1]
	cleanedDeviceAddr := strings.ReplaceAll(args[2], "0x", "")
	deviceAddr, err := strconv.ParseUint(cleanedDeviceAddr, 16, 8)
	if err != nil {
		panic(err)
	}

	if args[3] == "set" {
		if len(args) < 4 {
			usage()
			os.Exit(1)
		}
	}

	// Init device
	dev, err := gt911.New(busName, uint16(deviceAddr))
	if err != nil {
		panic(err)
	}
	defer dev.Close()
	defer func() {
		if r := recover(); r != nil {
			dev.Close()
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	if args[3] == "set" {
		val, err := strconv.Atoi(args[4])
		if err != nil {
			panic(err)
		}

		if err := dev.SetScreenTouchLevel(val); err != nil {
			panic(err)
		}

		if err := dev.Apply(); err != nil {
			panic(err)
		}

		return
	}

	// Get
	val, err := dev.GetScreenTouchLevel()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}

func usage() {
	fmt.Printf("Usage: %s [i2cbus] [device_addr_hex] (set|get) [value]\n", os.Args[0])
}
