package models

import (
	"database/sql"
	"fmt"

	"github.com/Kakashi944/Appointment_GRPC/config"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Appointment struct {
	ResourceType string
	ID           uint
	Text         string
	Identifier   []string
	Priority     uint
}

type Text struct {
	Status string `json:"status"`
	Div    string `json:"div"`
}

type Identifier struct {
	System string `json:"system"`
	Value  string `json:"value"`
}

func ConnectSQL() (db1 *sql.DB, err error) {
	config.InitializeAppConfig()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBUsername,
		config.AppConfig.DBPassword, config.AppConfig.DBDatabase)

	db1, _ = sql.Open("postgres", psqlconn)

	err = db1.Ping()

	return db1, err
}

func InsertAppointment(appointment Appointment) (id uint, err error) {
	db, _ := ConnectSQL()
	insert := `INSERT INTO "Appoint"."appointments" (resource_type, text, identifier) VALUES ($1,$2,$3) RETURNING id`

	err = db.QueryRow(insert, appointment.ResourceType, appointment.Text, pq.Array(appointment.Identifier)).Scan(&id)

	return id, err

}
