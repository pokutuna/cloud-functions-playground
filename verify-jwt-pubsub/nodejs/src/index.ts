import express = require('express');
import { logHttpRequest } from '@pokutuna/requestlog-cloudfunctions';

import { verify as verifyByGoogleAuthLibrary } from './google-auth-library';

import { Request, Response } from 'express';
import { PubsubMessage } from '@google-cloud/pubsub/build/src/publisher';

const projectId = 'pokutuna-dev';

export const app = express();
app.use(express.json());
app.use(logHttpRequest({ projectId }));

app.get('/', (req, res) => res.send('ok'));

const email =
  'pubsub-verification-example@pokutuna-dev.iam.gserviceaccount.com';
const audience = 'verify-jwt-pubsub';

app.post(
  '/google-auth-library',
  verifyByGoogleAuthLibrary({ email, audience }),
  (req: Request, res: Response) => {
    const message: PubsubMessage = req.body.message;
    const data = Buffer.from(
      (message.data || '').toString(),
      'base64'
    ).toString();

    console.log(JSON.stringify((req as any).claims));
    res.send(data);
  }
);

app.post('/jsonwebtoken')
