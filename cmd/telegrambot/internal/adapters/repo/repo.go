package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sipki-tech/database/connectors"
	"github.com/sipki-tech/database/migrations"
	"tgbot/cmd/telegrambot/internal/app"
)

var _ app.Repo = &Repo{}

type Repo struct {
	sql *sql.DB
}

func New(ctx context.Context) (*Repo, error) {
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
		return nil, fmt.Errorf("migration.Parse: %w", err)
	}

	err = migrations.Run(ctx, driverName, &connector, migrations.Up, migrates)
	if err != nil {
		return nil, fmt.Errorf("migration.Run: %w", err)
	}

	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	return &Repo{sql: conn}, nil
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) Create(ctx context.Context, upd app.UserInfo) error {

	quests, err := json.Marshal(upd.Quests)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	const query = `INSERT INTO user_table (id, quest_number, last_message, right_answer, finished, quests) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = r.sql.ExecContext(ctx, query, upd.ChatID, upd.QuestNumber, upd.LastMessageID, upd.CountRightAnswer, upd.Finished, quests)
	if err != nil {
		return fmt.Errorf("sql.ExecContext: %w", err)
	}
	return nil
}

func (r *Repo) Update(ctx context.Context, user app.UserInfo) error {

	quests, err := json.Marshal(user.Quests)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	const query = `UPDATE user_table SET quest_number=$1, last_message=$2, right_answer=$3,finished=$4, quests=$5 WHERE id=$6`
	_, err = r.sql.ExecContext(ctx, query, user.QuestNumber, user.LastMessageID, user.CountRightAnswer, user.Finished, quests, user.ChatID)
	if err != nil {
		return fmt.Errorf("sql.ExecContext: %w", err)
	}
	return nil
}

// GetInfo TODO refact
func (r *Repo) Get(ctx context.Context, chatID int) (*app.UserInfo, error) {
	update := UserInfo{}
	var marsh []byte
	var quest []app.Questions

	const query = `SELECT * FROM user_table WHERE id=$1`
	row := r.sql.QueryRowContext(ctx, query, chatID)
	err := row.Scan(&update.ChatID, &update.QuestNumber, &update.LastMessageID, &update.CountRightAnswer, &update.Finished, &marsh)
	if err != nil {
		return nil, fmt.Errorf("sql.QueryRowContext: %w", convertErr(err))
	}

	err = json.Unmarshal(marsh, &quest)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	user := update.convert(quest)

	return user, nil
}
