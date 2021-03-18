import cors from 'cors';
import express, { NextFunction, Request, Response, Router } from 'express';
import rateLimit from 'express-rate-limit';
import helmet from 'helmet';
import * as http from 'http';
import { Socket } from 'net';

import IRestChannel from '../interface';
import MusicsController from './controllers/MusicsController';
import TokensController from './controllers/TokensController';
import UsersController from './controllers/UsersController';
import Authentication from '../middlewares/Authentication';
import Config from '@config/index';
import { HttpStatusCode } from '@constants/index';
import BaseError from '@constants/BaseError';
import { ErrorInvalidToken, InternalError } from '@constants/errors';
import IErrorHandler from '@handlers/ErrorHandler/interface';
import IMusicsIntegration from '@integrations/MusicsIntegration/interface';
import IUsersIntegration from '@integrations/UsersIntegration/interface';
import ILoggerProvider from '@providers/LoggerProvider/interface';
import { delay } from '@utils/index';

export default class ExpressRestChannel implements IRestChannel {
  private express: express.Express;
  private server: http.Server;
  private sockets: Map<number, Socket>;

  private authenticationProvider: Authentication;
  private errorHandler: IErrorHandler;
  private loggerProvider: ILoggerProvider;

  private musicsController: MusicsController;
  private tokensController: TokensController;
  private usersController: UsersController;

  // prettier-ignore
  constructor(
    musicsIntegration: IMusicsIntegration,
    usersIntegration: IUsersIntegration,
    errorHandler: IErrorHandler,
    loggerProvider: ILoggerProvider,
  ) {
    this.express = express();
    this.server = new http.Server();
    this.sockets = new Map();

    this.authenticationProvider = new Authentication();
    this.errorHandler = errorHandler;
    this.loggerProvider = loggerProvider;

    this.usersController = new UsersController(usersIntegration);
    this.tokensController = new TokensController(usersIntegration);
    this.musicsController = new MusicsController(musicsIntegration);
  }

  public start(): void {
    this.initMiddlewares();

    this.initRouter();

    const PORT = Config.channels.rest.port;
    const HOST = Config.channels.rest.host;

    this.server = this.express.listen(PORT, HOST, () => {
      this.loggerProvider.info(`Rest server is running on port ${PORT}.`);
    });

    this.server.on('connection', (socket: Socket) => {
      const socketId = this.sockets.size + 1;
      this.sockets.set(socketId, socket);

      socket.once('close', () => {
        this.sockets.delete(socketId);
      });
    });
  }

  public async stop(): Promise<void> {
    this.loggerProvider.info('Stopping rest server ...');

    if (this.sockets.size > 0) {
      await this.waitForSocketsToClose();
    }

    return new Promise((resolve, reject) => {
      this.server.close((error: Error | undefined) => {
        if (error) {
          this.errorHandler.handleError(error);
          reject(error);
        } else {
          this.loggerProvider.info('Rest server stopped.');
          resolve();
        }
      });
    });
  }

  private async waitForSocketsToClose() {
    for (let counter = 5; counter > 0; counter--) {
      if (this.sockets.size > 0) {
        this.loggerProvider.info(`Waiting ${counter} more ${counter !== 1 ? 'seconds' : 'second'} for all connections to close ...`);
        await delay(1);
      } else {
        break;
      }
    }

    this.loggerProvider.info('Forcing all connections to close now.');
    this.sockets.forEach(socket => {
      socket.destroy();
    });
  }

  private initMiddlewares(): void {
    this.express.use(helmet());

    const OneHour = 60 * 60 * 1000;
    const limit = rateLimit({
      max: 100,
      windowMs: OneHour,
      message: 'Too many requests.',
      headers: true,
      statusCode: HttpStatusCode.TOO_MANY_REQUESTS,
    });
    this.express.use(limit);

    this.express.use(cors());
    this.express.use(express.json({ limit: '10kb' }));

    this.express.use(this.logging.bind(this));
  }

  private initRouter(): void {
    const router = Router();

    router.get('/', this.checkAccess.bind(this), this.usersController.show.bind(this.usersController));
    router.post('/', this.usersController.create.bind(this.usersController));
    router.patch('/', this.checkAccess.bind(this), this.usersController.update.bind(this.usersController));
    router.delete('/', this.checkAccess.bind(this), this.usersController.delete.bind(this.usersController));

    router.post('/token', this.tokensController.create.bind(this.tokensController));

    router.get('/music/:id', this.musicsController.stream.bind(this.musicsController));

    this.express.use(router);
  }

  private checkAccess(request: Request, response: Response, next: NextFunction): void | Response {
    const authenticationHeader = request.headers.authorization;

    if (!authenticationHeader) {
      const error = new ErrorInvalidToken();
      return response.status(error.statusCode).json({ error: error.message });
    }

    try {
      const id = this.authenticationProvider.authentication(authenticationHeader);

      request.userId = id;

      return next();

      //
    } catch (error) {
      if (error instanceof BaseError) {
        return response.status(error.statusCode).json({ error: error.message });
      }

      const internalError = new InternalError();
      return response.status(internalError.statusCode).json({ error: internalError.message });
    }
  }

  private logging(request: Request, response: Response, next: NextFunction): void {
    this.loggerProvider.info('', {
      'http-version': request.httpVersionMajor + '.' + request.httpVersionMinor,
      method: request.method,
      URL: request.originalUrl || request.url,
      'user-agent': request.headers['user-agent'],
      'remote-addr': request.ip || (request.connection && request.connection.remoteAddress) || undefined,
    });

    return next();
  }
}
