NUM_DEVICES=1
DURATION=7
INTERVAL=2

run-devices:
	GATEWAY_HOST=localhost GATEWAY_PORT=28000 go run ./*.go -num-devices=${NUM_DEVICES} -duration=${DURATION} -interval=${INTERVAL}

install:
	go get ./