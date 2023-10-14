package data 
import "database/sql"

type Models struct {
	Books BookModel
}

func newModels(db *sql.DB) Models {
	return Models{
		Books: MookModel{DB: db},
	}
}