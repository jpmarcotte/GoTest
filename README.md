# Sample Go Webserver
## A webserver with a single endpoint and configurable listening port

### Summary
This is a sample Go webserver that connects to an SQLite Database.
An in-memory database store is included for local testing where SQLite may not be available,
and to demonstrate how different persistant storage engines can be used.

### Requirements
- Local development will benefit from a recent build of Go and SQLite.
- To run as a proof of concept, all that is required is Docker

### Instructions

- If not already viewing from an archive, clone and move into the Repo:
  ```bash
  git clone git@github.com:jpmarcotte/GoTest.git
  cd GoTest
  ```
- Update `.env` to include the desired port ([51701][ncc-1701] is the default)
- Build and run the docker container
  ```bash
  docker compose build
  docker compose up -d
  ```
- open the link with the appropriate port (change the port if changed from default):
  http://localhost:51701/employees/
- alternatively fetch using `curl`:
  ```bash
  curl http://localhost:51701/employees/
  ```
### Cleanup
- Stop docker
  ```bash
  docker compose down
  ```
- Remove existing docker image
  ```bash
  docker image rm gotest-gotest-app
  ```
  
[ncc-1701]: https://static.wikia.nocookie.net/memoryalpha/images/b/be/USS_Enterprise_%28NCC-1701%29%2C_ENT.jpg/revision/latest?cb=20171022133400&path-prefix=en