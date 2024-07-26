.phony: build docs

dev-docs:
	@source /opt/homebrew/opt/nvm/nvm.sh && nvm use && npm run docs:dev

