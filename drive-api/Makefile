.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy drive-api \
		--runtime=nodejs16 \
		--trigger-http \
		--allow-unauthenticated \
		--region=asia-northeast1
