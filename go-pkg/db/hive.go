package db

import (
	"context"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (db *DB) GetHive(req laruche.Namespace) (*laruche.Hive, error) {
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"name", req.GetName()}},
			bson.D{{"author", req.GetAuthor()}},
		}},
	}
	ret := &laruche.Hive{}
	err := db.Hives.FindOne(context.TODO(), filter).Decode(ret)
	return ret, err
}

func (db *DB) HiveExist(hive *laruche.Hive) bool {
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"name", hive.Name}},
			bson.D{{"author", hive.Author}},
		}},
	}
	return nil == db.Hives.FindOne(context.TODO(), filter).Err()
}

func (db *DB) InsertHive(hive *laruche.Hive) (string, error) {
	if db.HiveExist(hive) {
		return "", status.Error(codes.AlreadyExists, hive.GetNamespaceStr()+" already exist")
	}
	_, err := db.Hives.InsertOne(context.TODO(), hive)
	if err != nil {
		return "", err
	}
	return hive.GetNamespaceStr(), nil
}

func (db *DB) SetHive(req *laruche.Hive) error {
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"name", req.GetName()}},
			bson.D{{"author", req.GetAuthor()}},
		}},
	}
	update := bson.D{
		{"$set", req},
	}
	_, err := db.Hives.UpdateOne(context.TODO(), filter, update)
	return err
}
