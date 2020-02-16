.PHONY: deploy
deploy:
	$(MAKE) deploy-nodejs

.PHONY: deploy-nodejs
deploy-nodejs:
	yarn --cwd=nodejs compile
	$(GCLOUD) functions deploy verify-jwt-pubsub-nodejs \
		--runtime=nodejs10 \
		--entry-point=app \
		--source=./nodejs \
		--trigger-http

.PHONY: deploy-golang
deploy-golang:
	$(GCLOUD) functions deploy verify-jwt-pubsub-golang \
		--runtime=go113 \
		--entry-point=App \
		--source=./golang \
		--trigger-http


TOPIC := pubsub-verify-example
SUBSCRIPTION := pubsub-verify-example
SERVICE_ACCOUNT := pubsub-verification-example

.PHONY: create-topic
create-topic:
	$(GCLOUD) pubsub topics create $(TOPIC)

.PHONY: create-subscriptions
create-subscriptions:

	$(GCLOUD) pubsub subscriptions create $(SUBSCRIPTION)-nodejs-google-auth-library \
		--topic=$(TOPIC) \
		--push-endpoint="https://us-central1-pokutuna-dev.cloudfunctions.net/verify-jwt-pubsub-nodejs/google-auth-library" \
		--push-auth-service-account="$(SERVICE_ACCOUNT)@$(PROJECT).iam.gserviceaccount.com" \
		--push-auth-token-audience="verify-jwt-pubsub"

	$(GCLOUD) pubsub subscriptions create $(SUBSCRIPTION)-nodejs-jsonwebtoken \
		--topic=$(TOPIC) \
		--push-endpoint="https://us-central1-pokutuna-dev.cloudfunctions.net/verify-jwt-pubsub-nodejs/jsonwebtoken" \
		--push-auth-service-account="$(SERVICE_ACCOUNT)@$(PROJECT).iam.gserviceaccount.com" \
		--push-auth-token-audience="verify-jwt-pubsub"


.PHONY: publish-message
publish-message:
	$(GCLOUD) pubsub topics publish $(TOPIC) \
    --message='{ "key": "value" }'

.PHONY: create-service-account
create-service-account:
	$(GCLOUD) iam service-accounts create $(SERVICE_ACCOUNT)
