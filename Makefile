PROJECT_NAME:=wasty
TODAY=`date +"%Y-%m-%d"`

serve:
	go run cmd/server/server.go

build_gene:
	go build -o bin/gene cmd/generate/generate.go

generate:
	bin/gene config/config.yml image
	bin/gene config/config.yml execution
	bin/gene config/config.yml address
	bin/gene config/config.yml waybill
	bin/gene config/config.yml passwd
	bin/gene config/config.yml user
	bin/gene config/config.yml firm
	bin/gene config/config.yml car
	bin/gene config/config.yml driver
	bin/gene config/config.yml page

build: build_gene generate
	go build -o bin/wasty cmd/server/server.go