#!/usr/bin/env bash

go run ./cmd/go-compose/ build -target android ./cmd/demo/kitchen/
adb uninstall org.gioui.experiment
adb install GioApp.apk