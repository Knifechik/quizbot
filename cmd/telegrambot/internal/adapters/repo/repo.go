package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/sipki-tech/database/connectors"
	"github.com/sipki-tech/database/migrations"
	"log"
	"tgbot/cmd/telegrambot/internal/app"
)

var _ app.Repo = &Repo{}

type Repo struct {
	sql *sql.DB
}

func New(ctx context.Context) *Repo {
	const driverName = "postgres"
	const migrateDir = "./cmd/telegrambot/migrate"
	const dsn = "dbname=postgres password=pass123 user=user123 sslmode=disable host=localhost port=5432"
	connector := connectors.PostgresDB{
		User:     "user123",
		Password: "pass123",
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		Parameters: &connectors.PostgresDBParameters{
			ApplicationName: "",
			Mode:            connectors.PostgresSSLDisable,
			SSLRootCert:     "",
			SSLCert:         "",
			SSLKey:          "",
		},
	}

	migrates, err := migrations.Parse(migrateDir)
	if err != nil {
		panic(err)
	}

	err = migrations.Run(ctx, driverName, &connector, migrations.Up, migrates)
	if err != nil {
		panic(err)
	}

	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	return &Repo{sql: conn}
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) CheckChatExist(ctx context.Context, chatID int) bool {
	var exist bool
	const query = `SELECT EXISTS(SELECT 1 FROM user_table WHERE id=$1)`
	r.sql.QueryRowContext(ctx, query, chatID).Scan(&exist)
	return exist
}

func (r *Repo) CheckFinished(ctx context.Context, chatID int) bool {
	var finished bool
	const query = `SELECT finished FROM user_table WHERE id=$1`
	r.sql.QueryRowContext(ctx, query, chatID).Scan(&finished)
	return finished
}

func (r *Repo) CreateChat(ctx context.Context, upd app.UserInfo) {

	quests, err := json.Marshal(upd.Quests)
	if err != nil {
		log.Println(err)
	}

	const query = `INSERT INTO user_table (id, quest_number, last_message, right_answer, finished, quests) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = r.sql.ExecContext(ctx, query, upd.ChatID, upd.QuestNumber, upd.LastMessageID, upd.CountRightAnswer, upd.Finished, quests)
	if err != nil {
		log.Println(err)
	}
}

func (r *Repo) SaveMessage(ctx context.Context, user app.UserInfo) {

	quests, err := json.Marshal(user.Quests)
	if err != nil {
		log.Println(err)
	}

	const query = `UPDATE user_table SET quest_number=$1, last_message=$2, right_answer=$3, quests=$4 WHERE id=$5`
	_, err = r.sql.ExecContext(ctx, query, user.QuestNumber, user.LastMessageID, user.CountRightAnswer, quests, user.ChatID)
	if err != nil {
		log.Println(err)
	}
}

func (r *Repo) GetInfo(ctx context.Context, chatID int) app.UserInfo {
	update := UserInfo{}
	var marsh []byte
	var quest []app.Questions
	const query = `SELECT * FROM user_table WHERE id=$1`
	row := r.sql.QueryRowContext(ctx, query, chatID)
	err := row.Scan(&update.ChatID, &update.QuestNumber, &update.LastMessageID, &update.CountRightAnswer, &update.Finished, &marsh)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(marsh, &quest)
	if err != nil {
		log.Println(err)
	}

	u := update.convert()
	u.Quests = quest

	return u
}

func (r *Repo) PlusCounter(ctx context.Context, u app.UserInfo) {
	const query = `UPDATE user_table SET quest_number=$1 WHERE id=$2`
	_, err := r.sql.ExecContext(ctx, query, u.QuestNumber, u.ChatID)
	if err != nil {
		log.Println(err)
	}
}

func (r *Repo) PlusAnswer(ctx context.Context, u app.UserInfo) {
	const query = `UPDATE user_table SET right_answer=$1 WHERE id=$2`
	_, err := r.sql.ExecContext(ctx, query, u.CountRightAnswer, u.ChatID)
	if err != nil {
		log.Println(err)
	}
}

func (r *Repo) SetFinished(ctx context.Context, u app.UserInfo) {
	const query = `UPDATE user_table SET finished=$1 WHERE id=$2`
	_, err := r.sql.ExecContext(ctx, query, true, u.ChatID)
	if err != nil {
		log.Println(err)
	}

}
