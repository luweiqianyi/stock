package gorm_model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

const (
	dsn = "root:123456@tcp(127.0.0.1:3306)/stock?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
)

type User struct {
	gorm.Model
	UserName string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func TestGormCreate(t *testing.T) {
	db, err := OpenDB(dsn)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	db.Create(&User{UserName: "nicklaus", Password: "123456"})

	var user User
	tx := db.First(&user, "username = ?", "nicklaus")
	if tx.RowsAffected <= 0 {
		fmt.Printf("user: %v not found\n", "nicklaus")
		return
	}

}

func TestGormUpdate(t *testing.T) {
	db, err := OpenDB(dsn)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	var user User
	tx := db.First(&user, "username = ?", "nicklaus")
	if tx.RowsAffected <= 0 {
		fmt.Printf("user: %v not found\n", "nicklaus")
		return
	}

	// 不能缺少Where子句，并且Where子句要放在Update前面
	tx = db.Model(&user).Where("username=?", "nicklaus").Update("password", "666666")
	if tx.RowsAffected <= 0 {
		fmt.Printf("user: %v not found\n", "nicklaus")
		return
	}
}

func TestGormDelete(t *testing.T) {
	db, err := OpenDB(dsn)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// 对于软删除的记录，First返回的没有这样的记录，虽然记录在数据库中仍然存在
	//var user User
	//tx := db.First(&user, "username = ?", "nicklaus")
	//if tx.RowsAffected <= 0 {
	//	fmt.Printf("user: %v not found\n", "nicklaus")
	//	return
	//}

	// 不能缺少Where子句，并且Where子句要放在Delete前面
	//tx = db.Where("username=?", "nicklaus").Delete(&user)
	// 下面这句转化成的SQL为：SELECT * FROM `users` WHERE username = 'nicklaus' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	//tx = db.Unscoped().Where("username=?", "nicklaus").Delete(&user)
	tx := db.Exec("delete from users where username=?", "leebai")
	fmt.Printf("error: %v\n", tx.Error)

}

func OpenDB(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	return db, err
}
