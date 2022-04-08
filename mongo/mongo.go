package tokens

import (
	"context"
	"fmt"

	"testMEDOS/users"

	"go.mongodb.org/mongo-driver/bson"
)

func (t *Tokens) checkDb(refreshToken string, guid string) (bool, error) {
	err := t.db.Ping(context.Background(), nil)
	if err != nil {
		return false, err
	}
	collection := t.db.Database("auth").Collection("users")

	if err != nil {
		return false, err
	}

	cursor, err := collection.Find(context.Background(), bson.M{"user": bson.M{"guid": guid}})
	if err != nil {
		return false, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			return false, err
		}
		str := fmt.Sprintf("%v", episode["refreshtoken"])
		res := CheckTokenHash(refreshToken, str)
		if res {
			id := episode["_id"]
			_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
			if err != nil {
				return false, nil
			}
			return true, nil
		}
	}
	return false, nil
}

func (t *Tokens) createDb(u users.User, at AuthToken) error {
	err := t.db.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	collection := t.db.Database("UserSession").Collection("sessions")
	var s Session
	s.User = u
	s.RefreshToken, err = HashToken(at.RefreshToken)
	fmt.Println(s.RefreshToken)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), s)
	if err != nil {
		return err
	}

	return nil
}
