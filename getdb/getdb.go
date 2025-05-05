package getdb

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// DB接続情報を格納する構造体
type User struct {
	StudentID string `json:"student_id" db:"student_id"` // varchar -> string, JSONとDBタグを追加 (任意)
}

type Job struct {
	JobID       int    `json:"job_id" db:"job_id"`             // int (自動採番) -> int
	JobName     string `json:"job_name" db:"job_name"`         // varchar -> string
	Location    string `json:"location" db:"location"`         // varchar -> string
	DetailsLink string `json:"details_link" db:"details_link"` // varchar -> string
}

type Shift struct {
	ShiftID   int       `json:"shift_id" db:"shift_id"`     // int (自動採番) -> int
	StudentID string    `json:"student_id" db:"student_id"` // varchar (FK) -> string
	JobID     int       `json:"job_id" db:"job_id"`         // int (FK) -> int
	ShiftDate time.Time `json:"shift_date" db:"shift_date"` // date -> time.Time (日付のみが意味を持つ)
	StartTime time.Time `json:"start_time" db:"start_time"` // time -> time.Time (時刻のみが意味を持つ)
	EndTime   time.Time `json:"end_time" db:"end_time"`     // time -> time.Time (時刻のみが意味を持つ)
}

func GetDB() {
	// DB接続情報
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=shiftgo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("DB接続エラー:", err)
		return
	}
	defer db.Close()

	// DB接続確認
	err = db.Ping()
	if err != nil {
		fmt.Println("DB接続エラー:", err)
		return
	}
	fmt.Println("DB接続成功")

}
