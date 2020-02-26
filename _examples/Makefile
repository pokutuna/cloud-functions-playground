PROJECT := pokutuna-playground
GCLOUD := gcloud --project=$(PROJECT)
SERVICE_ACCOUNT := builder

# サービスアカウントキーを作る
.PHONY: generate-service-account-key
generate-service-account-key:
	-$(GCLOUD) iam service-accounts create $(SERVICE_ACCOUNT)
	-$(GCLOUD) projects add-iam-policy-binding $(PROJECT) \
		--member=serviceAccount:$(SERVICE_ACCOUNT)@$(PROJECT).iam.gserviceaccount.com \
		--role=roles/cloudbuild.builds.editor
	$(GCLOUD) iam service-accounts keys create keys/key.json \
		--iam-account=$(SERVICE_ACCOUNT)@$(PROJECT).iam.gserviceaccount.com

# デプロイ対象として function を GCS にアップロードしておく
.PHONY: upload-function
upload-function:
	zip -rj function.zip function
	gsutil cp function.zip gs://$(PROJECT)/gomi/function.zip

# Cloud Build サービスアカウントに、Cloud Functions へのデプロイ権限をあげる
.PHONY: add-function-deployer-policy
add-function-deployer-policy:
	$(eval PROJECT_NUMBER = $(shell $(GCLOUD) projects list --filter=$(PROJECT) --format="value(PROJECT_NUMBER)"))
	$(GCLOUD) projects add-iam-policy-binding $(PROJECT) \
		--member=serviceAccount:$(PROJECT_NUMBER)@cloudbuild.gserviceaccount.com \
		--role=roles/cloudfunctions.developer \
		--role=roles/iam.serviceAccountUser
