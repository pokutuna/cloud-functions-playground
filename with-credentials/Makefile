SERVICE_ACCOUNT1 := credentials-empty-roles
SERVICE_ACCOUNT2 := credentials-account

.PHONY: deploy
deploy:
	yarn compile

	$(GCLOUD) functions deploy with-credentials-default \
		--runtime=nodejs8 \
		--entry-point=app \
		--trigger-http \
		--allow-unauthenticated

	$(GCLOUD) functions deploy with-credentials-empty-roles \
		--runtime=nodejs8 \
		--entry-point=app \
		--trigger-http \
		--allow-unauthenticated \
		--set-env-vars="GOOGLE_APPLICATION_CREDENTIALS=./$(SERVICE_ACCOUNT1)-key.json"

	$(GCLOUD) functions deploy with-credentials \
		--runtime=nodejs8 \
		--entry-point=app \
		--trigger-http \
		--allow-unauthenticated \
		--set-env-vars="GOOGLE_APPLICATION_CREDENTIALS=./$(SERVICE_ACCOUNT2)-key.json"


.PHONY: create-service-account
create-service-accounts:
	$(GCLOUD) iam service-accounts create $(SERVICE_ACCOUNT1)
	$(GCLOUD)	iam service-accounts create $(SERVICE_ACCOUNT2)
	$(GCLOUD) projects add-iam-policy-binding $(PROJECT) \
		--member=serviceAccount:$(SERVICE_ACCOUNT2)@$(PROJECT).iam.gserviceaccount.com \
		--role=roles/logging.logWriter
		--role=roles/datastore.user

create-service-account-keys:
	$(GCLOUD) iam service-accounts keys create $(SERVICE_ACCOUNT1)-key.json \
		--iam-account=$(SERVICE_ACCOUNT1)@$(PROJECT).iam.gserviceaccount.com
	$(GCLOUD) iam service-accounts keys create $(SERVICE_ACCOUNT2)-key.json \
		--iam-account=$(SERVICE_ACCOUNT2)@$(PROJECT).iam.gserviceaccount.com


.PHONY: request-default
request-default:
	$(MAKE) request FUNCTION=with-credentials-default

.PHONY: request-empty-roles
request-empty-roles:
	$(MAKE) request FUNCTION=with-credentials-empty-roles

.PHONY: request-with-credentials
request-with-credentials:
	$(MAKE) request FUNCTION=with-credentials

.PHONY: request
request:
	curl -X POST https://us-central1-${PROJECT}.cloudfunctions.net/${FUNCTION}/logging
	sleep 1
	curl -X POST https://us-central1-${PROJECT}.cloudfunctions.net/${FUNCTION}/firestore
	sleep 1
	curl -X POST https://us-central1-${PROJECT}.cloudfunctions.net/${FUNCTION}/bigquery
