.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy spreadsheet --entry-point=spreadsheet --runtime nodejs16 --trigger-http --allow-unauthenticated
