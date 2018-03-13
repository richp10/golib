// Copyright (c) 2013 Jack Christensen -- MIT License
// Additional changes copyright Richard Phillips - MIT License
// Stripped down inet data type for postgres

package inet

import (
	"net"
	"errors"

	"reflect"
	"fmt"
	"database/sql/driver"
)

type Status byte

const (
	Undefined Status = iota
	Null
	Present
)

// Inet represents both inet and cidr PostgreSQL types.
type Inet struct {
	IPNet  *net.IPNet
	Status Status
}

func (m *Inet) Set(src interface{}) error {
	if src == nil {
		*m = Inet{Status: Null}
		return nil
	}

	switch value := src.(type) {
	case net.IPNet:
		*m = Inet{IPNet: &value, Status: Present}
	case *net.IPNet:
		*m = Inet{IPNet: value, Status: Present}
	case net.IP:
		bitCount := len(value) * 8
		mask := net.CIDRMask(bitCount, bitCount)
		*m = Inet{IPNet: &net.IPNet{Mask: mask, IP: value}, Status: Present}
	case string:
		_, ipnet, err := net.ParseCIDR(value)
		if err != nil {
			return err
		}
		*m = Inet{IPNet: ipnet, Status: Present}
	default:
		if originalSrc, ok := underlyingPtrType(src); ok {
			return m.Set(originalSrc)
		}
		mess := fmt.Sprintf("cannot convert %v to Inet", value)
		return errors.New(mess)
	}

	return nil
}

func (m *Inet) Get() interface{} {
	switch m.Status {
	case Present:
		return m.IPNet
	case Null:
		return nil
	default:
		return m.Status
	}
}

// Scan implements the database/sql Scanner interface.
func (m *Inet) Scan(src interface{}) error {
	if src == nil {
		*m = Inet{Status: Null}
		return nil
	}

	switch src := src.(type) {
	case string:
		return m.decodeText([]byte(src))
	case []byte:
		srcCopy := make([]byte, len(src))
		copy(srcCopy, src)
		return m.decodeText(srcCopy)
	}

	return errors.New("cannot scan: " + src.(string))
}

// Value implements the database/sql/driver Valuer interface.
func (m *Inet) Value() (driver.Value, error) {
	switch m.Status {
	case Present:
		return m.IPNet, nil
	case Null:
		return nil, nil
	default:
		return m.Status, nil
	}
}

func (m *Inet) decodeText(src []byte) error {
	if src == nil {
		*m = Inet{Status: Null}
		return nil
	}

	var ipnet *net.IPNet
	var err error

	if ip := net.ParseIP(string(src)); ip != nil {
		ipv4 := ip.To4()
		if ipv4 != nil {
			ip = ipv4
		}
		bitCount := len(ip) * 8
		mask := net.CIDRMask(bitCount, bitCount)
		ipnet = &net.IPNet{Mask: mask, IP: ip}
	} else {
		_, ipnet, err = net.ParseCIDR(string(src))
		if err != nil {
			return err
		}
	}

	*m = Inet{IPNet: ipnet, Status: Present}
	return nil
}

// Put in common package if use more pgx types
// underlyingPtrType dereferences a pointer

func underlyingPtrType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	}

	return nil, false
}
