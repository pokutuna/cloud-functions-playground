.PHONY: deploy
deploy:
	$(GCLOUD) functions deploy envs-node16-gen2 \
		--gen2 \
		--runtime=nodejs16 \
		--entry-point=gen2 \
		--trigger-http \
		--allow-unauthenticated \
		--region=asia-northeast1 \
		--serve-all-traffic-latest-revision
