package sqlite3_repository

// func NewConnectionMock() (*sql.DB, sqlmock.Sqlmock) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		log.Fatalf("erro ao fazer conexão do mock, %s", err.Error())
// 	}
// 	return db, mock
// }

/*
func TestNewProductRepository(t *testing.T) {
	t.Run("Delete Products - success", func(t *testing.T) {
		db, mocks := NewConnectionMock()

		productDelete := entities.ProductRequest{ID: 10}
		mocks.ExpectExec(regexp.QuoteMeta(QUERY_DELETE_PRODUCT)).
			WithArgs(productDelete.ID).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnError(nil)

		productRepo := NewProductRepository(db)
		err := productRepo.DeleteProducts(&productDelete)
		assert.Nil(t, err)
	})
	t.Run("Delete Products - fail in query", func(t *testing.T) {
		db, mocks := NewConnectionMock()

		productDelete := entities.ProductRequest{ID: 10}
		expectError := errors.New("errors")
		mocks.ExpectExec(regexp.QuoteMeta(QUERY_DELETE_PRODUCT)).
			WithArgs(productDelete.ID).WillReturnError(expectError)

		productRepo := NewProductRepository(db)
		err := productRepo.DeleteProducts(&productDelete)
		assert.Error(t, err)
		assert.Equal(t, expectError.Error(), err.Error())
	})
	t.Run("Delete Products - fail rows affected", func(t *testing.T) {
		db, mocks := NewConnectionMock()

		productDelete := entities.ProductRequest{ID: 10}
		expectError := errors.New("errors")
		mocks.ExpectExec(regexp.QuoteMeta(QUERY_DELETE_PRODUCT)).
			WithArgs(productDelete.ID).
			WillReturnResult(sqlmock.NewResult(0, 0)).
			WillReturnError(expectError)

		productRepo := NewProductRepository(db)
		err := productRepo.DeleteProducts(&productDelete)
		assert.Error(t, err)
		assert.Equal(t, expectError.Error(), err.Error())
	})
}
*/