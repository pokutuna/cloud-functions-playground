.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy express-accesslog --entry-point=app --runtime=nodejs8 --trigger-http --allow-unauthenticated
