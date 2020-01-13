.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy stackdriver-trace-node --entry-point=app --runtime nodejs10 --trigger-http --allow-unauthenticated

# project 名をハードコードしていていまいち
.PHONY: create-service-account
create-service-account:
	$(GCLOUD) iam service-accounts create stackdriver-trace-agent
	$(GCLOUD) projects add-iam-policy-binding pokutuna-dev \
		--member=serviceAccount:stackdriver-trace-agent@pokutuna-dev.iam.gserviceaccount.com \
		--role=roles/cloudtrace.agent

.PHONY: create-service-account-key
create-service-account-key:
	$(GCLOUD) iam service-accounts keys create key.json \
		--iam-account stackdriver-trace-agent@pokutuna-dev.iam.gserviceaccount.com
