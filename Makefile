PROJECT_NAME:=wasty
TODAY=`date +"%Y-%m-%d"`

serve:
	go run cmd/server/server.go