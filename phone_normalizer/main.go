package main

import (
	"fmt"
	"os"

	"github.com/gopher-training/phone_normalizer/phoneNormalizer"
	"github.com/gopher-training/phone_normalizer/setup"
	"github.com/jinzhu/gorm"
)

func init() {
	setup.Database()
}

func main() {
	db, err := gorm.Open("postgres", setup.ConnConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	phoneNumbers := []setup.PhoneNumber{}
	db.Find(&phoneNumbers)
	fmt.Println("Source database:")
	fmt.Println(phoneNumbers)
	normalizedPhoneNumbers, err := phoneNormalizer.NormalizeNumbers(phoneNumbers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, v := range normalizedPhoneNumbers {
		if v.PhoneNumber == "" {
			db.Delete(&v)
			continue
		}
		db.Save(&v)
	}
	newPhoneNumbers := []setup.PhoneNumber{}
	db.Find(&newPhoneNumbers)
	fmt.Println("Normalized database:")
	fmt.Println(newPhoneNumbers)
}
