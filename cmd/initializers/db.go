package initializers

import (
    "log"
    "os"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "scanNstore/cmd/models"
)

var DB *gorm.DB


func ConnectDB() {
    var err error

   
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the Database! \n", err.Error())
        os.Exit(1)
    }

    log.Println("Running Migrations")

    // Run migrations for User, Receipt, and Item models
    err = DB.AutoMigrate( &models.Receipt{})
    if err != nil {
        log.Fatal("Migration Failed: \n", err.Error())
        os.Exit(1)
    }

    log.Println("🚀 Connected Successfully to the Database")
}
