NUM_DEVICES=2
DURATION=10
INTERVAL=3

run-devices:
	GATEWAY_HOST=localhost GATEWAY_PORT=28000 go run ./*.go -num-devices=${NUM_DEVICES} -duration=${DURATION} -interval=${INTERVAL}

install:
	go get ./