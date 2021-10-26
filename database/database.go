// the database module
package database

// import packages
import (
    _ "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var (
    // Database instance => DB Gorm connector
    DBConn *gorm.DB
)
