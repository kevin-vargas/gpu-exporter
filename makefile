build:
	GOOS=windows GOARCH=amd64 go build -o bin/gpu-exporter.exe cmd/main.go
