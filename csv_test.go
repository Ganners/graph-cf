package cf

import (
	"reflect"
	"testing"
)

func TestRatingsFromCSV(t *testing.T) {
	fixture := "./data/fixture.csv"
	expectedRatings := GetMockRatings()

	ratings, err := RatingsFromCSV(fixture)
	if err != nil {
		t.Errorf("did not expect to see an error, got %v", err)
	}
	if !reflect.DeepEqual(ratings, expectedRatings) {
		t.Errorf("expected %#v got %#v", ratings, expectedRatings)
	}
}
