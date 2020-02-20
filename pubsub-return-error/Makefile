TOPIC := pubsub-error

.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy pubsub-return-error \
		--runtime=nodejs10 \
		--entry-point=app \
		--trigger-topic=$(TOPIC) \
		--retry

.PHONY: create-topic
create-topic:
	$(GCLOUD) pubsub topics create $(TOPIC)


.PHONY: publish-ok
publish-ok:
	$(GCLOUD) pubsub topics publish $(TOPIC) --message='{ "key": "value" }'

.PHONY: publish-throw
publish-throw:
	$(GCLOUD) pubsub topics publish $(TOPIC) --message='{ "error": true, "throw": true }'

.PHONY: publish-return
publish-return:
	$(GCLOUD) pubsub topics publish $(TOPIC) --message='{ "error": true, "throw": false }'
