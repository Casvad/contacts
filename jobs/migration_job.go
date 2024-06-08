package jobs

import (
	"contacts/utils/env"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"net/url"
)

type MigrationJob struct {
}

func (r *MigrationJob) Execute() {

	m, err := migrate.New(env.MigrationPath, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.DbUser, url.QueryEscape(env.DbPassword), env.DbHost, env.DbPort, env.DbName))
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		if err.Error() != "no change" {
			panic(err)
		}
	}
}

func ProvideMigrationJob() MigrationJob {

	return MigrationJob{}
}
