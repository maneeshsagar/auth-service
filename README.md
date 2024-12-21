# Token Service





### Gin, Mysql, Migration, Docker and docker-compose is used to complete this project



### Steps to run the project

clone the Repository
```bash
 [git clone github.com/maneeshsagar/auth-service](https://github.com/maneeshsagar/auth-service.git)
```

run the docker compose file by below command
```bash
cd auth-service
docker-compose up --build -d
```


#### Test the APIs

1. Sign Up

```bash
curl --location 'localhost:8080/auth/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"Maneesh Sagar",
    "email":"aks@abc.com",
    "password":"ps"
}'
```
2. Sign In to get the tokens

``` bash
curl --location 'localhost:8080/auth/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"aks@abc.com",
    "password":"ps"
}'
```

3. Get the Profile of the user using access token (add the access token which you received in the signin repsonse)

```bash

curl --location 'localhost:8080/v1/profile' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzQ3NzQ2NDEsImlhdCI6MTczNDc3Mzc0MX0.sitjjQdhq4_dmWV2IJnP1s7AACcj2j-Ha0CYi1YHgpg'

```

4. If access token is expired then use this one to get the new token based on the refresh token

``` bash
curl --location 'localhost:8080/auth/refresh' \
--header 'Content-Type: application/json' \
--data '{
     "refreshToken": "GOvnii0sVw9tWx5ba4zdQO7ze6d5lleJ"
}'

```

#### Test with postman

import the Auth-Service.postman_collection.json file from the auth-service folder into the postman

and hit the apis
