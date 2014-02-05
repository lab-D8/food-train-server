package testing
import (
       "encoding/json"
       "fmt"
       "github.com/lab-D8/food-train-server-error"
       "github.com/lab-D8/food-train-server-user"
       "labix.org/v2/mgo"
       "labix.org/v2/mgo/bson"
       "net/http"
)

type Response map[string]interface{}
var host="localhost"
var testdb_name = "catan-testing" 

func HandlerGenerator(handle func(*HttpRequest, mgo.Database) (Response)) (handler func(http.ResponseWriter, *http.Request)){
     return func(w http.ResponseWriter, r *http.Request) {
          r.ParseForm()
	  w.Header().Set("Content-Type", "application/json")
	  session, err := mgo.Dial(host)
	  if err != nil {
	       fmt.Fprint(w, Response{"return_code": error.DBERROR, "description": error.GetDescription(error.DBERROR)})
	       return
	  }
	  defer session.Close()
	  session.SetMode(mgo.Monotonic, true)
	  db := session.DB(testdb_name)
	  response := handle(r, *db)
	  fmt.Fprint(w, response)
     }
}

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

//cleans up the databsae so we have a shiny new one for the next test
//right now we only have users to worry about
//note that collections are created automatically on use so they don't need to be manually recreated
func CleanUp(db mgo.Database) (returnCode int){
     returnCode = error.SUCCESS
     err := DropCollection("users");
     if err != nil { returnCode = error.DBERROR }
     return returnCode
}

