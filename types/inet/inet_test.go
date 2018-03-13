// Copyright (c) 2013 Jack Christensen -- MIT License
// Additional changes copyright Richard Phillips - MIT License

package inet_test

import (
	"net"
	"reflect"
	"testing"

	"github.com/richp10/golib/types/inet"
)

// Todo re-expand these tests - see pgx..
func TestInetSet(t *testing.T) {
	successfulTests := []struct {
		source interface{}
		result inet.Inet
	}{
		{source: mustParseCIDR(t, "127.0.0.1/32"), result: inet.Inet{IPNet: mustParseCIDR(t, "127.0.0.1/32"), Status: inet.Present}},
		{source: mustParseCIDR(t, "127.0.0.1/32").IP, result: inet.Inet{IPNet: mustParseCIDR(t, "127.0.0.1/32"), Status: inet.Present}},
		{source: "127.0.0.1/32", result: inet.Inet{IPNet: mustParseCIDR(t, "127.0.0.1/32"), Status: inet.Present}},
	}

	for i, tt := range successfulTests {
		var r inet.Inet
		err := r.Set(tt.source)
		if err != nil {
			t.Errorf("%d: %v", i, err)
		}

		if !reflect.DeepEqual(r, tt.result) {
			t.Errorf("%d: expected %v to convert to %v, but it was %v", i, tt.source, tt.result, r)
		}
	}
}

func mustParseCIDR(t testing.TB, s string) *net.IPNet {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		t.Fatal(err)
	}

	return ipnet
}
