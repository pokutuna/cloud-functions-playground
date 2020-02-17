.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy write-json-stdout \
	  --runtime=nodejs10 \
	  --entry-point=app \
	  --trigger-http
