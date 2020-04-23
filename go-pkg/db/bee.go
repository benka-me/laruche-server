package db

import (
	"context"
	"fmt"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (db *DB) GetBee(req laruche.Namespace) (*laruche.Bee, error) {
	bee := &laruche.Bee{}
	filter := bson.D{{"$and", bson.A{
		bson.D{{"name", req.GetName()}},
		bson.D{{"author", req.GetAuthor()}},
	}}}
	err := db.Bees.FindOne(context.TODO(), filter).Decode(&bee)

	return bee, err
}

func (db *DB) GetAllBees(ctx context.Context) (laruche.Beez, error) {
	bees := make(laruche.Beez, 0)
	cur, err := db.Bees.Find(context.TODO(), bson.D{})
	if err != nil {
		return bees, err
	}
	if err = cur.All(ctx, &bees); err != nil {
		return bees, err
	}
	return bees, nil
}
func (db *DB) GetBees(req []laruche.Namespace) (laruche.Beez, string) {
	bees := make(laruche.Beez, 0)
	var msgs string
	for _, r := range req {
		b, err := db.GetBee(r)
		if err != nil {
			msgs = fmt.Sprintf("%s\n%s", msgs, err.Error())
		} else {
			bees = append(bees, b)
		}
	}
	return bees, msgs
}

func (db *DB) NamespaceExist(req laruche.Namespace) bool {
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"$and", bson.A{
				bson.D{{"name", req.GetName}},
				bson.D{{"author", req.GetAuthor}},
			}}},
		}},
	}
	return nil == db.Bees.FindOne(context.TODO(), filter).Err()
}

func (db *DB) BeeExist(req *laruche.Bee) bool {
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"$and", bson.A{
				bson.D{{"name", req.Name}},
				bson.D{{"author", req.Author}},
			}}},
			bson.D{{"repo", req.Port}},
		}},
	}
	return nil == db.Bees.FindOne(context.TODO(), filter).Err()
}

func (db *DB) InsertBee(bee *laruche.Bee) (string, error) {
	if db.BeeExist(bee) {
		return "", status.Error(codes.AlreadyExists, bee.GetNamespaceStr()+" already exist")
	}
	_, err := db.Bees.InsertOne(context.TODO(), bee)
	if err != nil {
		return "", err
	}

	return bee.GetNamespaceStr(), nil
}

func (db *DB) SetBee(bee *laruche.Bee) error {
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"$and", bson.A{
				bson.D{{"name", bee.Name}},
				bson.D{{"author", bee.Author}},
			}}},
			bson.D{{"repo", bee.Port}},
		}},
	}
	update := bson.D{{"$set",
		bee,
	}}
	_, err := db.Bees.UpdateOne(context.TODO(), filter, update)
	return err
}
