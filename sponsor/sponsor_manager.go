package sponsor

import (
	"database/sql"
	"fmt"
	"github.com/kraxarn/website/config"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"os"
)

const dbUrl string = "http://sponsor.ajay.app/database.db"

type Manager struct {
	db   *sql.DB
	path string
}

type Time struct {
	startTime, endTime float64
}

func NewManager() (Manager, error) {
	m := Manager{
		db:   nil,
		path: config.GetPath("sponsor.db"),
	}

	pathStat, err := os.Stat(m.path)
	if os.IsNotExist(err) {
		_ = m.update()
	} else {
		fmt.Println("database was last updated", pathStat.ModTime().Format("2006-01-02"))
	}

	m.db, err = sql.Open("sqlite3", m.path)
	return m, err
}

func (manager *Manager) update() error {
	response, err := http.Get(dbUrl)
	if err != nil {
		return err
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println("failed to close response body:", err)
		}
	}()

	outFile, err := os.Create(manager.path)
	if err != nil {
		return err
	}

	_, err = io.Copy(outFile, response.Body)

	return err
}

func (manager *Manager) GetTimes(videoId string) ([]Time, error) {
	times := make([]Time, 0)
	stmt, err := manager.db.Prepare("select startTime, endTime from sponsorTimes where videoID = ?")
	if err != nil {
		return times, err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("failed to close:", err)
		}
	}()

	rows, err := stmt.Query(videoId)
	if err != nil {
		return times, err
	}

	for rows.Next() {
		var time Time
		err = rows.Scan(&time.startTime, &time.endTime)
		if err != nil {
			return times, err
		}
		times = append(times, time)
	}

	return times, nil
}
