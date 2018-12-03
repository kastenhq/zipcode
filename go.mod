module github.com/kastenhq/zipcode

replace github.com/kastenhq/zipcode/cmd/zipcode => ./cmd/zipcode

replace github.com/kastenhq/zipcode/pkg/zipcode => ./pkg/zipcode

require (
	github.com/kastenhq/zipcode/cmd/zipcode v0.0.0-20181203215836-7bedf43fb22f // indirect
	github.com/kastenhq/zipcode/pkg/zipcode v0.0.0-20181203215836-7bedf43fb22f // indirect
	github.com/lib/pq v1.0.0 // indirect
)
