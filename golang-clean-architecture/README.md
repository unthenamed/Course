# Aplikasi Enigma Laundry

### Deskripsi
Sebuah api backend untuk aplikasi laundy, dengan postgres sql sebagai databasenya.
---

## Contoh Penggunaan API

### Customer API

#### Create Customer
```curl
curl -X POST '{{endpoint}}/customers' -H 'Content-Type:application/json' -H 'Accept:application/json' -d '{
  "id": 12,
  "name": "Contoh",
  "phoneNumber": "0857-1213-2312",
  "address": "Jawilan"
}'
```
#### Get Customer

```curl
curl -X GET '{{endpoint}}/customers/10' -H 'Accept:application/json'
```

#### Update Customer

```curl
curl -X PUT '{{endpoint}}/customers/12' -H 'Content-Type:application/json' -H 'Accept:application/json' -d '{
	"phoneNumber": "0812-2512-1672"
}'
```


#### Delete Customer

```curl
curl -X DELETE '{{endpoint}}/customers/12' -H 'Accept:application/json'
```


### Employee API

#### Create Employee

```curl
curl -X POST '{{endpoint}}/employees' -H 'Content-Type:application/json' -H 'Accept:application/json' -d '{
  "name": "umi hanif",
  "phoneNumber": "0823-1382-1789",
  "address": "cikande"
}'
```


#### Get Employee

```curl
curl -X GET '{{endpoint}}/employees/11' -H 'Content-Type:application/json' -H 'Accept:application/json'
```


#### Update Employee

```curl
curl -X PUT '{{endpoint}}/employees/13' -H 'Content-Type:application/json' -H 'Accept:application/json' -d '{
	"address": "cikande kibin"
}'
```


#### Delete Employee

```curl
curl -X DELETE '{{endpoint}}/employees/13' -H 'Content-Type:application/json' -H 'Accept:application/json' -d '{
	"address": "cikande kibin"
}'
```


### Product API

#### Create Product

```curl
curl -X POST '{{endpoint}}/products' -H 'Content-Type:application/json' -H 'Accept:application/json' -d '{
	"name": "Cuci Boneka",
  "price": 5000,
  "unit": "1kg" 
}'
```
#### Get Product
```curl
curl -X GET '{{endpoint}}/products/10' -H 'Content-Type:application/json' -H 'Accept:application/json'
```
#### List Product

```curl
curl -X GET '{{endpoint}}/products' -H 'Content-Type:application/json' -H 'Accept:application/json'
```

#### Update Product

```curl
curl -X PUT '{{endpoint}}/products/10' -H 'Content-Type:application/json' -H 'Accept:application/json' -d '{
	"price": 10000
}'
```
#### Delete Product

```curl
curl -X DELETE '{{endpoint}}/products/10' -H 'Content-Type:application/json' -H 'Accept:application/json' -d '{
	"name": "Cuci Boneka",
  "price": 5000,
  "unit": "1kg" 
}'

```


### Transaction API

#### Create Transaction

```curl
curl -X POST '{{endpoint}}/transactions' -H 'Accept:application/json' -H 'Content-Type:application/json' -d '{
	"billDate": "15-12-2024",
	"entryDate": "15-12-2024",
	"finishDate": "20-12-2024",
	"employeeId": 9,
	"customerId": 7,
	"billDetails": [
		{
			"productId": 5,
			"qty": 2
		}
	]
}'
```

#### Get Transaction

```curl
curl -X GET '{{endpoint}}/transactions/50' -H 'Accept:application/json' -H 'Content-Type:application/json'
```

#### List Transaction

Pattern string date : `dd-MM-yyyy`

```curl
curl -X GET '{{endpoint}}/transactions?startDate=15-12-2024&endDate=20-12-2024&productName=Cuci' -H 'Accept:application/json' -H 'Content-Type:application/json'
```




