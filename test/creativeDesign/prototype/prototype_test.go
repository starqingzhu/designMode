package test

import (
	"designmode/creativeDesign/prototype"
	"fmt"
	"testing"
	"time"
)

func TestProtoType(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	words := prototype.KeyWords{
		"testA": &prototype.KeyWord{
			Word:      "testA",
			Visit:     1,
			UpdatedAt: &updateAt,
		},
		"testB": &prototype.KeyWord{
			Word:      "testB",
			Visit:     2,
			UpdatedAt: &updateAt,
		},
		"testC": &prototype.KeyWord{
			Word:      "testC",
			Visit:     3,
			UpdatedAt: &updateAt,
		},
	}

	now := time.Now()
	updatedWords := []*prototype.KeyWord{
		{
			Word:      "testB",
			Visit:     10,
			UpdatedAt: &now,
		},
	}

	got := words.Clone(updatedWords)
	fmt.Println("testA:", words["testA"], got["testA"])
	fmt.Println("testB:", words["testB"], got["testB"])
	fmt.Println("testC:", words["testC"], got["testC"])
	// assert.Equal(t, words["testA"], got["testA"])
	// assert.NotEqual(t, words["testB"], got["testB"])
	// assert.NotEqual(t, updatedWords[0], got["testB"])
	// assert.Equal(t, words["testC"], got["testC"])

	fmt.Printf("ProtoType end ...\n")
}
