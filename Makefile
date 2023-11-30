CONFIG = config/config-docker.yml
CONFIG_LOCAL = config/config-local.yml


local:
			@(echo "Check imports to download")
			@go mod tidy
			@(echo "Launching app")
			@go run cmd/main.go

config:
			@(echo "Creating configs for launch. Don't forget change sample credentials.")
			@(cp ./config/config-sample.yml $(CONFIG_LOCAL))


.PHONY: local config