package cf

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"strconv"
)

func RatingsFromCSV(file string) (Ratings, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(data))
	ratings := make(Ratings, 0, 32)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return ratings, err
		}
		if len(record) != 3 {
			return ratings, errors.New("expecting 3 columns")
		}
		rating, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return ratings, err
		}

		ratings = append(ratings, Rating{
			Item:   record[0],
			User:   record[1],
			Rating: rating,
		})
	}

	return ratings, nil
}
