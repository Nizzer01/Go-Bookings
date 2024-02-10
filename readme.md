# Booking System

Repo for a booking system Built in Go


External Dependencies:
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [alex edwards scs session management](https://github.com/alexedwards/scs)
- Uses [no-surf](https://github.com/justinas/nosurf)
- Uses go-simple-mail (https://github.com/xhit/go-simple-mail)

Dev Setup:
- Using https://postgresapp.com/ for database (app variant)
- Soda/Buffalo framework for migrations & DB management (https://gobuffalo.io/documentation/database/pop/)
  Note installed via homebrew: brew install gobuffalo/tap/pop
- Use MailHog for local email testing 
- Usi Foundation for Emails to niceify our emails, (https://get.foundation/emails/getting-started.html)

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
  - soda reset (runs down migrations then the up migrations) May need to disconnect datagrip from DB first
