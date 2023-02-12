.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy bq-remote-fetch \
		--runtime=go119 \
		--entry-point=RemoteFunction \
		--region=asia-northeast1 \
		--trigger-http \
		--allow-unauthenticated
