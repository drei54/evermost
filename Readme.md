###### Guide:
```
https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj
```

###### How to run
```
go run main.go
```

###### soldier
```
http://localhost:8080/soldier/verify
request:
{
	"Data":[[1,1,1,0,0],[1,1,1,1,1],[0,0,0,0,0]]
}

response:
{
    "code": 200,
    "message": "SUCCESS",
    "data": [
        {
            "magazine": "[1 1 1 0 0]",
            "verified": "false"
        },
        {
            "magazine": "[1 1 1 1 1]",
            "verified": "true"
        },
        {
            "magazine": "[0 0 0 0 0]",
            "verified": "false"
        }
    ]
}
```

###### store
```
http://localhost:8080/kitara-store
response:
{
	code: 200,
	message: "SUCCESS",
	data: {
		productId: 1,
		productName: "T-SHIRT A",
		productQuantity: 3
	}
}

http://localhost:8080/kitara-store/request?productId=1&quantity=1
response:
{
	code: 200,
	message: "SUCCESS"
}
```

###### 5 lokasi:
```
[# # # # # # # #]
[# . . . . . . #]
[# . # # # . . #]
[# . . . # . # #]
[# X # . . . . #]
[# # # # # # # #]
1 langkah ke utara, 2 langkah ke timur dan 1 langkah ke selatan
3 langkah ke utara, 4 langkah ke timur dan 1 langkah ke selatan
3 langkah ke utara, 4 langkah ke timur dan 2 langkah ke selatan
3 langkah ke utara, 4 langkah ke timur dan 3 langkah ke selatan
3 langkah ke utara, 5 langkah ke timur dan 1 langkah ke selatan
[# # # # # # # #]
[# . . . . . . #]
[# . # # # K K #]
[# . . . # K # #]
[# X # K . K . #]
[# # # # # # # #]
```