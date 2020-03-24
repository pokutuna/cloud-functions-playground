.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy metadata-server \
		--entry-point=app \
		--runtime=nodejs8 \
		--trigger-http
