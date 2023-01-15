dev:
	/home/pablo/go/bin/air -c .air.toml

build:
	go build -o ./dist/app .

swagger:
	/home/pablo/go/bin/swag fmt && /home/pablo/go/bin/swag init