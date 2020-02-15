# $ make {directory}
# $ make {directory} TASK={task}

export PROJECT := pokutuna-dev
export GCLOUD := gcloud beta --project $(PROJECT)

TASK := deploy
SUBDIRS := $(subst /.,,$(wildcard */.))

.PHONY: $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@ ${TASK}
