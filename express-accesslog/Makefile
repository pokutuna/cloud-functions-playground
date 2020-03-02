.PHONY: deploy
deploy: key.json
	$(GCLOUD) functions deploy express-accesslog \
		--entry-point=app \
		--runtime=nodejs10 \
		--trigger-http \
		--allow-unauthenticated \
		--set-env-vars=GOOGLE_APPLICATION_CREDENTIALS=./key.json

key.json:
	cp ../_credentials/key.json .
