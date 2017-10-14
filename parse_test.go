package stringargv_test

import (
	"testing"

	. "github.com/cloverstd/parse-string-argv"
)

const (
	cmd1 = `docker run -it --rm     -p    8080:80/tcp       -v /data:/data:ro ubuntu echo "this is a text" 'this is a single quotes text'`
	cmd2 = `docker run -it --rm -p 8080:80/tcp -v /data:/data:ro ubuntu echo "this is a text 'this is a single quotes text'"`
	cmd3 = `docker run -it --rm -p 8080:80/tcp -v /data:/data:ro ubuntu echo "this is a text 'this is a single quotes text'`
	cmd4 = `docker1run1-it11111--rm1-p18080:80/tcp1-v1/data:/data:ro1ubuntu1echo1"this is a text"1'this is a single quotes text'`
)

func TestParseEmpty(t *testing.T) {
	argv, err := Parse("")
	if err != nil {
		t.Error(err)
	} else if len(argv) != 0 {
		t.Errorf("parse string argv length failed: %d %s", len(argv), argv)
	}
}

func TestParseInvalidCMDString(t *testing.T) {
	argv, err := Parse(`""""""''''`)
	if err == nil {
		t.Errorf("parse string argv failed: %d %s", len(argv), argv)
	} else if len(argv) != 0 {
		t.Errorf("parse string argv length failed: %d %s", len(argv), argv)
	}
}
func TestParseSpace(t *testing.T) {
	argv, err := Parse(cmd1)
	if err != nil {
		t.Error(err)
	} else if len(argv) != 12 {
		t.Errorf("parse string argv length failed: %d %s", len(argv), argv)
	} else if argv[11] != "this is a single quotes text" {
		t.Errorf("parse string argv failed: %d - %d", len(argv[11]), len("this is a single quotes text"))
	}
}
func TestParseQuoteInQuote(t *testing.T) {

	argv, err := Parse(cmd2)
	if err != nil {
		t.Error(err)
	} else if len(argv) != 11 {
		t.Errorf("parse string argv length failed: %d %s", len(argv), argv)
	}
}
func TestParseUncloseQuote(t *testing.T) {

	argv, err := Parse(cmd3)
	if err == nil {
		t.Error("parse invalid cmd string failed")
	} else if len(argv) != 0 {
		t.Errorf("parse invalid cmd string failed: %d", len(argv))
	}
}
func TestParseOtherSplit(t *testing.T) {
	argv, err := ParseSplit(cmd4, '1')
	if err != nil {
		t.Error(err)
	} else if len(argv) != 12 {
		t.Errorf("parse string argv length failed: %d %s", len(argv), argv)
	}

}
func TestParseSingle(t *testing.T) {
	argv, err := ParseSplit(cmd1, '1')
	if err != nil {
		t.Error(err)
	} else if len(argv) != 1 {
		t.Errorf("parse string argv length failed: %d %s", len(argv), argv)
	}
}
