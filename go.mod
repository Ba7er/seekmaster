module github/Ba7er/seekmaster

go 1.22.0

replace github/Ba7er/seekmaster/internals/http => ../internals/http

require github.com/go-sql-driver/mysql v1.8.1

require filippo.io/edwards25519 v1.1.0 // indirect

replace github/Ba7er/seekmaster/internals/services => ./internals/services
