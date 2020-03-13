.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy state-and-error --entry-point=app --runtime nodejs10 --trigger-http --allow-unauthenticated
