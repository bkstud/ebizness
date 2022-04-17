package main

import (
	"os"
	"project/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func db_init() {

	if _, err := os.Stat("./test.db"); err == nil {
		os.Remove("./test.db")
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// // Migrate the schema
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Worker{})
	db.AutoMigrate(&models.Purchase{})
	db.AutoMigrate(&models.Shift{})

	p1 := models.Product{Code: "D42", Price: 100}
	p2 := models.Product{Code: "D45", Price: 200}
	p3 := models.Product{Code: "D46", Price: 300}
	for _, p := range []*models.Product{&p1, &p2, &p3} {
		db.Create(p)
	}

	c1 := models.Customer{Firstname: "Adam", Surname: "Kowalski"}
	c2 := models.Customer{Firstname: "Janusz", Surname: "Nowak"}
	db.Create(&c1)
	db.Create(&c2)

	w1 := models.Worker{Firstname: "Hugo", Surname: "Rower", Position: "Sorter"}
	w2 := models.Worker{Firstname: "Monika", Surname: "Grozna", Position: "Manager"}
	w3 := models.Worker{Firstname: "Zordon", Surname: "Wygodny", Position: "Employee"}
	db.Create(&w1)
	db.Create(&w2)
	db.Create(&w3)

	purchase1 := models.Purchase{Product: p1, Customer: c1}
	purchase2 := models.Purchase{Customer: c2, Product: p3}
	db.Create(&purchase1)
	db.Create(&purchase2)

	// YYYY-MM-DD HH:MM
	s1 := models.Shift{Worker: w1, EnterTime: "2022-05-22 08:00", ExitTime: "2022-05-22 16:00"}
	s2 := models.Shift{Worker: w2, EnterTime: "2022-05-22 10:00", ExitTime: "2022-05-22 18:00"}
	s3 := models.Shift{Worker: w3, EnterTime: "2022-05-22 09:00", ExitTime: "2022-05-22 17:30"}
	db.Create(&s1)
	db.Create(&s2)
	db.Create(&s3)

	// db.Create(&models.Product{Code: "D43", Price: 100})

	// // Read
	// var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
	// return db
}
