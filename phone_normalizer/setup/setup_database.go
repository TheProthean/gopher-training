package setup

import (
	"fmt"

	"github.com/jinzhu/gorm"

	//We need this to connect to postgres and setup database
	_ "github.com/lib/pq"
)

//PhoneNumber is a struct where we store all our phone numbers, so we can then update our database
type PhoneNumber struct {
	ID          int `gorm:"primary_key"`
	PhoneNumber string
}

//DatabaseName is database name, that already should exist
var DatabaseName = "phone_numbers_db"

//ConnConfig is our connection parameters string that we can reuse in main
var ConnConfig = fmt.Sprintf("user=postgres password=postgres dbname=%s", DatabaseName)

//Database is a function where we create/update our phone numbers table
func Database() {
	db, err := gorm.Open("postgres", ConnConfig)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	if db.HasTable(&PhoneNumber{}) {
		db.DropTable(&PhoneNumber{})
	}
	db.AutoMigrate(&PhoneNumber{})
	db.Create(&PhoneNumber{PhoneNumber: "1234567890"})
	db.Create(&PhoneNumber{PhoneNumber: "123 456 7891"})
	db.Create(&PhoneNumber{PhoneNumber: "(123) 456 7892"})
	db.Create(&PhoneNumber{PhoneNumber: "(123) 456-7893"})
	db.Create(&PhoneNumber{PhoneNumber: "123-456-7894"})
	db.Create(&PhoneNumber{PhoneNumber: "123-456-7890"})
	db.Create(&PhoneNumber{PhoneNumber: "1234567892"})
	db.Create(&PhoneNumber{PhoneNumber: "(123)456-7892"})
	fmt.Println("Successfully created/updated database.")
}
