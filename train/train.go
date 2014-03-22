package train

import (
	"fmt"
	"github.com/lab-D8/food-train-server/error"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Train struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Creator      string
	Where        string
	When         string
	ActiveHash   string
	Privacy      string
}

// Gets the train info using the id.
func (t *Train) Query(db mgo.Database) (returnCode int){
	result := Train{}
	err := db.C("trains").Find(bson.M{"_id": t.Id}).One(&result)

	if err != nil{
		return error.NOTFOUND
	} else {
		t.Id = result.Id
		t.Creator = result.Creator
		t.Where = result.Where
		t.When = result.When
		t.ActiveHash = result.ActiveHash
		t.Privacy = result.Privacy
		return error.SUCCESS
	}
}

// Create a train.
// Expects the user exists
func (t *Train) Create(db mgo.Database) (returnCode int){
	t.Id = bson.NewObjectId()
	
	err := db.C("trains").Insert(t)

	if err != nil {
		return error.DBERROR;
	}

	return error.SUCCESS
}
