import jwt = require('jsonwebtoken');
import jwksClient = require('jwks-rsa');

import createError = require('http-errors');

import { Request, Response, NextFunction } from 'express';

interface Options {
  email: string;
  audience: string;
}

export function verify(options: Options) {
  const certClient = jwksClient({
    jwksUri: 'https://www.googleapis.com/oauth2/v3/certs',
  });
  const getKey = (header: jwt.JwtHeader, callback: jwt.SigningKeyCallback) => {
    certClient.getSigningKey(
      header.kid || '',
      (err: Error | null, key: jwksClient.SigningKey) =>
        err ? callback(err, undefined) : callback(null, key.getPublicKey())
    );
  };

  return (req: Request, res: Response, next: NextFunction) => {
    const bearer = req.header('Authorization') || '';
    const [, token] = bearer.split('Bearer ');
    jwt.verify(
      token,
      getKey,
      {
        audience: options.audience,
        complete: true,
      },
      (err?: Error, decoded?: any) => {
        (req as any).claims = decoded;

        if (err || options.email !== decoded?.payload?.['email']) {
          const error =
            err || new Error(`Unexpected email ${decoded?.payload?.['email']}`);
          next(createError(403, error));
        }

        next();
      }
    );
  };
}
