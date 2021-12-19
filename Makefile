build: 
	docker build -t j-thompson12/serverless-api .

compile: 
	docker run -e GOOS=linux -e GOARCH=amd64 -v $$(pwd):/app -w /app golang:1.13 go build -ldflags="-s -w" -o bin/serverless-api

plan: compile
	docker run -v $$(pwd)/main.tf:/srv/main.tf -v $$(pwd)/terraform.tfstate:/srv/terraform.tfstate -v $$(pwd)/bin:/srv/bin -e AWS_ACCESS_KEY_ID=AKIAZ4JUGFUIWTPKFNFG -e AWS_SECRET_ACCESS_KEY=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 -e AWS_DEFAULT_REGION=us-west-2 j-thompson12/terraform-aws plan

apply: compile
	docker run -v $$(pwd)/main.tf:/srv/main.tf -v $$(pwd)/terraform.tfstate:/srv/terraform.tfstate -v $$(pwd)/bin:/srv/bin -e AWS_ACCESS_KEY_ID=AKIAZ4JUGFUIWTPKFNFG -e AWS_SECRET_ACCESS_KEY=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 -e AWS_DEFAULT_REGION=us-west-2 j-thompson12/terraform-aws apply -auto-approve

destroy:
	docker run -v $$(pwd)/main.tf:/srv/main.tf -v $$(pwd)/terraform.tfstate:/srv/terraform.tfstate -v $$(pwd)/bin:/srv/bin -e AWS_ACCESS_KEY_ID=AKIAZ4JUGFUIWTPKFNFG -e AWS_SECRET_ACCESS_KEY=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 -e AWS_DEFAULT_REGION=us-west-2 j-thompson12/terraform-aws destroy -auto-approve
