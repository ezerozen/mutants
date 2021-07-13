# mutants

- Install [Golang](https://golang.org/dl/)
- Install [Docker](https://www.docker.com/get-started)
- Run
````
docker-compose up
````

- Endpoints
````
curl --location --request POST 'localhost:8080/mutant' \
--header 'Content-Type: application/json' \
--data-raw '{
"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}'
````
````
curl --location --request GET 'http://localhost:8080/stats'
````