package token

import (
	"github.com/iph/catan/util"
	"github.com/iph/catan/error"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

// this is the duration for the token,
// as of now it can be any time.Duration quantity
const EXPIRATION_DURATION = 5 * time.Hour

type Token struct {
	Email      string
	Expiration time.Time
	Hash       string
}

func (t *Token) New(db mgo.Database) (returnCode int) {
	t.Hash = util.GenerateSHA512Hash(time.Now().String()) // TODO: Should we make this better?
	t.Expiration = time.Now()
	t.Expiration = t.Expiration.Add(EXPIRATION_DURATION)

	collection := db.C("tokens")

	err := collection.Insert(*t)

	if err != nil {
		// BAD DB, VERY BAD!
		returnCode = error.DBERROR
	} else {
		// kk, we like the db.
		returnCode = error.SUCCESS
	}
	return returnCode
}

func (t *Token) Check(db mgo.Database) (returnCode int) {
	token := Token{}
	var err = db.C("tokens").Find(bson.M{"email": t.Email, "hash": t.Hash}).One(&token)

	if err != nil {
		// The token was not in the db.
		return error.NOTFOUND
	} else {
		// TODO: Check expiration:
		t.Expiration = token.Expiration
		// The token is legit.
		return error.SUCCESS
	}

}
