db=minpham-family-db

.PHONY: dev
dev:
	wrangler dev

.PHONY: migration-create
migration-create:
	wrangler d1 migrations create $(db) $(args)

.PHONY: migration-apply
migration-apply:
	wrangler d1 migrations apply $(db) $(args)

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.18.0
	tinygo build -o ./build/app.wasm -target wasm -no-debug ./...

.PHONY: deploy
deploy:
	wrangler deploy

.PHONY: schema-seed
schema-seed:
	wrangler d1 execute $(db) --file=./seed.sql
