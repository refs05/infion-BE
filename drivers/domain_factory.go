package drivers

import (
	threadsDomain "infion-BE/businesses/threads"
	threadsDB "infion-BE/drivers/databases/threads"

	"gorm.io/gorm"
)

//NewThreadsRepository Factory with threads domain
func NewThreadsRepository(conn *gorm.DB) threadsDomain.Repository {
	return threadsDB.NewMySQLRepository(conn)
}