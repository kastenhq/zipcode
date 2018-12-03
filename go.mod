module github.com/kastenhq/zipcode

require (
	github.com/kastenhq/zipcode/cmd/zipcode v0.0.0
	github.com/kastenhq/zipcode/pkg/zipcode v0.0.0
	github.com/lib/pq v1.0.0 // indirect
)

replace github.com/kastenhq/zipcode/cmd/zipcode => ./cmd/zipcode

replace github.com/kastenhq/zipcode/pkg/zipcode => ./pkg/zipcode
