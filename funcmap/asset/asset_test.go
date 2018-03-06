// Copyright (c) 2016 Blue Jay - MIT License
// Additional changes copyright Richard Phillips - MIT License

package asset

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMegaChecks(t *testing.T) {

	Convey("Should pass megachecks", t, func() {

		cmd := exec.Command("megacheck")
		res, _ := cmd.Output()
		So(string(res[:]), ShouldBeEmpty)

	})
}

func init() {
	// Clear log messages from output (specifically TestCSSMissing())
	log.SetOutput(ioutil.Discard)
}

// TestKeysExist ensures keys exist.
func TestKeysExist(t *testing.T) {
	Convey("Should ensure key exists", t, func() {

		config := Info{
			Folder: "testdata",
		}

		fm := config.Map("/")

		js := false
		if _, ok := fm["JS"]; ok {
			js = true
		}
		So(js, ShouldBeTrue)

		css := false
		if _, ok := fm["CSS"]; ok {
			css = true
		}
		So(css, ShouldBeTrue)

	})
}

// TestCSS ensures CSS parses correctly.
func TestCSS(t *testing.T) {
	Convey("Should ensures CSS parses correctly", t, func() {

		config := Info{
			Folder: "testdata",
		}

		fm := config.Map("/")

		temp, err := template.New("test").Funcs(fm).Parse(`{{CSS "/test.css" "all"}}`)
		So(err, ShouldBeNil)

		buf := new(bytes.Buffer)

		err = temp.Execute(buf, nil)
		So(err, ShouldBeNil)

		expected := `<link media="all" rel="stylesheet" type="text/css" href="/test.css?`
		received := buf.String()

		gotit := strings.HasPrefix(received, expected)
		So(gotit, ShouldBeTrue)
	})
}

// TestCSS ensures CSS from internet parses correctly.
func TestCSSInternet(t *testing.T) {
	Convey("Should ensures CSS from internet parses correctly.", t, func() {

		config := Info{
			Folder: "testdata",
		}

		fm := config.Map("/")

		temp, err := template.New("test").Funcs(fm).Parse(`{{CSS "//test.css" "all"}}`)
		So(err, ShouldBeNil)

		buf := new(bytes.Buffer)

		err = temp.Execute(buf, nil)
		So(err, ShouldBeNil)

		expected := `<link media="all" rel="stylesheet" type="text/css" href="//test.css" />`
		received := buf.String()

		So(expected, ShouldEqual, received)

	})
}

// TestCSSMissing ensures file is missing error is thrown.
func TestCSSMissing(t *testing.T) {
	Convey("Should ensures file is missing error is thrown..", t, func() {

		config := Info{
			Folder: "testdata2",
		}

		fm := config.Map("/")

		temp, err := template.New("test").Funcs(fm).Parse(`{{CSS "test.css" "all"}}`)
		So(err, ShouldBeNil)

		buf := new(bytes.Buffer)

		err = temp.Execute(buf, nil)
		So(err, ShouldBeNil)

		expected := `<!-- CSS Error: test.css -->`
		received := buf.String()

		So(expected, ShouldEqual, received)

	})
}

// TestJS ensures JS parses correctly.
func TestJS(t *testing.T) {
	Convey("Should ensures file is missing error is thrown..", t, func() {

		config := Info{
			Folder: "testdata",
		}

		fm := config.Map("/")

		temp, err := template.New("test").Funcs(fm).Parse(`{{JS "test.js"}}`)
		So(err, ShouldBeNil)

		buf := new(bytes.Buffer)

		err = temp.Execute(buf, nil)
		So(err, ShouldBeNil)

		expected := `<script type="text/javascript" src="/test.js?`
		received := buf.String()

		gotit := strings.HasPrefix(received, expected)
		So(gotit, ShouldBeTrue)

	})
}

// TestJS ensures JS from internet parses correctly.
func TestJSInternet(t *testing.T) {
	config := Info{
		Folder: "testdata",
	}

	fm := config.Map("/")

	temp, err := template.New("test").Funcs(fm).Parse(`{{JS "//test.js"}}`)
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = temp.Execute(buf, nil)
	if err != nil {
		t.Fatal(err)
	}

	expected := `<script type="text/javascript" src="//test.js"></script>`
	received := buf.String()

	if expected != received {
		t.Errorf("\n got: %v\nwant: %v", received, expected)
	}
}

// TestJSMissing ensures file is missing error is thrown.
func TestJSMissing(t *testing.T) {
	config := Info{
		Folder: "testdata2",
	}

	fm := config.Map("/")

	temp, err := template.New("test").Funcs(fm).Parse(`{{JS "test2.js"}}`)
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = temp.Execute(buf, nil)
	if err != nil {
		t.Fatal(err)
	}

	expected := `<!-- JS Error: test2.js -->`
	received := buf.String()

	if expected != received {
		t.Errorf("\n got: %v\nwant: %v", received, expected)
	}
}

// TestIMG ensures IMG parses correctly.
func TestIMG(t *testing.T) {
	config := Info{
		Folder: "testdata",
	}

	fm := config.Map("/")

	temp, err := template.New("test").Funcs(fm).Parse(`{{IMG "test.gif" "Loader"}}`)
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = temp.Execute(buf, nil)
	if err != nil {
		t.Fatal(err)
	}

	expected := `<img src="/test.gif?`
	received := buf.String()

	if !strings.HasPrefix(received, expected) {
		t.Errorf("\n got: %v\nwant: %v", received, expected)
	}
}

// TestIMG ensures IMG from internet parses correctly.
func TestIMGInternet(t *testing.T) {
	config := Info{
		Folder: "testdata",
	}

	fm := config.Map("/")

	temp, err := template.New("test").Funcs(fm).Parse(`{{IMG "//test.gif" "Loader"}}`)
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = temp.Execute(buf, nil)
	if err != nil {
		t.Fatal(err)
	}

	expected := `<img src="//test.gif" alt="Loader" />`
	received := buf.String()

	if expected != received {
		t.Errorf("\n got: %v\nwant: %v", received, expected)
	}
}

// TestIMGMissing ensures file is missing error is thrown.
func TestIMGMissing(t *testing.T) {
	config := Info{
		Folder: "testdata2",
	}

	fm := config.Map("/")

	temp, err := template.New("test").Funcs(fm).Parse(`{{IMG "test2.gif" "Loader"}}`)
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = temp.Execute(buf, nil)
	if err != nil {
		t.Fatal(err)
	}

	expected := `<!-- IMG Error: test2.gif -->`
	received := buf.String()

	if expected != received {
		t.Errorf("\n got: %v\nwant: %v", received, expected)
	}
}

// TTestBaseURI ensure URI is handled correctly.
func TestBaseURI(t *testing.T) {
	config := Info{
		Folder: "testdata",
	}

	fm := config.Map("/newbase/")

	temp, err := template.New("test").Funcs(fm).Parse(`{{CSS "/test.css" "all"}}`)
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = temp.Execute(buf, nil)
	if err != nil {
		t.Fatal(err)
	}

	expected := `<link media="all" rel="stylesheet" type="text/css" href="/newbase/test.css?`
	received := buf.String()

	if !strings.HasPrefix(received, expected) {
		t.Errorf("\n got: %v\nwant: %v", received, expected)
	}
}
