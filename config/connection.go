package config

import (
	"database/sql"
	"dataons-service/models/company"
	"dataons-service/models/department"
	"dataons-service/models/division"
	"dataons-service/models/employee"
	"fmt"
	grmsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func MigrateTable(db *gorm.DB) {
	err := db.AutoMigrate(&company.Company{}, &department.Department{}, &division.Division{}, &employee.Employee{})
	if err != nil {
		log.Fatal(err)
		return
	}

	var datacompany = []company.Company{
		{NameCompany: "PT Indodev Niaga Internet", IsActive: 1, Address: "Jl. Tegal Rotan Raya No.78, Sawah Baru, Kec. Ciputat, Kota Tangerang Selatan, Banten 15413"},
	}

	if err := db.FirstOrCreate(&datacompany).Error; err != nil {
		log.Fatal("Gagal mengisi compay awal:", err)
	}

	var dataDepart = []department.Department{}
	if len(datacompany) > 0 {
		for i := 0; i < len(datacompany); i++ {
			dataDepart = []department.Department{
				{IdCompany: datacompany[i].IdCompany, NameDepartment: "Board Of Directors", IsActive: 1},
				{IdCompany: datacompany[i].IdCompany, NameDepartment: "Information Technology", IsActive: 1},
			}
		}
	}

	if err := db.FirstOrCreate(&dataDepart).Error; err != nil {
		log.Fatal("Gagal mengisi department awal:", err)
	}

	var datadivision = []division.Division{}
	if len(dataDepart) > 0 {
		for i := 0; i < len(dataDepart); i++ {
			datadivision = []division.Division{
				{IdDepartment: dataDepart[i].IdDepartment, NameDivision: "ERP Development", IsActive: 1},
				{IdDepartment: dataDepart[i].IdDepartment, NameDivision: "Tech Development", IsActive: 1},
				{IdDepartment: dataDepart[i].IdDepartment, NameDivision: "Software Maintenance", IsActive: 1},
				{IdDepartment: dataDepart[i].IdDepartment, NameDivision: "Quality Assurance", IsActive: 1},
			}
		}
	}

	if err := db.FirstOrCreate(&datadivision).Error; err != nil {
		log.Fatal("Gagal mengisi data division awal:", err)
	}

	var dataemployee = []employee.Employee{}

	if len(datadivision) > 0 {
		for i := 0; i < len(datadivision); i++ {
			dataemployee = []employee.Employee{
				{IdDivision: datadivision[i].IdDivision, NameEmployee: "OVIEK SHAGYA GHINULUR", NPK: "2021.001", IsActive: 1},
			}
		}
	}

	if err := db.FirstOrCreate(&dataemployee).Error; err != nil {
		log.Fatal("Gagal mengisi data employee awal:", err)
	}

	fmt.Println("Data awal berhasil ditambahkan.")

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
