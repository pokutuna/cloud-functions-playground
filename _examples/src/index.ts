import { CloudBuildClient } from '@google-cloud/cloudbuild';
import { google } from '@google-cloud/cloudbuild/build/protos/protos';

const projectId = 'pokutuna-playground';

const build = {
  steps: [
    {
      name: 'alpine:latest',
      entrypoint: '/bin/echo',
      args: ['start'],
    },
    {
      name: 'gcr.io/cloud-builders/gcloud',
      args: [
        'functions',
        'deploy',
        'hello',
        '--runtime=nodejs10',
        '--entry-point=app',
        '--trigger-http',
        `--source=gs://${projectId}/gomi/function.zip`,
      ],
    },
  ],
  timeout: { seconds: 600 },
  // timeout: '600s',
};

(async () => {
  const client = new CloudBuildClient();
  const res = await client.createBuild({
    projectId,
    build,
  });
  console.log(res);
})();
