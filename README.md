# vip-management-system

Simple REST API built with Go and MySQL database

## Install and Run
Clone this repository
```
git clone https://github.com/williammfu/vip-management-system
```

Create database
```
mysql -u <username> -p <password> vip_system < vip_system.sql
```

Install dependencies
```
go mod init
go mod tidy
```

Edit MySQL database configuration on config.go
```
dbInfo := DBInfo{
	Username: "root",
	Password: "",
	Host:     "localhost",
	Port:     "3306"}
```


Build and Run
```
go build
./vip-management-system
```

## Author
William Fu