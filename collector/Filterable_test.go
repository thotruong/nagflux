package collector

import (
	"testing"
)

var testTargetFilterData = []struct {
	perfFilter   Filterable
	targetToTest string
	expected     bool
}{
	{Filterable{Filter: All}, All, true},
	{Filterable{Filter: All}, "foo", true},
	{Filterable{Filter: "foo"}, All, true},
	{Filterable{Filter: "foo"}, "foo", true},
	{Filterable{Filter: "bar"}, "foo", false},
	{Filterable{Filter: "foo"}, "bar", false},
	{Filterable{Filter: ""}, "", true},
	{Filterable{Filter: ""}, "foo", false},
	{Filterable{Filter: "foo"}, "", false},
	{Filterable{Filter: "foo,bar"}, "", false},
	{Filterable{Filter: "foo,bar"}, "foo", true},
	{Filterable{Filter: "foo,bar"}, "bar", true},
	{Filterable{Filter: "foo,bar"}, "bar,foo", true},
	{Filterable{Filter: ""}, "foo,bar", false},
	{Filterable{Filter: "foo"}, "foo,bar", true},
	{Filterable{Filter: "bar"}, "foo,bar", true},
	{Filterable{Filter: "bar,foo"}, "foo,bar", true},
}

func TestFilterable_TestTargetFilter(t *testing.T) {
	for i, data := range testTargetFilterData {
		if data.perfFilter.TestTargetFilter(data.targetToTest) != data.expected {
			t.Errorf("%d went wrong. Filterable: %s, ToTest: %s", i, data.perfFilter, data.targetToTest)
		}
	}
}

func TestFilterable_TestTargetFilterObj(t *testing.T) {
	for i, data := range testTargetFilterData {
		if data.perfFilter.TestTargetFilterObj(Filterable{Filter: data.targetToTest}) != data.expected {
			t.Errorf("%d went wrong. Filterable: %s, ToTest: %s", i, data.perfFilter, data.targetToTest)
		}
	}
}
