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

	values.Del("xyz")

	if values.EncodeWithOrder() != "abc=xyz" {
		t.Fatalf("failed to delete item")
	}
}

func TestValues_EncodeWithOrder(t *testing.T) {
	values := getValues()

	if values.EncodeWithOrder() != "xyz=abc&abc=xyz" {
		t.FailNow()
	}

	values[OrderKey] = []string{"abc", "xyz"}

	if values.EncodeWithOrder() != "abc=xyz&xyz=abc" {
		t.Fatalf("failed to manually rearrange order")
	}

	values.Del("abc")

	if values.EncodeWithOrder() != "xyz=abc" && len(values[OrderKey]) != 1 {
		t.Fatalf("failed to fully delete item")
	}
}
