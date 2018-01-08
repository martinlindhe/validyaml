update-vendor:
	dep ensure
	dep ensure -update
	dep prune

release:
	goreleaser --rm-dist
