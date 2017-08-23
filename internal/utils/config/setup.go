package config

import (
	"errors"

	"github.com/namsral/flag"
)

var (
	Port int

	//Additional config vars declaration goes here
	//TODO
)

//Initialize set the values of all config vars and check of availability of
// all required parameters
func Initialize() error {
	flag.IntVar(&Port, "port", 0, "Port to run on")

	//Set config value here. Set default value to zero value
	// if parameter is required
	//TODO

	flag.Parse()

	return IsAllRequiredConfigAvailable()
}

//IsAllRequiredConfigAvailable check whether or not all required
// Environment variables are available. It return an error if any
// env var is missing. Value of config parameter if checked against
// it's zero values.
func IsAllRequiredConfigAvailable() error {
	if Port == 0 {
		return errors.New("please set port value")
	}
	//TODO

	return nil
}
