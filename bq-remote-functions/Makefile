.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy bq-remote-functions \
		--runtime=go119 \
		--trigger-http \
		--allow-unauthenticated
