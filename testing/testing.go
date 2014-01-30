package testing
import (
       "github.com/lab-D8/food-train-server-error"
       "labix.org/v2/mgo"
       "labix.org/v2/mgo/bson"
)

func (u *User) CreateUser(db mgo.Database, email, token) (returnCode int) {
     collection = db.C("users")
     returnCode = error.SUCCESS

     // we still want to check if the user is in the (hopefully clean) database
     // because we don't want to assume anything
     if u.Query(db) != error.SUCCESS {
     	err := collection.Insert(*u)
	if err != nil {
	   returnCode = error.DBERROR
	}
     }

     return returnCode
}

func CleanUp(db mgo.Database) {
     DropDatabase(db);
}

