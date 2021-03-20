import { Request, Response } from 'express';
import fs from 'fs';
import path from 'path';

import Config from '@config/index';
import { HttpStatusCode } from '@constants/index';
import BaseError from '@constants/BaseError';
import { InternalError } from '@constants/errors';
import IMusicsIntegration from '@integrations/MusicsIntegration/interface';
import { getFileStatus } from '@utils/index';

export default class MusicsController {
  private musicsIntegration: IMusicsIntegration;

  constructor(musicsIntegration: IMusicsIntegration) {
    this.musicsIntegration = musicsIntegration;
  }

  public async stream(request: Request, response: Response) {
    const { id } = request.params;

    try {
      const music = await this.musicsIntegration.getMusic(id);

      const storagePath = Config.storage.path;
      const musicFile = path.resolve(storagePath, music.file);

      const stat = await getFileStatus(musicFile);

      response.writeHead(HttpStatusCode.OK, {
        'Content-Type': 'audio/mp3',
        'Content-Length': stat.size,
      });

      const stream = fs.createReadStream(musicFile);
      stream.pipe(response);

      //
    } catch (error) {
      if (error instanceof BaseError) {
        return response.status(error.statusCode).json({ error: error.message });
      }

      const internalError = new InternalError();
      return response.status(internalError.statusCode).json({ error: internalError.message });
    }
  }
}