# Booking System

Repo for a booking system Built in Go


External Dependencies:
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [alex edwards scs session management](https://github.com/alexedwards/scs)
- Uses [no-surf](https://github.com/justinas/nosurf)

Dev Setup:
- Using https://postgresapp.com/ for database (app variant)
- Soda/Buffalo framework for migrations & DB managment (https://gobuffalo.io/documentation/database/pop/)
  Note installed via homebrew: brew install gobuffalo/tap/pop

Run:
Created ./run.sh

Test Notes:
- Coverage alias set up in local .zshrc 
  - alias coverage="go test -coverprofile=coverage.out && go tool cover -html=coverage.out"
- Run all tests from root:
  - go test -v ./...



Soda:
- Initialise DB from migrations
  - soda migrate
  - soda migrate down
