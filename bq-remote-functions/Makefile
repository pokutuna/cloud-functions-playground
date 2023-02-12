.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy bq-remote-functions \
		--runtime=go119 \
		--entry-point=RemoteFunction \
		--region=asia-northeast1 \
		--trigger-http \
		--allow-unauthenticated
