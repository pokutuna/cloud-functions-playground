
.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy params-limit --runtime=nodejs16 --entry-point=app --trigger-http --allow-unauthenticated
