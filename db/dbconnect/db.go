package dbconnect

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/richp10/dat"
	"github.com/richp10/dat/sqlx-runner"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Connect() (*runner.DB, error) {
	// create a normal database connection through database/sql
	db, err := sql.Open("postgres", dsn())
	if err != nil {
		return nil, err
	}

	// ensures the database can be pinged with 15 min backoff
	runner.MustPing(db)

	// set to reasonable values for production
	viper.SetDefault("SetMaxIdleConns", 20)
	viper.SetDefault("SetMaxOpenConns", 40)

	db.SetMaxIdleConns(viper.GetInt("SetMaxIdleConns"))
	db.SetMaxOpenConns(viper.GetInt("SetMaxOpenConns"))

	// set this to enable interpolation
	dat.EnableInterpolation = true

	// Check things like sessions closing. Disable in production
	dat.Strict = false

	// Log any query over 10ms as warnings. (optional)
	viper.SetDefault("LogQueriesThreshold", 10)
	t := viper.GetInt("LogQueriesThreshold")
	runner.LogQueriesThreshold = time.Duration(t) * time.Millisecond

	return runner.NewDB(db, "postgres"), nil
}

// DSN creates a postgres connection string
func dsn() string {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		viper.GetString("DBUSER"),
		viper.GetString("DBPASS"),
		viper.GetString("DBHOST"),
		viper.GetInt("DBPORT"),
		viper.GetString("DBNAME"),
	)
	if !viper.GetBool("DBSSL") {
		dsn += "?sslmode=disable"
	}
	return dsn
}
