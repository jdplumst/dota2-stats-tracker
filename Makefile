MAKEFLAGS += -j2

run-client:
	@echo "Running client..."
	@cd client && pnpm run dev

run-server:
	@echo "Running server..."
	@cd server && vercel dev ../ 

run: run-client run-server

install-server:
	@echo "Installing server dependencies..."
	@cd server && go get ./...

test-server:
	@echo "Testing server..."
	@cd server && go test ./api -v