package db

import "github.com/jinzhu/gorm"

var tables = []interface{}{}

func CreateTables(db *gorm.DB) error {
	return db.AutoMigrate(tables...).Error
}

func DropAllTables(db *gorm.DB) error {
	return db.DropTable(tables...).Error
}

func CreateAll(db *gorm.DB) error {
	if err := CreateTables(db); err != nil {
		return err
	}

	fkeys := map[interface{}][][2]string{
		//&Session{}: {
		//	{"user_name", "users(name)"},
		//},
	}

	for model, foreignKey := range fkeys {
		for _, fk := range foreignKey {
			if err := db.Debug().Model(model).AddForeignKey(
				fk[0], fk[1], "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
		}
	}

	//if err := db.Debug().Model(&User{}).AddUniqueIndex(
	//	"idx_user_name",
	//	"name").Error; err != nil {
	//	return err
	//}

	return nil
}
