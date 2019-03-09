.DEFAULT_GOAL := build
.PHONY: dep build install docker dockerpush

REPO=reale/ipa-cl-ea
LDFLAGS=-ldflags "-X github.com/reale/ipa-cl-ea/store.Sha=`git rev-parse HEAD`"

dep:
	@dep ensure

build: dep
	@go build $(LDFLAGS) -o ipa-cl-ea

install: dep
	@go install $(LDFLAGS)

docker:
	@docker build . -t $(REPO)

dockerpush:
	@docker push $(REPO)
