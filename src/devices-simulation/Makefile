NUM_DEVICES=2
DURATION=6
INTERVAL=2

run-devices:
	GATEWAY_HOST=localhost GATEWAY_PORT=28000 go run ./*.go -num-devices=${NUM_DEVICES} -duration=${DURATION} -interval=${INTERVAL}

install:
	go get ./