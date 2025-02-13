---
create_at: 2024.01.29
update_at: 2024.01.29
draft: false
tags: 
 - Go
 - Gorm
 - Database
---

# Gorm

# Install

這個是 gorm 的主要套件，接下來要根據使用的資料庫安裝各自的驅動。因為我目前只有用過 sqlite，所以先寫 sqlite 的部分。

```shell
go get -u gorm.io/gorm
```

## 安裝 Sqlite 驅動

```shell
go get -u gorm.io/driver/sqlite
```
# Connect

連接到資料庫，如果沒有該資料庫會幫你創建一個。

```go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
if err != nil {
    log.Fatelln("Failed to connect database.")
}
```

# Create table

AutoMigrate 會檢查資料庫內有沒有這個 table，沒有的話會創建一個新的。有的話會檢查 table 的欄位跟傳進去的這個 struct 的欄位有沒有一致，沒有一致的話會幫你變成一致的。所以才需要傳一個 struct 進去，為了給他檢查 table 的欄位。

```go
type Tasl struct {
    gorm.Model // 這行會幫你加上常用的 table 欄位
    Name string
	Status string
}

db.AutoMigrate(&Task{})
```

# Read

## Get one

```go
func GetOneByName(db *gorm.DB, name string) *Task {
	var result Task

	db.Where("name = ?", name).First(&result)

	return &result
}
```

## Get all

```go
func GetAll(db *gorm.DB) *[]Task {
	var results []Task

	db.Find(&results)

	return &results
}
```

# Update

```go
func UpdateStatusByName(db *gorm.DB, name string, pending bool) {
	if pending {
		db.Model(&Task{}).Where("name = ?", name).Update("status", "Pending")
	} else {
		db.Model(&Task{}).Where("name = ?", name).Update("status", "Done")
	}
}
```

在這裏 `db.Model()` 的用法應該是取得 table，像範例傳入了 Task 的實例，所以他應該會取得資料庫裡 Task 的 table。這應該是用在資料庫裡有很多 table 的情況下，不知道要對哪個 table 做操作，所以先把想要做操作的 table 拿出來。

# Delete

```go
func DeleteTaskByName(db *gorm.DB, name string) {
	record := GetOneByName(db, name)
	db.Delete(record)
}
```

# Close database

要先取得底層的資料庫連線池，然後使用它來關掉 DB。gorm 本身是沒有給clode方法的。

```go
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get db")
	}
	sqlDB.Close()
}
```

# Reference

1. [gorm.io](https://gorm.io/)
2. [Go中使用 SQLite 数据库(Gorm)
](https://blog.csdn.net/cnwyt/article/details/118904882)