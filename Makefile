run:
	go run ./cmd//eaglebank
api-build:
	docker build -t bankapi .
api-run:
	docker run --name bankapi --network=host bankapi
api-start:
	docker start bankapi
api-stop:
	docker stop bankapi
run-test:
	ginkgo -v  ./...
gen-mocks:
	go generate ./...
