# segmed
Simple web app which displays pictures to the user and saves the ones which are tagged by the user into a DB.

## Endpoints
- `Home`: displays 6 images to the user. Under each image there is `Tag` button which marks an image as tagged in the database. 
- `Show table`: after tagging the images, the user can see which of the images where tagged.
- `Contact`: generic contact page.    

## How to run

### Clone the repo
```bash
git clone https://github.com/Radu1990/segmed.git
```

### Pulling dependencies
```bash
go get github.com/mattn/go-sqlite3
go get "github.com/gorilla/mux"
```

### Runing locally
```bash
go run main.go
```
### Running tests
```bash
go test ./... -v
```

App is running on `localhost:3000`