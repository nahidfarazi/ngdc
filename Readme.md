# NGDC (Next-Gen Database Connector)

### Installation
``` sh
    go get -u github.com/nahidfarazi/ngdc
```
#### use fo Database connection

#### example
```go
       package main
    import (
        "gorm.io/gorm"
        "gorm.io/driver/postgres"
        "os"
    )
    type User struct{
        Name string
        Age string
    }
    func main (){
        dsn := os.Getenv("DB_URL")
        db :=  ngdc.ConnectDB(postgres.Open(dsn) &User{})
    }
```

## or

##### for postgres
```go
// postgres
    package main
    import (
        "gorm.io/gorm"
        "gorm.io/driver/postgres"
    )
    type User struct{
        Name string
        Age string
    }
    func main (){
        dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
        dialector := postgres.Open(dsn)
    db :=  ngdc.ConnectDB(dialector, &User{})
    }
```
##### for mysql
```go
// mysql
    package main
    import (
        "gorm.io/gorm"
        "gorm.io/driver/mysql"
    )
    type User struct{
        Name string
        Age string
    }
    func main (){
        dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
        dialector := mysql.Open(dsn)
    db :=  ngdc.ConnectDB(dialector, &User{})
    }
```
##### for tiDb
```go
// tiDB
    package main
    import (
        "gorm.io/gorm"
        "gorm.io/driver/mysql"
    )
    type User struct{
        Name string
        Age string
    }
    func main (){
        dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
        dialector := mysql.Open(dsn)
    db :=  ngdc.ConnectDB(dialector, &User{})
    }
```
##### for sqlite
```go
// sqlite
    package main
    import (
        "gorm.io/gorm"
        "gorm.io/driver/sqlite"
    )
    type User struct{
        Name string
        Age string
    }
    func main (){
        dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
        dialector := sqlite.Open(dsn)
    db :=  ngdc.ConnectDB(dialector, &User{})
    }
```
##### for sqlserver
```go
// sqlserver
    package main
    import (
        "gorm.io/gorm"
        "gorm.io/driver/sqlserver"
    )
    type User struct{
        Name string
        Age string
    }
    func main (){
        dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
        dialector := sqlserver.Open(dsn)
    db :=  ngdc.ConnectDB(dialector, &User{})
    }
```
##### for clickhouse
```go
// clickhouse
    package main
    import (
        "gorm.io/gorm"
        "gorm.io/driver/mysql"
    )
    type User struct{
        Name string
        Age string
    }
    func main (){
        dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
        dialector := clickhouse.Open(dsn)
    db :=  ngdc.ConnectDB(dialector, &User{})
    }
```

#### ** you can pas more than one struct object **
