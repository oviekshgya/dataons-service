package config

import (
	"database/sql"
	"fmt"
	grmsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func MigrateTable(db *gorm.DB) {
	//db.AutoMigrate(&product.Product{}, &category.Category{})

}

func CheckDb() {
	configuration := GetConfig()

	// Menghubungkan ke MySQL tanpa memilih database tertentu
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", configuration.DatabaseMysql.Username, configuration.DatabaseMysql.Password, configuration.DatabaseMysql.Host, configuration.DatabaseMysql.Port)
	db, err := gorm.Open(grmsql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke MySQL:", err)
	}

	// Membuat database jika belum ada
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan instance DB:", err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			fmt.Println("Gagal mendapatkan instance DB:", err)
		}
	}(sqlDB)

	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + configuration.DatabaseMysql.Dbname).Error; err != nil {
		log.Fatal("Gagal membuat database:", err)
	}

	fmt.Println("Database berhasil dibuat atau sudah tersedia.")
}

func SetupMainMysqlDB() *gorm.DB {
	CheckDb()
	configuration := GetConfig()

	database := configuration.DatabaseMysql.Dbname
	host := configuration.DatabaseMysql.Host
	port := configuration.DatabaseMysql.Port
	username := configuration.DatabaseMysql.Username
	password := configuration.DatabaseMysql.Password

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	DBMain, err := gorm.Open(grmsql.Open(addr), &gorm.Config{
		//Logger: logger.New(
		//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		//	logger.Config{
		//		LogLevel:                  logger.Info,
		//		IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
		//		Colorful:                  true, // Enable colorful output// Log level: logger.Silent, logger.Error, logger.Warn, logger.Info
		//	},
		//),
	})

	if err != nil {
		log.Println("err connect db main:", err)
	}
	defer MigrateTable(DBMain)
	return DBMain
}
