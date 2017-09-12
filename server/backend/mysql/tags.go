package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"git.oschina.net/k2ops/jarvis/server/api/helper"
	"encoding/json"
)

func (m *JarvisMysqlBackend) GetHostTags(id string) (tags []byte, err error) {
	db := m.db
	log.WithFields(log.Fields{
		"sql": "SELECT tags FROM jarvis.hosts WHERE systemId="+id,
	}).Info("search alarms")
	rows, err := db.Query("SELECT tags FROM jarvis.hosts WHERE systemId=?", id)
	if err != nil {
		return
	}
	for rows.Next() {
		err = rows.Scan(
			&tags,
		)
		if err != nil {
			return
		}
	}
	return
}

func (m *JarvisMysqlBackend) AttachTag(id string, tag string) (err error) {
	db := m.db

	// update hosts table
	log.WithFields(log.Fields{
		"sql": "UPDATE jarvis.hosts SET tags=JSON_ARRAY_APPEND(tags, '$', ?) WHERE systemId=? and !JSON_CONTAINS(tags, ?)",
	}).Info("attach tag to host")
	result, err := db.Exec("UPDATE jarvis.hosts SET tags=JSON_ARRAY_APPEND(tags, '$', ?) WHERE systemId=? and !JSON_CONTAINS(tags, ?)",
		tag,
		id,
		helper.SafeMarshalJsonArray([]string{tag}),
	)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	log.WithFields(log.Fields{
		"tag": tag,
		"host": id,
		"rows_affected": affected,
	}).Info("attach tag to host")

	if affected==0 {
		return
	}

	// update tags table
	result, err = db.Exec("UPDATE jarvis.tags SET attached=JSON_ARRAY_APPEND(attached, '$', ?) WHERE tag=? AND !JSON_CONTAINS(attached, ?)",
		id,
		tag,
		helper.SafeMarshalJsonArray([]string{id}),
	)

	affected, err = result.RowsAffected()
	if affected==0 {
		// new tag, we should insert a record
		result, err = db.Exec("INSERT INTO jarvis.tags(tag, attached) VALUES (?, ?)", tag, helper.SafeMarshalJsonArray([]string{id}))
		if err != nil {
			return
		}
	}

	return
}


func (m *JarvisMysqlBackend) RemoveTag(id string, tag string) (err error) {
	db := m.db
	// update hosts table
	result, err := db.Exec("UPDATE jarvis.hosts SET tags=JSON_REMOVE(	tags, replace(JSON_SEARCH(tags, 'one', ?), '\"', '')) WHERE JSON_SEARCH(tags, 'one', ?) IS NOT NULL AND systemId=?",
		tag,
		tag,
		id,
	)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	log.WithFields(log.Fields{
		"tag": tag,
		"host": id,
		"rows_affected": affected,
	}).Info("remove tag from host")

	if affected==0 {
		return
	}

	// update tags table
	result, err = db.Exec("UPDATE jarvis.tags SET attached=JSON_REMOVE(attached, replace(JSON_SEARCH(attached, 'one', ?), '\"', '')) WHERE JSON_SEARCH(attached, 'one', ?) IS NOT NULL AND tag=?",
		id,
		id,
		tag,
	)

	return
}

func (m *JarvisMysqlBackend) ListTags() (map[string][]string, error) {
	db := m.db
	//
	tags := make(map[string][]string)

	// query tags table
	rows, err := db.Query("SELECT * FROM jarvis.tags WHERE JSON_CONTAINS_PATH(attached, 'all', '$[0]')")
	if err != nil {
		return tags, err
	}

	for rows.Next() {
		var tag string
		var attached []byte
		err = rows.Scan(
			&tag,
			&attached,
		)
		if err != nil {
			return tags, err
		}

		ids := make([]string, 10)
		err = json.Unmarshal(attached, &ids)
		if err != nil {
			return tags, err
		}
		tags[tag] = ids
	}
	return tags, err
}
