export PROJECT := pokutuna-dev
export GCLOUD := gcloud beta --project $(PROJECT)

SUBDIRS := $(subst /.,,$(wildcard */.))

.PHONY: $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@ deploy
