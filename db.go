package main

import (
	"database/sql"
	"net/url"
)

const (
	DefaultUsername = "sa"
	DefaultPassword = ""
)

type DBOptions struct {
	Server   string
	Database string

	Username string
	Password string

	ApplicationName string
}

func (o *DBOptions) connect() (*sql.DB, error) {
	query := url.Values{}
	query.Add("database", o.Database)
	if o.ApplicationName != "" {
		query.Add("app name", o.ApplicationName)
	}

	connectionString := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(o.Username, o.Password),
		Host:     o.Server,
		RawQuery: query.Encode(),
	}
	urlString := connectionString.String()
	return sql.Open("sqlserver", urlString)
}
