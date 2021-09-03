package main

import (
	"database/sql"
	"time"
)

type Test struct {
	Name string
	SQL  string
}

type TestResult struct {
	Duration time.Duration
}

func RunTest(db *sql.DB, test Test) (TestResult, error) {
	var result TestResult

	startTime := time.Now()
	_, err := db.Exec(test.SQL)
	endTime := time.Now()
	if err != nil {
		return TestResult{}, err
	}

	result.Duration = endTime.Sub(startTime)

	return result, nil
}
