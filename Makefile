OUTPUT = main # will be archived
SERVICE_NAME = s3Consumer
PACKAGED_TEMPLATE = packaged.yaml # will be archived
TEMPLATE = template.yaml
VERSION = 0.1
STACK_NAME = $(SERVICE_NAME)stack

.PHONY: test
test:
	go test ./...

clean:
	rm -f $(OUTPUT)
	rm -f $(PACKAGED_TEMPLATE)

main: ./cmd/$(SERVICE_NAME)-lambda/main.go
	go build -o $(OUTPUT) ./cmd/$(SERVICE_NAME)-lambda/main.go
	
build-local:
	go build -o $(OUTPUT) ./cmd/$(SERVICE_NAME)/main.go

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) main

# create a lambda deployment package
$(ZIPFILE): clean lambda
	zip -9 -r $(ZIPFILE) $(OUTPUT)

.PHONY: build
build: clean lambda

.PHONY: package
package:
		aws s3 cp open-api-integrated.yaml s3://uchenchostorage/open-api/$(SERVICE_NAME)/open-api-integrated.yaml
		aws cloudformation package --template-file $(TEMPLATE) --s3-bucket uchenchostorage --output-template-file $(PACKAGED_TEMPLATE) --capabilities CAPABILITY_IAM
		aws cloudformation deploy --template-file $(PACKAGED_TEMPLATE) --stack-name $(STACK_NAME) --capabilities CAPABILITY_IAM

.PHONY: application
application: test build package

run: build-local
	@echo ">> Running application ..."
	PORT=6000 \
	MONGO_URL="mongodb://localhost:27017" \
	MONGO_DB_NAME=stGerald \
	./$(OUTPUT)
