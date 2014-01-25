package main

import (
	"fmt"
	"github.com/iph/catan/error"
	"github.com/iph/catan/token"
	"github.com/iph/catan/user"
	"labix.org/v2/mgo"
	"testing"
)

func TestMain(t *testing.T) {
	db := makeDB("localhost", "test")
	var newUser = map[string]string{
		"email":    "seanmyers0608@gmail.com",
		"password": "herpaderp",
		"fname":    "Sean",
		"lname":    "Myers",
	}

	user1 := user.User{newUser["email"], newUser["password"], newUser["fname"],
		newUser["lname"], 0, 0, token.Token{}}
	

	fmt.Printf("Catan!\n")

	fmt.Printf("Adding user to DB: ")
	if user1.New(db) == error.ALREADY {
		fmt.Printf("%s\n", error.GetDescription(error.ALREADY))
	}

	user1.Email = "seanmyers0608@gmail.com"

	var newUser2 = map[string]string{
		"email":    "zaxcoding@gmail.com",
		"password": "hearthstone",
		"fname":    "Zach",
		"lname":    "Sadler",
	}

	user2 := user.User{newUser2["email"], newUser2["password"], newUser2["fname"],
		newUser2["lname"], 0, 0, token.Token{}}

	// Test New
	user2Err := user2.New(db)
	fmt.Printf("Adding second user to DB: ")
	if user2Err == error.ALREADY {
		fmt.Printf("%s\n", error.GetDescription(error.ALREADY))
	} else if user2Err == error.SUCCESS {
		fmt.Printf("%s\n", error.GetDescription(error.SUCCESS))
	}

	// Test AddFriend
	friendshipErr := user1.AddFriend(db, user2.Email)
	if friendshipErr == error.SUCCESS {
		fmt.Printf("Friendship added: ")
	} else if friendshipErr == error.NOTFOUND {
		fmt.Printf("One of the users is not in DB: ")
	} else if friendshipErr == error.ALREADY {
		fmt.Printf("Friendship already exits: ")
	} else if friendshipErr == error.DBERROR {
		fmt.Printf("Insert failed: ")
	}
	fmt.Printf("%s\n", error.GetDescription(friendshipErr))

	// Test RemoveFriend
	if friendshipErr == error.ALREADY {
		removeErr := user1.RemoveFriend(db, user2.Email)
		fmt.Printf("Testing removal of friendship: %s\n", error.GetDescription(removeErr))
	}


	// Test Query
	queryErr := user1.Query(db)
	fmt.Printf("Querying user: %s\n", error.GetDescription(queryErr))

	// Test AddWin/AddLoss
	fmt.Printf("The user before addding a win and 2 losses:\n%s\n", user1)
	user1.AddWin(db)
	user1.AddLoss(db)
	user1.AddLoss(db)

	user1.Query(db)
	fmt.Printf("The user after adding the win and losses:\n%s\n", user1)

	roomUser := user.User{}
	roomUser.Email = "seanmyers0608@gmail.com"

	roomErr, testRoom := roomUser.CreateGameroom(db)
	fmt.Printf("Testing gameroom.Create: %v\n", error.GetDescription(roomErr))
//	testRoom.Query(db)
	fmt.Printf("Here's the room: %v\n", testRoom)
	roomQueryErr := testRoom.Query(db)
	fmt.Printf("Query gave error: %v\n", error.GetDescription(roomQueryErr))

	fmt.Printf("Here's the room after query: %v\n", testRoom)


	/*****************************
	 * Added to test functionality
	 * of EmailValidation
	 *****************************/
	testEmail := user.User{}
	testEmail.Email = "zaxcoding@gmail.com"
	testEmailErr := testEmail.EmailValidation(db)

	fmt.Printf("Testing EmailValidation: %v\n", error.GetDescription(testEmailErr))

	/*****************************
	 * END Added to test functionality
	 * of EmailValidation
	 *****************************/

}


func makeDB(host string, database_name string) (mgo.Database){
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	db := session.DB(database_name)
	return *db
}

func clear_database(host string, database_name string){
	db := makeDB(host, database_name)
	db.DropDatabase()
}
