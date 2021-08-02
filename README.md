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

## Sample Response
`POST` /api/vips


Request Body
```json
{
    "name": "Naomi Osaka",
    "country_of_origin": "Japan",
    "eta": "2021-07-23T11:40:45Z",
    "photo": "https://staticg.sportskeeda.com/editor/2021/05/a2752-16206547493861-800.jpg",
    "attributes": [
        "black jersey",
        "white visor"
    ]
}
```


Response Body
```json
{
	"ok": true,
	"message": "Success"
}
```

`GET` /api/vips
```json
{
    "ok": true,
    "data": [
        {
            "id": 1019521482,
            "name": "Naomi Osaka",
            "country_of_origin": "Japan",
            "eta": "2021-07-23T00:00:00Z",
            "photo": "https://staticg.sportskeeda.com/editor/2021/05/a2752-16206547493861-800.jpg",
            "arrived": true,
            "attributes": [
                "black jersey",
                "white visor"
            ]
        },
        {
            "id": 1954788059,
            "name": "Taufik Hidayat",
            "country_of_origin": "Indonesia",
            "eta": "2021-07-20T00:00:00Z",
            "photo": "https://staticg.sportskeeda.com/wp-content/uploads/2012/02/taufik-hidayat.jpg",
            "arrived": false,
            "attributes": [
                "blue jeans",
                "red jacket indonesian team"
            ]
        },
        {
            "id": 2980432840,
            "name": "An San",
            "country_of_origin": "South Korea",
            "eta": "2021-07-20T00:00:00Z",
            "photo": "https://staticg.sportskeeda.com/editor/2021/07/a1951-16276115578462-800.jpg",
            "arrived": true,
            "attributes": [
                "bucket hat",
                "white black jacket"
            ]
        },
        {
            "id": 4217253817,
            "name": "Novak Djokovic",
            "country_of_origin": "Serbia",
            "eta": "2021-07-20T00:00:00Z",
            "photo": "https://staticg.sportskeeda.com/editor/2021/07/a6359-16277321137185-800.jpg",
            "arrived": false,
            "attributes": [
                "blue shorts",
                "red jacket serbia team"
            ]
        }
    ],
    "message": "success"
}
```


`GET` /api/vips/1019521482

Response Body
```json
{
    "ok": true,
    "data": {
        "id": 1019521482,
        "name": "Naomi Osaka",
        "country_of_origin": "Japan",
        "eta": "2021-07-23T00:00:00Z",
        "photo": "https://staticg.sportskeeda.com/editor/2021/05/a2752-16206547493861-800.jpg",
        "arrived": false,
        "attributes": [
            "black jersey",
            "white visor"
        ]
    },
    "message": "Success"
}
```

`PATCH` /api/vips/1019521482/arrived
Request Body
```json
{
	"arrived": true;
}
```

Response Body
```json
{
	"ok": true,
	"message": "Success"
}
```

## Author
William Fu