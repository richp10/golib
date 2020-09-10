// Copyright (c) 2015 Peter Goldstein -- MIT License
// Additional changes copyright Richard Phillips - MIT License

package dbtype

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"net"
)

// A wrapper for transferring Inet values back and forth easily.
type Inet struct {
	Inet  net.IP
	Valid bool
}

// Copyright (c) 2013 Jack Christensen -- MIT License
// Additional changes copyright Richard Phillips - MIT License
// Allows us to set using a net.IP value or string "127.0.0.1"
func (i *Inet) Set(src interface{}) error {
	if src == nil {
		i.Inet = nil
		i.Valid = false
		return nil
	}

	switch value := src.(type) {
	case net.IP:
		i.Inet = value // already correct type..
		i.Valid = true

	case string:
		_, ipnet, err := net.ParseCIDR(value)
		if err != nil {
			return err
		}

		i.Inet = ipnet.IP
		i.Valid = true

	default:
		mess := fmt.Sprintf("cannot convert %v to Inet", value)
		return errors.New(mess)
	}

	return nil
}

// Scan implements the Scanner interface.
func (i *Inet) Scan(value interface{}) error {
	i.Inet = nil
	i.Valid = false
	if value == nil {
		i.Inet = nil
		i.Valid = false
		return nil
	}
	ipAsBytes, ok := value.([]byte)
	if !ok {
		return errors.New("could not convert scanned value to bytes")
	}
	parsedIP := net.ParseIP(string(ipAsBytes))
	if parsedIP == nil {
		i.Inet = nil
		i.Valid = false
		return nil
	}
	i.Valid = true
	i.Inet = parsedIP
	return nil
}

// Value implements the driver Valuer interface. Note if i.Valid is false
// or i.IP is nil the database column value will be set to NULL.
func (i Inet) Value() (driver.Value, error) {
	if !i.Valid || i.Inet == nil {
		return nil, nil
	}
	return i.Inet.String(), nil
}
