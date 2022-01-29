```bash
# coverage report
go test ./... -coverprofile="coverage.out"

# show coverage report
go tool cover -html=coverage

# covarge show
go test ./... --cover
```