package model

import (
	"fmt"
	"go.etcd.io/bbolt"
)

func attrGetAllBase(db *bbolt.DB, bucket []byte, key []byte) []byte {
	var data []byte
	db.View(func(tx *bbolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b0 := tx.Bucket(bucket)
		if b0 == nil {
			return nil
		}
		data = b0.Get(key)
		return nil
	})
	return data
}

func attrSave(db *bbolt.DB, bucket []byte, key []byte, data []byte) {
	db.Update(func(tx *bbolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b0 := tx.Bucket(bucket)
		if b0 == nil {
			return nil
		}
		err := b0.Put(key, data)
		if err != nil {
			fmt.Println(err)
		}
		return err
	})
}

func AttrGroupUserGetAll(db *bbolt.DB, groupId int64, userId int64) []byte {
	return attrGetAllBase(db, []byte("attrs_group_user"), []byte(fmt.Sprintf("%d-%d", groupId, userId)))
}

func AttrGroupUserSave(db *bbolt.DB, groupId int64, userId int64, data []byte) {
	attrSave(db, []byte("attrs_group_user"), []byte(fmt.Sprintf("%d-%d", groupId, userId)), data)
}

func AttrGroupGetAll(db *bbolt.DB, groupId int64) []byte {
	return attrGetAllBase(db, []byte("attrs_group"), []byte(fmt.Sprintf("%d", groupId)))
}

func AttrGroupSave(db *bbolt.DB, groupId int64, data []byte) {
	attrSave(db, []byte("attrs_group"), []byte(fmt.Sprintf("%d", groupId)), data)
}

func AttrUserGetAll(db *bbolt.DB, userId int64) []byte {
	return attrGetAllBase(db, []byte("attrs_user"), []byte(fmt.Sprintf("%d", userId)))
}

func AttrUserSave(db *bbolt.DB, userId int64, data []byte) {
	attrSave(db, []byte("attrs_user"), []byte(fmt.Sprintf("%d", userId)), data)
}