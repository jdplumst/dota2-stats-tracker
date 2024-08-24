MAKEFLAGS += -j2

run-client:
	@echo "Running client..."
	@cd client && pnpm run dev

run-server:
	@echo "Running server..."
	@cd server && vercel dev ../ 

run: run-client run-server
