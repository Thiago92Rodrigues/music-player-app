import Knex from 'knex';

import { PaginationRequest } from './dtos';
import IArtistsRepository from './interface';
import { AlbumsDb, ArtistsDb } from '../databaseEntities';
import { translateArtist, translateArtistsList } from '../translators';
import { AlbumsTable, ArtistsTable } from '@constants/index';
import Artist from '@entities/Artist';

export default class SQLArtistsRepository implements IArtistsRepository {
  private databaseConnection: Knex;

  constructor(databaseConnection: Knex) {
    this.databaseConnection = databaseConnection;
  }

  public async find(id: string): Promise<Artist | undefined> {
    // prettier-ignore
    const artist = await this.databaseConnection<ArtistsDb>(ArtistsTable)
      .where({ id })
      .first();

    if (!artist) {
      return;
    }

    // prettier-ignore
    const albums = await this.databaseConnection<AlbumsDb>(AlbumsTable)
      .where({ artist_id: artist.id });

    return translateArtist(artist, albums);
  }

  public async findAll(paginationRequest?: PaginationRequest): Promise<Array<Artist>> {
    if (paginationRequest) {
      const { offset, limit } = paginationRequest;

      // prettier-ignore
      const artists = await this.databaseConnection<ArtistsDb>(ArtistsTable)
        .offset(offset)
        .limit(limit)
        .orderBy('name', 'asc');

      return translateArtistsList(artists);
    }

    const artists = await this.databaseConnection<ArtistsDb>(ArtistsTable);

    return translateArtistsList(artists);
  }

  public async findByGenre(genre: number): Promise<Array<Artist>> {
    // prettier-ignore
    const artists = await  this.databaseConnection<ArtistsDb>(ArtistsTable)
      .where({ genre });

    return translateArtistsList(artists);
  }

  public async store({ id, name, country, foundationDate, members, description, genre, photos, facebookUrl, twitterUrl, instagramUrl, wikipediaUrl, favorites, followers }: Artist): Promise<void> {
    // prettier-ignore
    await this.databaseConnection<ArtistsDb>(ArtistsTable)
      .insert({ id, name, country, foundation_date: foundationDate, members, description, genre, photos, facebook_url: facebookUrl, twitter_url: twitterUrl, instagram_url: instagramUrl, wikipedia_url: wikipediaUrl, favorites, followers });
  }

  public async update({ id, name, country, foundationDate, members, description, genre, photos, facebookUrl, twitterUrl, instagramUrl, wikipediaUrl, favorites, followers }: Artist): Promise<void> {
    // prettier-ignore
    await this.databaseConnection<ArtistsDb>(ArtistsTable)
      .where({ id })
      .update({ name, country, foundation_date: foundationDate, members, description, genre, photos, facebook_url: facebookUrl, twitter_url: twitterUrl, instagram_url: instagramUrl, wikipedia_url: wikipediaUrl, favorites, followers });
  }

  public async delete(id: string): Promise<void> {
    // prettier-ignore
    await this.databaseConnection<ArtistsDb>(ArtistsTable)
      .where({ id })
      .del();
  }
}
