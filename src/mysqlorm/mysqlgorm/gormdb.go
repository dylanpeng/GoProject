//gorm demo
//DB Instance
package mysqlgorm

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbRepository *gorm.DB
var lock *sync.Mutex = &sync.Mutex{}

func GetDBInstance() (*gorm.DB, error) {
	if dbRepository == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbRepository == nil {
			var err error
			dbRepository, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true&loc=Local")
			if err != nil {
				fmt.Errorf("Create dbRepository failed")
				return nil, nil
			}
		}
	}
	return dbRepository, nil
}
