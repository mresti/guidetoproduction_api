package models

import (
	"api/helpers"
	"testing"
)

func TestVisitAPI_marshall(t *testing.T) {
	d := &VisitAPI{
		Visits: 1,
	}
	want := `{"visits":1}`
	helpers.TestJSONMarshal(t, d, want)
}

func TestMessageAPI_marshall(t *testing.T) {
	d := &MessageAPI{
		Message: "Hello",
	}
	want := `{"message":"Hello"}`
	helpers.TestJSONMarshal(t, d, want)
}

func TestNumberAPI_marshall(t *testing.T) {
	d := &NumberAPI{
		Number: 1,
	}
	want := `{"number":1}`
	helpers.TestJSONMarshal(t, d, want)
}
