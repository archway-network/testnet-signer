#!/usr/bin/make -f

install: go.sum
	go install -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name=archwayd" .