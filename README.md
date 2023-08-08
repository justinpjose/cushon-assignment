# Cushon Assignment
Implementation of the assignment given by Cushon

## Requirements
The requirements of the assignment can be found [here](docs/assignment.pdf)

## Design
Documentation and Diagrams for the design of the application can be found below:
- [Data Entities](/docs/data.md)
- [Infrastructure](docs/infrastructure.md)
- [Sequence](internal/handlers/createtransaction/post_transaction_sequence.drawio.png)
- [API](api/README.md)

## Prerequisite
Ensure that you have the following installed:
- [Docker](https://www.docker.com/)

## Usage
### Commands
1. `make check` - checks the repo for any vulnerability, lint issues and runs the unit tests
2. `make start` - runs the docker containers for the db and api in the background
3. `make stop` - stops the db and removes the docker container
4. `make start-local` - runs DB in docker container while the API runs locally on your host machine
5. `make logs-api` to get logs from the api docker container

### Instructions

#### Option 1 - Main
1. Run `make start` to run API and DB (both will be running in their own separate docker container)
2. Make requests to the API via Postman using the [Postman Collection](api/README.md#postman-collection) or curl commands in a separate shell. Endpoints available can be found in the [swagger spec](api/README.md#viewing-the-swagger-spec)
```
curl -X POST http://localhost:8080/v0/transactions -H "Content-Type: application/json" -d '{"customer_accounts_funds_id": 1, "amount": 25000}'
```
3. Check logs in the container to get information about the requests being made
- Run `make logs-api` to get logs from the api container
4. Connect to the DB locally via a database administration software - e.g. DBBeaver using the configurations at the top of the [Makefile](Makefile) to see the data itself in the database
6. Run `make stop` to stop the DB and API

#### Option 2 - Alternative
Run `make start-local` to run DB in docker container while the API runs locally on your host machine
- This option is useful if you want to see the logs of the API on your shell
- Remember to `Ctrl+C` to stop the API
- Remember to `make stop` to stop the DB container

## Scripts
Informtion about scripts can be found [here](scripts/README.md)

## Assumptions
- Each Company has specific set of funds that the employers can choose from
- Employer customer data and direct customer data are the same 
- We can purge data in our system otherwise a status needs to be added to the data entities where the status would become 'Deleted' even though the data still exists in the DB
- Each account has specific funds and those funds do not share with other account types
- Other assumptions can be found in the other READMEs located in the repo which are related to the topics discussed in that README

## Enchanments
- Run tests in parallel while being mindful of data race - https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/loopclosure
- Sensitive data (e.g. DB_PASS) should be retrieved from a vault or secrets managers
- Create sequence diagrams for the other handlers
- Add type to transactions to distinguish from personal, interest, etc
- Add profit/loss in investments so user can see a summary of how much they have gained
- Handling different types of currency
- Other endpoints can be added for more functionality
- Making backups of data
- Add more tests
- Add integration test for postgres db functions
- Add integration test at api-level but use BDD-style test
- Other enchanments can be found in the other READMEs located in the repo which are related to the topics discussed in that README
- Update db tables and models to add all the fields mentioned in the class diagram

## Notes
- Project Structuring was based on https://github.com/golang-standards/project-layout to follow idiomatic Go patterns

[//]: # (Reference Links)
[different-tests]: <https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/#TheTestifyPackage>
[testing-libraries]: <https://speedscale.com/blog/golang-testing-frameworks/#elementor-toc__heading-anchor-8>