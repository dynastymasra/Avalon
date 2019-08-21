package config

import "fmt"

type DatabaseConfig struct {
	host        string
	port        int
	name        string
	username    string
	password    string
	maxOpenConn int
	maxIdleConn int
	logEnabled  bool
}

func (d DatabaseConfig) Host() string {
	return d.host
}

func (d DatabaseConfig) Port() int {
	return d.port
}

func (d DatabaseConfig) Name() string {
	return d.name
}

func (d DatabaseConfig) Username() string {
	return d.username
}

func (d DatabaseConfig) Password() string {
	return d.password
}

func (d DatabaseConfig) MaxOpenConns() int {
	return d.maxOpenConn
}

func (d DatabaseConfig) MaxIdleConns() int {
	return d.maxIdleConn
}

func (d DatabaseConfig) LogEnabled() bool {
	return d.logEnabled
}

func (d DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable",
		d.name, d.username, d.password, d.host, d.port)
}
