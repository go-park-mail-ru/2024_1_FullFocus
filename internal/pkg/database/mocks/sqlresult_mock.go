package mock_database

// Mock Sql Result

type MockSQLResult struct {
	LastInsertedID int64
	RowsAffect     int64
}

func (r MockSQLResult) LastInsertId() (int64, error) {
	return r.LastInsertedID, nil
}

func (r MockSQLResult) RowsAffected() (int64, error) {
	return r.RowsAffect, nil
}
