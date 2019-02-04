package cf

import (
	"testing"
)

func GetMockRatings() Ratings {
	return Ratings{
		// Jack
		{
			User:   "Jack",
			Item:   "Bread",
			Rating: 1.,
		},
		{
			User:   "Jack",
			Item:   "Butter",
			Rating: 1.,
		},
		{
			User:   "Jack",
			Item:   "Milk",
			Rating: 1.,
		},

		// Mary
		{
			User:   "Mary",
			Item:   "Butter",
			Rating: 1.,
		},
		{
			User:   "Mary",
			Item:   "Milk",
			Rating: 1.,
		},
		{
			User:   "Mary",
			Item:   "Beef",
			Rating: 1.,
		},

		// Jane
		{
			User:   "Jane",
			Item:   "Bread",
			Rating: 1.,
		},
		{
			User:   "Jane",
			Item:   "Butter",
			Rating: 1.,
		},

		// Sayani
		{
			User:   "Sayani",
			Item:   "Bread",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Butter",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Milk",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Fish",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Beef",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Ham",
			Rating: 1.,
		},

		// John
		{
			User:   "John",
			Item:   "Fish",
			Rating: 1.,
		},
		{
			User:   "John",
			Item:   "Ham",
			Rating: 1.,
		},

		// Tom
		{
			User:   "Tom",
			Item:   "Fish",
			Rating: 1.,
		},
		{
			User:   "Tom",
			Item:   "Beef",
			Rating: 1.,
		},
		{
			User:   "Tom",
			Item:   "Ham",
			Rating: 1.,
		},

		// Peter
		{
			User:   "Peter",
			Item:   "Butter",
			Rating: 1.,
		},
		{
			User:   "Peter",
			Item:   "Fish",
			Rating: 1.,
		},
		{
			User:   "Peter",
			Item:   "Beef",
			Rating: 1.,
		},

		// OutOfSample
		{
			User:   "OutOfSample",
			Item:   "Eels",
			Rating: 1.,
		},
	}
}

func TestImplicitTopK(t *testing.T) {
	ratings := GetMockRatings()
	cf := NewSimpleGraphCF()
	cf.AddRatings(ratings)
	userTopK, err := cf.UserTopK("John", 3)
	if err != nil {
		t.Fatalf("did not expect to receive an error, got %v", err)
	}
	if len(userTopK) != 3 {
		t.Fatalf("expected 3 results returned for topK, got %d", len(userTopK))
	}
	if userTopK[0].Item != "Beef" {
		t.Errorf("expected the top item to be Beef, got %s", userTopK[0].Item)
	}

	userTopK, err = cf.UserTopK("OutOfSample", 3)
	if len(userTopK) != 0 {
		t.Errorf("did not expect to retrieve any recommendations, got %v", userTopK)
	}
}
