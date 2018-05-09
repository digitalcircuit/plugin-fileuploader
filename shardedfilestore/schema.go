package shardedfilestore

import (
	"github.com/rs/zerolog/log"
	"github.com/rubenv/sql-migrate"
)

func (store *ShardedFileStore) initDB() {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id: "1",
				Up: []string{`CREATE TABLE uploads(
					id VARCHAR(36) PRIMARY KEY,
					uploader_ip BLOB,
					sha256sum BLOB,
					created_at INTEGER(8)
				);`},
				Down: []string{"DROP TABLE uploads;"},
			},
			&migrate.Migration{
				Id: "2",
				Up: []string{`ALTER TABLE uploads
					ADD deleted INTEGER(1) DEFAULT 0 NOT NULL
				;`},
			},
		},
	}

	n, err := migrate.Exec(store.Db, store.DbConfig.DriverName, migrations, migrate.Up)
	if err != nil {
		log.Fatal().Err(err)
	}

	if n > 0 {
		log.Info().Int("count", n).Msg("Applied schema migrations")
	}
}
