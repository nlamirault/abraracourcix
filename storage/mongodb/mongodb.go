// Copyright (C) 2015, 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongodb

import (
	"errors"
	"fmt"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/storage"
)

const (
	label string = "mongodb"
)

var (
	errDuplicateEntry       = errors.New("Duplicate name exists for the shorturl")
	errCollectionNotCreated = errors.New("Collection could not be created")
	errNotFoundEntry        = errors.New("Entry not found")
)

type mongoDB struct {
	session    *mgo.Session
	database   string
	collection string
}

type mongoDocument struct {
	ID       bson.ObjectId `bson:"_id"`
	ShortURL string        `bson:"shorturl"`
	LongURL  string        `bson:"longurl"`
}

func init() {
	storage.RegisterStorage(label, newMongoDBStorage)
}

func newMongoDBStorage(conf *config.Configuration) (storage.Storage, error) {
	glog.V(1).Infof("Create storage using MongoDB : %s", conf.Storage)
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s", conf.Storage.MongoDB.Address))
	if err != nil {
		return nil, err
	}
	// defer session.Close()
	mongo := &mongoDB{
		session:    session,
		database:   conf.Storage.MongoDB.Database,
		collection: conf.Storage.MongoDB.Collection,
	}
	return mongo, nil
}

func (mongoDB *mongoDB) Name() string {
	return label
}

func (mongoDB *mongoDB) Init() error {
	glog.V(1).Infof("Initialize")
	collection := mongoDB.session.DB(mongoDB.database).C(mongoDB.collection)
	if collection == nil {
		return errCollectionNotCreated
	}
	index := mgo.Index{
		Key:      []string{"$text:shorturl"},
		Unique:   true,
		DropDups: true,
	}
	return collection.EnsureIndex(index)
}

func (mongoDB *mongoDB) List() ([][]byte, error) {
	session, collection, err := mongoDB.getCollection()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	query := collection.Find(nil)
	if query == nil {
		return nil, nil
	}
	urls := [][]byte{}
	var documents []mongoDocument
	if err := query.All(&documents); err != nil {
		return nil, err
	}
	for _, doc := range documents {
		urls = append(urls, []byte(doc.ShortURL))
	}

	return urls, nil
}

func (mongoDB *mongoDB) Get(key []byte) ([]byte, error) {
	glog.V(1).Infof("Search entry with key : %v", string(key))
	session, collection, err := mongoDB.getCollection()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	query := collection.Find(bson.M{"shorturl": string(key)})
	if query == nil {
		return nil, nil
	}
	nb, err := query.Count()
	if err != nil {
		return nil, err
	}
	if nb > 0 {
		url := mongoDocument{}
		if err := query.One(&url); err != nil {
			return nil, err
		}
		glog.V(2).Infof("Find : %s", url)
		return []byte(url.LongURL), nil
	}
	return nil, errNotFoundEntry
}

func (mongoDB *mongoDB) Put(key []byte, value []byte) error {
	glog.V(1).Infof("Put : %v %v", string(key), string(value))
	session, collection, err := mongoDB.getCollection()
	if err != nil {
		return err
	}
	defer session.Close()
	err = collection.Insert(
		&mongoDocument{
			ID:       bson.NewObjectId(),
			ShortURL: string(key),
			LongURL:  string(value),
		},
	)
	if err != nil {
		if mgo.IsDup(err) {
			err = errDuplicateEntry
		}
		return err
	}
	return nil
}

func (mongoDB *mongoDB) Delete(key []byte) error {
	glog.V(1).Infof("Delete : %v", string(key))
	return storage.ErrNotImplemented
}

func (mongoDB *mongoDB) Close() error {
	glog.V(1).Infof("Close")
	if mongoDB.session != nil {
		mongoDB.session.Close()
	}
	return nil
}

func (mongoDB *mongoDB) Print() error {
	glog.V(1).Infof("Storage backend: %s", label)
	return nil
}

func (mongoDB *mongoDB) getSession() (*mgo.Session, error) {
	if mongoDB.session != nil {
		return mongoDB.session.Copy(), nil
	}
	return nil, errors.New("No session found")
}

func (mongoDB *mongoDB) getCollection() (*mgo.Session, *mgo.Collection, error) {
	session, err := mongoDB.getSession()
	if err != nil {
		return nil, nil, err
	}
	return session, session.DB(mongoDB.database).C(mongoDB.collection), nil
}
