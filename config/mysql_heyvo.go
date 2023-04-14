package config

import (
	"Service-API/exception"
	"context"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"
)

func MysqlHeyvoUtilitiesConnection() *gorm.DB {
	ctx, cancel := NewMySQLHeyvoUtilitiesContext()
	defer cancel()

	sqlDB, err := sql.Open("mysql", os.Getenv("MYSQL_HEYVO_UTILITIES_HOST"))
	exception.PanicIfNeeded(err)

	err = sqlDB.PingContext(ctx)
	exception.PanicIfNeeded(err)

	mysqlPoolMax, err := strconv.Atoi(os.Getenv("MYSQL_HEYVO_UTILITIES_POOL_MAX"))
	exception.PanicIfNeeded(err)

	mysqlIdleMax, err := strconv.Atoi(os.Getenv("MYSQL_HEYVO_UTILITIES_IDLE_MAX"))
	exception.PanicIfNeeded(err)

	mysqlMaxLifeTime, err := strconv.Atoi(os.Getenv("MYSQL_HEYVO_UTILITIES_MAX_LIFE_TIME_MINUTE"))
	exception.PanicIfNeeded(err)

	mysqlMaxIdleTime, err := strconv.Atoi(os.Getenv("MYSQL_HEYVO_UTILITIES_MAX_IDLE_TIME_MINUTE"))
	exception.PanicIfNeeded(err)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(mysqlIdleMax)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(mysqlPoolMax)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlMaxLifeTime) * time.Minute)

	sqlDB.SetConnMaxIdleTime(time.Duration(mysqlMaxIdleTime) * time.Minute)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	exception.PanicIfNeeded(err)
	return gormDB
}

func NewMySQLHeyvoUtilitiesContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
