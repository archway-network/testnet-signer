#!/usr/bin/make -f

install:
	go install -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name=archwayd" .