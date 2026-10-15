package main

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"fmt"
	"log"
	"os"
	"time"
	"gorm.io/gorm/logger"
)

const (
	host = "localhost"
	port = 3306
	user = "root"
	password = ""
	dbname = "stock_db"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	newLogger := logger.New(
  		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
 		logger.Config{
    		SlowThreshold:              time.Second,   // Slow SQL threshold
    		LogLevel:                   logger.Silent, // Log level
    		Colorful:                 	true,          // Enable color
  },
)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(db)
	fmt.Println("Database connection success")

	db.AutoMigrate(&Product{})
	fmt.Println("Database migration success")

	//newProduct := Product{
	//	Name: "Pro 3",
	//	Price: 500000000.00,
	//	Quantity: 5,
	//}
//
	//createProductBD(db, &newProduct)
	//getProductAllDB(db)
	//getProductByIdDB(db, 4)
	//getProductByNameDB(db, "3")

	//updateProduct := &Product{
	//	Name: "Pro Max 20",
	//	Price: 50000.67,
	//	Quantity: 67,
	//}

	//updateProductDB(db, 2, updateProduct)

	//updateProduct := &Product{
	//	Name: "Nokia 3310",
	//	Price: 50000.67,
	//	Quantity: 67,
	//}

	//updateProductbyNameDB(db, "3", updateProduct)

	//deleteProductbyIDDB(db, 1)

	deleteProductbyNameDB(db, "Pro Max 20")

}

func createProductBD(db *gorm.DB, product *Product) {
	result := db.Create(product)

	if result.Error != nil {
		panic("failed to create product")
	}

	fmt.Println("Product created success")
}

func getProductAllDB(db *gorm.DB){
	var products []Product

	result := db.Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get products success")
	fmt.Println(products)
	//return products
}

func getProductByIdDB(db *gorm.DB, id uint) {
	var product Product

	result := db.First(&product, id)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get product by ID success")
	fmt.Println(product)
	//return product
}

func getProductByNameDB(db *gorm.DB, name string) {
	var product Product

	result := db.Where("name LIKE ?", "%" + name + "%").Find(&product)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get product by name success")
	fmt.Println(product)
	//return product
}

func updateProductDB(db *gorm.DB, id int, product *Product) {
	result := db.Model(&Product{}).Where("id = ?", id).Updates(product)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update product success")
	//return product
}

func updateProductbyNameDB(db *gorm.DB, name string, product *Product) {
	result := db.Model(&Product{}).Where("name LIKE ?", "%" + name + "%").Updates(product)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update product success")
	//return product
}

func deleteProductbyIDDB(db *gorm.DB, id int) {	
	// soft delete
	//result := db.Delete(&Product{}, id)
	// hard delete
	result := db.Unscoped().Delete(&Product{}, id)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Delete product success")
	//return product

}

func deleteProductbyNameDB(db *gorm.DB, name string) {
	// soft delete
	//result := db.Where("name LIKE ?", "%" + name + "%").Delete(&Product{})
	// hard delete
	result := db.Where("name LIKE ?", "%" + name + "%").Unscoped().Delete(&Product{})

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Delete product success")
	//return product
	} 