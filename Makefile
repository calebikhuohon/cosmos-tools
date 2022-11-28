lint:
	golangci-lint run

install:
	go install ./cmd/validator-status && go install ./cmd/vesting-account
