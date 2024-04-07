package mock_database

// Mock Sql Result
type MockSqlResult struct {
	LastInsertedId int64
	RowsAffect     int64
}

func (r MockSqlResult) LastInsertId() (int64, error) {
	return r.LastInsertedId, nil
}

func (r MockSqlResult) RowsAffected() (int64, error) {
	return r.RowsAffect, nil
}
