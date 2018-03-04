# TalenSearching
Talent Search Coding Test

## Testing
To run the unit tests use the following command `go test` 

## Run
To launch the webserver run the following command `./run.sh`

#### Adding Jobs Ads:
**POST** request http://localhost:8080/add 
with JSON body in the following format 

eg. for adding a premium ad
```json
{
	"ItemType": "premium"
}
```

On a successful request the JSON body of the following format will be returned:
```json
{
    "Classic": int,
    "Standout": int,
    "Premium": int
}
```
An example where the customer has (after adding a job) 1 classic ad, 2 standout ads and 3 premium ads. 
```json
{
    "Classic": 1,
    "Standout": 2,
    "Premium": 3
}
```
#### Deleting Job Ads:
**DELETE** request http://localhost:8080/delete?customer={customer}&type={jobtype}

eg. for customer Nike to delete a job ad of type 'standout' you would use the following **DELETE** request http://localhost:8080/delete?customer=nike&type=standout

On a successful request the JSON body of the following format will be returned:
```json
{
    "Classic": int,
    "Standout": int,
    "Premium": int
}
```
An example where the customer has (after deleting a job) 3 classic ad, 2 standout ads and 1 premium ads. 
```json
{
    "Classic": 3,
    "Standout": 2,
    "Premium": 1
}
```

#### Displaying Total:
**GET** request http://localhost:8080/total?customer={customer}

eg. to see customer Ford's total cost you would use the following **GET** request http://localhost:8080/total?customer=ford 

On a successful request the JSON body of the following format will be returned:
```json
{
    "TotalCost": float
}
```
An example where the total cost for the customer is $1,123.80:
```json
{
    "TotalCost": 1123.80
}
```
