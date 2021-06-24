package urlValues

import (
	"testing"
)

func getValues() Values {
	values := Values{}
	values.Add("xyz", "abc")
	values.Add("abc", "xyz")
	return values
}

func TestValues_Encode(t *testing.T) {
	values := getValues()

	if values.Encode() != "abc=xyz&xyz=abc" {
		t.FailNow()
	}
}

func TestValues_EncodeWithOrder(t *testing.T) {
	values := getValues()

	if values.EncodeWithOrder() != "xyz=abc&abc=xyz" {
		t.FailNow()
	}
}
