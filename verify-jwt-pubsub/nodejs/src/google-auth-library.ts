import { OAuth2Client, LoginTicket, TokenPayload } from 'google-auth-library';
import createError = require('http-errors');
import { Request, Response, NextFunction } from 'express';

export interface RequestWithClaims extends Request {
  claims: TokenPayload | undefined;
}

interface Options {
  email: string;
  audience: string;
}

// https://cloud.google.com/pubsub/docs/push#authentication_and_authorization_by_the_push_endpoint
export function verify(options: Options) {
  const authClient = new OAuth2Client();

  return (req: Request, res: Response, next: NextFunction) => {
    const bearer = req.header('Authorization') || '';
    const [, token] = bearer.split('Bearer ');
    authClient
      .verifyIdToken({
        idToken: token,
        audience: options.audience, // Audience(optional) on subscription details on console
      })
      .then((ticket: LoginTicket) => {
        const claims = ticket.getPayload();
        (req as RequestWithClaims).claims = claims;

        if (options.email !== claims?.email) {
          throw new Error(`Unexpected email ${claims?.email}`);
        }

        next();
      })
      .catch(err => next(createError(403, 'err')));
  };
}
