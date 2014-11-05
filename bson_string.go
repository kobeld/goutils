package goutils

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"strings"
)

func ToObjectId(idHex string) (id bson.ObjectId, err error) {
	if !bson.IsObjectIdHex(idHex) {
		err = fmt.Errorf("Invalid input to ObjectIdHex: %q", idHex)
		return
	}
	id = bson.ObjectId(idHex)
	return
}

// ["1", "2", "3"] -> [ObjectId("1"), ObjectId("2"), ObjectId("3")]
func TurnPlainIdsToObjectIds(idHexes []string) (r []bson.ObjectId, err error) {
	for _, id := range idHexes {
		if strings.Trim(id, " 　") == "" {
			continue
		}
		oId, err := ToObjectId(id)
		if err != nil {
			return r, err
		}
		r = append(r, oId)
	}
	return
}

// [ObjectId("1"), ObjectId("2"), ObjectId("3")]-> ["1", "2", "3"]
func TurnObjectIdToPlainIds(ids []bson.ObjectId) (r []string) {
	for _, id := range ids {
		r = append(r, id.Hex())
	}
	return
}

func IsInObjectIds(tragetId bson.ObjectId, ids []bson.ObjectId) bool {
	for _, id := range ids {
		if tragetId == id {
			return true
		}
	}
	return false
}
