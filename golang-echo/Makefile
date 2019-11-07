.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy golang-echo --runtime=go111 --entry-point=App --trigger-http --allow-unauthenticated
