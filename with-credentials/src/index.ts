import express = require('express');
import * as asyncHandler from 'express-async-handler';

import { Logging } from '@google-cloud/logging';
import { Firestore } from '@google-cloud/firestore';
import { BigQuery } from '@google-cloud/bigquery';

export const app = express();

app.get('/', (req, res) => {
  res.send('ok');
});

const projectId = 'pokutuna-dev';

app.post(
  '/logging',
  asyncHandler(async (req: express.Request, res: express.Response) => {
    const logging = new Logging({ projectId });
    const log = logging.log('test');

    const obj = {
      path: req.path,
    };
    const entry = log.entry(obj);
    console.log(entry.toJSON());
    await log.write(entry);
    return res.json(entry.toJSON());
  })
);

app.post(
  '/firestore',
  asyncHandler(async (req: express.Request, res: express.Response) => {
    const firestore = new Firestore();
    const ref = firestore.collection('with-credentials').doc(req.path);
    const doc = await ref.get();

    const data = {
      count: doc.exists ? (doc.data() as any).count + 1 : 0,
      updatedAt: new Date(),
    };
    console.log(JSON.stringify(data));
    await ref.set(data);

    return res.json(data);
  })
);

app.post(
  '/bigquery',
  asyncHandler(async (req: express.Request, res: express.Response) => {
    const client = new BigQuery();
    const [dataset] = await client.getDatasets();
    const ids = (dataset || []).map(d => d.id)
    console.log(JSON.stringify(ids));
    return res.json(ids);
  })
);
