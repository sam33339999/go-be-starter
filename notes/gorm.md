# GORM

```go
dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
    host,
    username,
    password,
    dbName,
    port,
)

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.GetGormLogger(),
})

if err != nil {
    logger.Info("DSN: ", dsn)
    logger.Panic(err)
}
```


> 建立 db 的 connect 物件 (*gorm.DB) 後，透過這個物件；
並且傳入對應的 structure ，且該 structure 有定義 tags 及實現對應的方法時，即可使用 gorm 中的方法去存取資料庫。



