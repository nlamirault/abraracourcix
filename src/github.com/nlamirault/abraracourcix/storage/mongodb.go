// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"errors"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DATABASE        string = "abraracourcix"
	URLS_COLLECTION string = "urls"
	DEFAULT_URL     string = "127.0.0.1:27017"
)

// Mongo represents the MongoDB database client
type Mongo struct {
	Session *mgo.Session
}

type mongoDocument struct {
	Id       bson.ObjectId `bson:"_id"`
	ShortUrl string        `bson:"shorturl"`
	LongUrl  string        `bson:"longurl"`
}

// NewMongo instantiates a new MongoDB database client
func NewMongo(url string) (*Mongo, error) {
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s", url))
	if err != nil {
		return nil, err
	}
	// defer session.Close()
	mongo := &Mongo{Session: session}
	err = mongo.Setup()
	if err != nil {
		return nil, err
	}
	return mongo, nil
}

func (db *Mongo) GetSession() (*mgo.Session, error) {
	if db.Session != nil {
		return db.Session.Copy(), nil
	}
	return nil, errors.New("No session found")
}

func (db *Mongo) Setup() error {
	collection := db.Session.DB(DATABASE).C(URLS_COLLECTION)
	if collection == nil {
		return errors.New("Collection could not be created")
	}
	index := mgo.Index{
		Key:      []string{"$text:shorturl"},
		Unique:   true,
		DropDups: true,
	}
	collection.EnsureIndex(index)
	return nil
}

func (db *Mongo) GetCollection() (*mgo.Session, *mgo.Collection, error) {
	session, err := db.GetSession()
	if err != nil {
		return nil, nil, err
	}
	return session, session.DB(DATABASE).C(URLS_COLLECTION), nil
}

// Get a value given its key
func (db *Mongo) Get(key []byte) ([]byte, error) {
	log.Printf("[DEBUG] [abraracourcix] Get : %v", string(key))
	url := mongoDocument{}
	session, collection, err := db.GetCollection()
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
		err = query.One(&url)
		if err != nil {
			return nil, err
		}
		log.Printf("[INFO] [abraracourcix] Find : %v", url)
		return []byte(url.LongUrl), nil
	}
	return nil, nil
}

// Put a value at the specified key
func (db *Mongo) Put(key []byte, value []byte) error {
	log.Printf("[DEBUG] [abraracourcix] Put : %v %v", string(key), string(value))
	session, collection, err := db.GetCollection()
	if err != nil {
		return err
	}
	defer session.Close()
	err = collection.Insert(
		&mongoDocument{
			Id:       bson.NewObjectId(),
			ShortUrl: string(key),
			LongUrl:  string(value),
		},
	)
	if err != nil {
		if mgo.IsDup(err) {
			err = errors.New("Duplicate name exists for the shorturl")
		}
		return err
	}
	return nil
}

// Delete the value at the specified key
func (db *Mongo) Delete(key []byte) error {
	log.Printf("[DEBUG] [abraracourcix] Delete : %v", string(key))
	return nil
}

// Close backend informations
func (db *Mongo) Close() {
	log.Printf("[DEBUG] [abraracourcix] Close")
	if db.Session != nil {
		db.Session.Close()
	}
}

// Print backend informations
func (db *Mongo) Print() {
	log.Printf("[DEBUG] [abraracourcix] Print")
}
