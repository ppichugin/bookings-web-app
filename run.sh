#!/usr/bin/env bash
go build -o build/bookings ./cmd/web/. || exit 0
./build/bookings
