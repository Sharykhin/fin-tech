.PHONY: serve

serve:
	WEB_API_ADDRESS=:8004 go run -race cmd/api/main.go

migrations:
	echo 12