package gameroom

import (
	"github.com/iph/catan/error"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type GameRoom struct {
	Id 					bson.ObjectId `bson:"_id,omitempty"`
	Creator     string
	Players     []string
	Invitations []string
}

/************************************
 * The go equivalent of toString()
 ***********************************/
func (g GameRoom) String() string {
	return fmt.Sprintf("Id, %s, Creator: %s, Players: %v, Invitations: %v",
		g. Id, g.Creator, g.Players, g.Invitations)
}

/************************************
 * Return Codes:
 * SUCCESS if user was in database,
 * NOTFOUND otherwise
 ************************************/
func (g *GameRoom) Query(db mgo.Database) (returnCode int) {
	result := GameRoom{}
	err := db.C("game_room").Find(bson.M{"_id": g.Id}).One(&result)	

	if err != nil {
		// GameRoom not found
		return error.NOTFOUND
	} else {
		// GameRoom found
		g.Id = result.Id
		g.Creator = result.Creator
		g.Players = result.Players
		g.Invitations = result.Invitations
		return error.SUCCESS	
	}
}
