package db

import "github.com/AlineConchetta/myapi/API-PostgreSQL/configs"

func OpenConnection(+ sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=/% password=%s dbname=% sslmode=disable",
	conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err =! nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}