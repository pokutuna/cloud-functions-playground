import express = require('express');
import { logHttpRequest } from '@pokutuna/requestlog-cloudfunctions';

import { verify as verifyByGoogleAuthLibrary } from './google-auth-library';
import { verify as verifyByJsonWebToken } from './jsonwebtoken';

import { Request, Response } from 'express';
import { PubsubMessage } from '@google-cloud/pubsub/build/src/publisher';

const projectId = 'pokutuna-dev';

export const app = express();
app.use(express.json());
app.use(logHttpRequest({ projectId }));

app.get('/', (req, res) => res.send('ok'));

const verfiyOptions = {
  email: 'pubsub-verification-example@pokutuna-dev.iam.gserviceaccount.com',
  audience: 'verify-jwt-pubsub',
};

const handler = (req: Request, res: Response) => {
  const message: PubsubMessage = req.body.message;
  const data = Buffer.from(
    (message.data || '').toString(),
    'base64'
  ).toString();

  const info = {
    path: req.path,
    claims: (req as any).claims,
    data,
  };

  console.log(JSON.stringify(info));
  res.send(info);
};

app.post(
  '/google-auth-library',
  verifyByGoogleAuthLibrary(verfiyOptions),
  handler
);

app.post('/jsonwebtoken', verifyByJsonWebToken(verfiyOptions), handler);
