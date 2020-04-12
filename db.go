package main

import (
	"database/sql"
	"fmt"
	"time"
)

type dbManager struct {
	db *sql.DB
}

func (dbm *dbManager) saveURL(url string, path string, options urlOptions) (URLDetail, error) {
	var ud URLDetail
	qs := `INSERT INTO urldetail (url, path, is_custom, created_at) VALUES ($1, $2, $3, $4)`
	_, err := dbm.db.Exec(qs, url, path, options.isCustom, time.Now().UTC())
	if err != nil {
		return ud, err
	}

	ud = URLDetail{URL: url, Path: path, IsCustom: options.isCustom}
	return ud, nil
}

func (dbm *dbManager) updateUrlPath(url, path string, options urlOptions) (URLDetail, error) {
	var ud URLDetail
	qs := `UPDATE urldetail SET path = $1, is_custom = $2 WHERE url = $3`
	_, err := dbm.db.Exec(qs, path, options.isCustom, url)
	if err != nil {
		return ud, err
	}
	return URLDetail{URL: url, Path: path, IsCustom: options.isCustom}, nil
}

func (dbm *dbManager) getURLByPath(path string) (URLDetail, error) {
	return dbm.get("path", path)
}

func (dbm *dbManager) getPathByURL(url string) (URLDetail, error) {
	return dbm.get("url", url)
}

func (dbm *dbManager) get(field string, by string) (URLDetail, error) {
	var ud URLDetail
	qs := fmt.Sprintf(`SELECT url, path, is_custom FROM urldetail WHERE %s = $1`, field)
	row := dbm.db.QueryRow(qs, by)
	if err := row.Scan(&ud.URL, &ud.Path, &ud.IsCustom); err != nil {
		return ud, err
	}
	return ud, nil
}
