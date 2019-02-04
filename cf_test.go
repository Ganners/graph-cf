package cf

import "testing"

func TestImplicitTopK(t *testing.T) {
	ratings := Ratings{
		// Jack
		{
			User:   "Jack",
			Item:   "Bread",
			Rating: 1,
		},
		{
			User:   "Jack",
			Item:   "Butter",
			Rating: 1,
		},
		{
			User:   "Jack",
			Item:   "Milk",
			Rating: 1,
		},
		{
			User:   "Jack",
			Item:   "Fish",
			Rating: 0,
		},
		{
			User:   "Jack",
			Item:   "Beef",
			Rating: 0,
		},
		{
			User:   "Jack",
			Item:   "Ham",
			Rating: 0,
		},

		// Mary
		{
			User:   "Mary",
			Item:   "Bread",
			Rating: 0,
		},
		{
			User:   "Mary",
			Item:   "Butter",
			Rating: 1,
		},
		{
			User:   "Mary",
			Item:   "Milk",
			Rating: 1,
		},
		{
			User:   "Mary",
			Item:   "Fish",
			Rating: 0,
		},
		{
			User:   "Mary",
			Item:   "Beef",
			Rating: 1,
		},
		{
			User:   "Mary",
			Item:   "Ham",
			Rating: 0,
		},

		// Jane
		{
			User:   "Jane",
			Item:   "Bread",
			Rating: 1,
		},
		{
			User:   "Jane",
			Item:   "Butter",
			Rating: 1,
		},
		{
			User:   "Jane",
			Item:   "Milk",
			Rating: 0,
		},
		{
			User:   "Jane",
			Item:   "Fish",
			Rating: 0,
		},
		{
			User:   "Jane",
			Item:   "Beef",
			Rating: 0,
		},
		{
			User:   "Jane",
			Item:   "Ham",
			Rating: 0,
		},

		// Sayani
		{
			User:   "Sayani",
			Item:   "Bread",
			Rating: 1,
		},
		{
			User:   "Sayani",
			Item:   "Butter",
			Rating: 1,
		},
		{
			User:   "Sayani",
			Item:   "Milk",
			Rating: 1,
		},
		{
			User:   "Sayani",
			Item:   "Fish",
			Rating: 1,
		},
		{
			User:   "Sayani",
			Item:   "Beef",
			Rating: 1,
		},
		{
			User:   "Sayani",
			Item:   "Ham",
			Rating: 1,
		},

		// John
		{
			User:   "John",
			Item:   "Bread",
			Rating: 0,
		},
		{
			User:   "John",
			Item:   "Butter",
			Rating: 0,
		},
		{
			User:   "John",
			Item:   "Milk",
			Rating: 0,
		},
		{
			User:   "John",
			Item:   "Fish",
			Rating: 1,
		},
		{
			User:   "John",
			Item:   "Beef",
			Rating: 0,
		},
		{
			User:   "John",
			Item:   "Ham",
			Rating: 1,
		},

		// Tom
		{
			User:   "Tom",
			Item:   "Bread",
			Rating: 0,
		},
		{
			User:   "Tom",
			Item:   "Butter",
			Rating: 0,
		},
		{
			User:   "Tom",
			Item:   "Milk",
			Rating: 0,
		},
		{
			User:   "Tom",
			Item:   "Fish",
			Rating: 1,
		},
		{
			User:   "Tom",
			Item:   "Beef",
			Rating: 1,
		},
		{
			User:   "Tom",
			Item:   "Ham",
			Rating: 1,
		},

		// Peter
		{
			User:   "Peter",
			Item:   "Bread",
			Rating: 0,
		},
		{
			User:   "Peter",
			Item:   "Butter",
			Rating: 1,
		},
		{
			User:   "Peter",
			Item:   "Milk",
			Rating: 0,
		},
		{
			User:   "Peter",
			Item:   "Fish",
			Rating: 1,
		},
		{
			User:   "Peter",
			Item:   "Beef",
			Rating: 1,
		},
		{
			User:   "Peter",
			Item:   "Ham",
			Rating: 0,
		},
	}

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
}
