package repository
import (
	"fmt"
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
)
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}
func NewPostgresDB(cfg Config, logrus *logrus_log.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logrus.Fatalf("failed check db configs.%v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logrus.Fatalf("fail ping to db %v", err)
		return nil, err
	}
	return db, nil
}
