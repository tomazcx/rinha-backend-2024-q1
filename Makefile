help: ## Prints available commands
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[.a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

go: down ## Run the containers
	docker-compose up

go.build: down ## Build and run the containers
	docker-compose up --build

dev: down ## Run in development mode
	docker-compose -f docker-compose.dev.yml up

dev.build: down ## Build and run the containers in development mode
	docker-compose -f docker-compose.dev.yml up --build

down:  ## Remove containers
	docker-compose down --remove-orphans -v

destroy: ## Remove containers, images and volumes
	docker-compose down --rmi all -v --remove-orphans 
