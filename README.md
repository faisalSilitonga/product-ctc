# Architecture

Using Clean Architecture in Go [Repo](https://github.com/eminetto/clean-architecture-go-v2)

Based on [Post: Clean Architecture, 2 years later](https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/)

Clean Architecture based on => [Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)


# API Documentation

API Documentation in [Postman](https://documenter.getpostman.com/view/13179409/TVes7mX8#07cc5871-16a1-475d-ad43-2ce13ec31d73)

## Build

  make

## Run tests

  make test

## Run tests

  make run

  
## API requests 

### Add Product

```
curl -X "POST" "http://localhost:8080/v1/product" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
		"name": "Product Test 1",
		"sku": "TEST-01",
		"catalogid": 294,
		"price":100,
		"quantity":1
		}'
```
### Search Product

```
curl "http://localhost:8080/v1/product?name=test" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Get Product by ID

```
curl "http://localhost:8080/v1/product/{id}" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```


### Show Product

```
curl "http://localhost:8080/v1/product" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```


### Update Product

```
curl -X "POST" "http://localhost:8080/v1/product/update" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
		"id": {id},
		"name": "Product Test 1",
		"sku": "TEST-01",
		"catalogid": 294,
		"price":100,
		"quantity":1,
		"status":1
		}'
```

### Delete Product

```
curl -X "DELETE" "http://localhost:8080/v1/product/{id}" \
     -H 'Content-Type: application/json'
```

