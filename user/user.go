package user

import (
	"fmt"
	"github.com/iph/catan/error"
	"github.com/iph/catan/util"
	"github.com/iph/catan/gameroom"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"math/rand"
	"net/smtp"
	"strconv"
)

type User struct {
	Email    string
	Phone    string
	Fname    string
	Lname    string
}

type Friendship struct {
	User1 string
	User2 string
}

type EmailHash struct {
	Email string
	Hash  string
}

/************************************
 * The go equivalent of toString()
 ***********************************/
func (u User) String() string {
	return fmt.Sprintf("Email: %s, Name: %s %s",
		u.Email, u.Fname, u.Lname)
}

/************************************
 * This uses the gmail account
 * 'email.validator.catan@gmail.com'
 * to send an email authorization.
 * Return Codes:
 * SUCESS if sending email worked
 * WEFUCKEDUP if the smtp.SendMail fails
 * DBERROR if insert to DB failed
 * NOTFOUND if user not found in DB
 ***********************************/
/* TODO: CHANGE THIS TO A DIFFERENT VALIDATION
func (u *User) EmailValidation(db mgo.Database) (returnCode int) {

	err := u.Query(db)

	if err != error.SUCCESS {
		return error.NOTFOUND
	}

	// this is for the gmail account I made, we will eventually change this
	auth := smtp.PlainAuth("",
		"email.validator.catan@gmail.com",
		"catanisfun",
		"smtp.gmail.com")

	// pick a random number and tack it onto the end of their email, then hash
	randomNumber := rand.Int()
	hash := util.GenerateSHA512Hash(u.Email + strconv.Itoa(randomNumber))

	// send them the email with their hash embedded in the link
	emailContents := fmt.Sprintf("Welcome to Settlers of Catan, %s %s!\n\n"+
		"To verify your email, please click this link:"+
		"127.0.0.1:8080/user/verify-email?"+
		"user=%s&hash=%s", u.Fname, u.Lname, u.Email, hash)

	// see http://golang.org/pkg/net/smtp/#SendMail for arguments
	emailErr := smtp.SendMail("smtp.gmail.com:587",
		auth,
		"email-validator@catan.com",
		[]string{u.Email},
		[]byte(emailContents))

	// this means an error with SendMail
	if emailErr != nil {
		fmt.Print("ERROR: attempting to send a mail ", err)
		return error.WEFUCKEDUP
	}

	// insert into the table 'user_hash'
	column := EmailHash{u.Email, hash}
	insertErr := db.C("user_hash").Insert(column)

	if insertErr != nil {
		return error.DBERROR
	}

	return error.SUCCESS

}
*/
/************************************
 * Return Codes:
 * SUCCESS if hash matches up
 * DBERROR if our RemoveAll call fails
 * NOTFOUND if hash doesn't match
 ***********************************/
func (u *User) CheckEmailHash(db mgo.Database, theirHash string) (returnCode int) {
	returnCode = error.NOTFOUND

	result := EmailHash{}

	db.C("user_hash").Find(bson.M{"email": u.Email, "hash": theirHash}).One(&result)

	if result.Hash == theirHash {
		_, err := db.C("user_hash").RemoveAll(result)

		if err == nil {
			returnCode = error.SUCCESS
		} else {
			returnCode = error.DBERROR
		}
	}

	return returnCode
}


/************************************
 * Return codes:
 * SUCCESS if insert worked
 * DBERROR if issue during insert
 * ALREADY if user already in DB
 ***********************************/
func (u *User) New(db mgo.Database) (returnCode int) {
	collection := db.C("users")

	// Check if user is in database
	if u.Query(db) != error.SUCCESS {
		err := collection.Insert(*u)

		if err != nil {
			// There was an error
			returnCode = error.DBERROR
		} else {
			// Insert was succsseful
			returnCode = error.SUCCESS
		}
	} else {
		// User is already in database
		returnCode = error.ALREADY
	}

	return returnCode
}

/************************************
 * Return Codes:
 * SUCCESS		if friendship was in database,
 * NOTFOUND 	if one or both of the users
 *						is not in the DB
 * NOTFRIENDS	if they aren't friends
 * TODO: Phone is optional as well.
 ************************************/
func (u *User) QueryFriend(db mgo.Database, friendEmail string) (returnCode int) {

	friend := User{}
	friend.Email = friendEmail

	// Check that they're both in the database
	if u.Query(db) != error.SUCCESS || friend.Query(db) != error.SUCCESS {
		return error.NOTFOUND
	}

	// See if they're already friends, both ways since commutative
	result := Friendship{}
	result2 := Friendship{}
	db.C("friends").Find(bson.M{"user1": u.Email, "user2": friendEmail}).One(&result)
	db.C("friends").Find(bson.M{"user2": u.Email, "user1": friendEmail}).One(&result2)

	// if they are, say so
	if (result.User1 == u.Email && result.User2 == friendEmail) ||
		(result.User1 == friendEmail && result.User1 == u.Email) {
		return error.SUCCESS
	}

	// otherwise not found
	return error.NOTFRIENDS
}

/************************************
 * Return codes:
 * SUCCESS if insert worked
 * NOTFOUND if one or both of the users
 *					are not in the DB
 * ALREADY if friendship already in DB
 * DBERROR if the insert failed
TODO: Make it so a friend is a phone as well.
 ***********************************/
func (u *User) AddFriend(db mgo.Database, friendEmail string) (returnCode int) {

	alreadyFriends := u.QueryFriend(db, friendEmail)

	if alreadyFriends == error.SUCCESS {
		returnCode = error.ALREADY
	} else if alreadyFriends == error.NOTFOUND {
		returnCode = error.NOTFOUND
	} else if alreadyFriends == error.NOTFRIENDS {

		testFriendship := Friendship{u.Email, friendEmail}

		err3 := db.C("friends").Insert(testFriendship)

		if err3 == nil {
			returnCode = error.SUCCESS
		} else {
			returnCode = error.DBERROR
		}
	}

	return returnCode
}

/************************************
 * Return codes:
 * SUCCESS 		if removal worked
 * DBERROR 		if removal failed
 * NOTFOUND 	if one or both of the users
 *						are not in the DB
 * NOTFRIENDS if the users aren't friends
 ***********************************/
func (u *User) RemoveFriend(db mgo.Database, friendEmail string) (returnCode int) {

	alreadyFriends := u.QueryFriend(db, friendEmail)

	if alreadyFriends == error.SUCCESS {
		testFriendship := Friendship{u.Email, friendEmail}
		testFriendship2 := Friendship{friendEmail, u.Email}

		changeInfo1, _ := db.C("friends").RemoveAll(testFriendship)
		changeInfo2, _ := db.C("friends").RemoveAll(testFriendship2)

		if changeInfo1.Removed > 0 || changeInfo2.Removed > 0 {
			returnCode = error.SUCCESS
		} else {
			returnCode = error.DBERROR
		}nn
	} else if alreadyFriends == error.NOTFOUND {
		returnCode = error.NOTFOUND
	} else if alreadyFriends == error.NOTFRIENDS {
		returnCode = error.NOTFRIENDS
	}

	return returnCode
}

/************************************
 * Return Codes:
 * SUCCESS if user was in database,
 * NOTFOUND otherwise
 ************************************/
func (u *User) Query(db mgo.Database) (returnCode int) {
	email := u.Email
	phone := u.Phone

	result := User{}
	var err = db.C("users").Find(bson.M{"email": email, "phone": phone}).One(&result)

	if err != nil {
		// User not found
		return error.NOTFOUND
	} else {
		// User found
		u.Email = result.Email
		u.Phone = result.Phone
		u.Fname = result.Fname
		u.Lname = result.Lname
		return error.SUCCESS
	}
}

