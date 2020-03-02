# $ make {directory}
# $ make {directory} TASK={task}

export PROJECT := pokutuna-dev
export GCLOUD := gcloud beta --project $(PROJECT)

TASK := deploy
SUBDIRS := $(subst /.,,$(wildcard */.))

.PHONY: $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@ ${TASK}

.PHONY: generate-service-account-key
generate-service-account-key:
	$(GCLOUD) iam service-accounts keys create _credentials/key.json \
		--iam-account=$(PROJECT)@appspot.gserviceaccount.com
