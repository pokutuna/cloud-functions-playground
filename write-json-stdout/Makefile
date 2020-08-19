.PHONY: deploy
deploy:
	$(MAKE) deploy-function RUNTIME=nodejs8 SOURCE=./node
	$(MAKE) deploy-function RUNTIME=nodejs10 SOURCE=./node
	$(MAKE) deploy-function RUNTIME=nodejs12 SOURCE=./node
	$(MAKE) deploy-function RUNTIME=go111 SOURCE=./golang
	$(MAKE) deploy-function RUNTIME=go113 SOURCE=./golang
	$(MAKE) deploy-function RUNTIME=python37 SOURCE=./python
	$(MAKE) deploy-function RUNTIME=python38 SOURCE=./python
	$(MAKE) deploy-function-java RUNTIME=java11 SOURCE=./java

.PHONY: deploy-function
deploy-function:
	$(GCLOUD) functions deploy write-json-stdout-$(RUNTIME) \
		--runtime=$(RUNTIME) \
		--source=$(SOURCE) \
		--entry-point=App \
		--trigger-http \
		--allow-unauthenticated \
		--set-env-vars=GCF_RUNTIME=$(RUNTIME)

.PHONY: deploy-function-java
deploy-function-java:
	# --entry-point requires {package}.{Class}
	$(GCLOUD) functions deploy write-json-stdout-$(RUNTIME) \
		--runtime=$(RUNTIME) \
		--source=$(SOURCE) \
		--entry-point=functions.App \
		--trigger-http \
		--allow-unauthenticated \
		--set-env-vars=GCF_RUNTIME=$(RUNTIME)
