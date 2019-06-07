all:

build:
	go build -v -o ./bin/subscriptions ./cmd/subscriptions/...

install:
	go install ./cmd/subscriptions/...

run:
	go run ./cmd/subscriptions/...

tests t:
	go test -v ./internal/...

lint-all l:
	bash ./scripts/golangci-lint.sh
	bash ./scripts/consistent.sh

add-ssh-key:
	openssl aes-256-cbc -K $(encrypted_694e6c89a81b_key) -iv $(encrypted_694e6c89a81b_iv) -in travis_key.enc -out /tmp/travis_key -d
	chmod 600 /tmp/travis_key
	ssh-add /tmp/travis_key

docker-login:
	docker login -u $(REGISTRY_USER) -p $(REGISTRY_PASS)

docker-build:
	docker build -t $(REGISTRY_REPO):$(VERSION) .

docker-push: docker-login
	docker push $(REGISTRY_REPO):$(VERSION)

deploy:
	ssh -o "StrictHostKeyChecking no" $(HOST_USER)@$(HOST) make run-music-subscriptions

deploy-staging:
	ssh -o "StrictHostKeyChecking no" $(STAGING_USER)@$(STAGING_HOST) make run-music-subscriptions
