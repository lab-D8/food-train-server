package main

import (
	"encoding/json"
	"fmt"
	"github.com/iph/catan/error"
//	"github.com/iph/catan/token"
	"github.com/iph/catan/user"
	"labix.org/v2/mgo"
	"net/http"
)

type Response map[string]interface{}

var host = "localhost"
var database_name = "catan"

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

func main() {
	http.HandleFunc("/friends/add", HandlerGenerator(AddFriendFunc))
	http.HandleFunc("/friends/remove", HandlerGenerator(RemoveFriendFunc))
	http.ListenAndServe(":8080", nil)
}

func HandlerGenerator(handle func(*http.Request, mgo.Database) (Response)) (handler func(http.ResponseWriter, *http.Request)){
	return func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()
		w.Header().Set("Content-Type", "application/json")
		session, err := mgo.Dial(host)
		if err != nil {
			fmt.Fprint(w, Response{"return_code": error.DBERROR, "description": error.GetDescription(error.DBERROR)})
			return
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		db := session.DB(database_name)		
		
		response := handle(r, *db)
		fmt.Fprint(w, response)
		
	}
}



/************************************
 * See user.RemoveFriend for return codes
 ************************************/
func RemoveFriendFunc(r *http.Request, db mgo.Database) (Response) {
	var ret_err = error.SUCCESS
	var user_email = r.FormValue("user")
	var friend_email = r.FormValue("friend")

	first_user := user.User{user_email, "", "", "", 0, 0, token.Token{}}
	ret_err = first_user.RemoveFriend(db, friend_email)

	// return the error code from RemoveFriend, or NOTFOUND if one or both users don't exist
	return Response{"return_code": ret_err, "description": error.GetDescription(ret_err)}
}

// AddFriendFunc:
// 		Largely stylized after NewUserFunc below.
//    Requires a {user, friend}.
// RETURNS:
//		SUCCESS if both users exist the friendship is a new one
//		NOTFOUND if one or both users do not exist in the DB
//		ALREADY if the friendship already exists
//		DBERROR if something bad with DB happened
func AddFriendFunc(r *http.Request, db mgo.Database) (Response) {
	var ret_err = error.SUCCESS
	var user_email = r.FormValue("user")
	var friend_email = r.FormValue("friend")

	first_user := user.User{user_email, "", "", "", 0, 0, token.Token{}}
	ret_err = first_user.AddFriend(db, friend_email)

	// return the error code from AddFriend, or NOTFOUND if one or both users don't exist
	return Response{"return_code": ret_err, "description": error.GetDescription(ret_err)}
}

// NewUserFunc:
//    Requires a {first_name, last_name, password, email}.
//    Writes a json encoded value back to the user of success, or error.

