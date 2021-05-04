import * as grpc from 'grpc';

import { AddTrack, CreatePlaylist, DeletePlaylist, GetPlaylist, GetPlaylists, RemoveTrack, UpdatePlaylist, UpdateTrack } from './dtos';
import IPlaylistsIntegration from './interface';
import { PlaylistsClient } from '../proto/playlists_service_grpc_pb';
import {
  AddTrackRequest,
  CreatePlaylistRequest,
  DeletePlaylistRequest,
  GetPlaylistRequest,
  GetPlaylistsRequest,
  Playlist,
  PlaylistsList,
  RemoveTrackRequest,
  Track,
  UpdatePlaylistRequest,
  UpdateTrackRequest,
} from '../proto/playlists_service_pb';
import { translatePlaylistEntity, translatePlaylistEntityList, translateTrackEntity } from '../translators';
import Config from '@config/index';
import PlaylistEntity, { Track as TrackEntity } from '@entities/Playlist';

export default class PlaylistsIntegration implements IPlaylistsIntegration {
  private client: PlaylistsClient;

  constructor() {
    const ADDRESS = Config.integrations.playlists_service;

    this.client = new PlaylistsClient(ADDRESS, grpc.credentials.createInsecure());
  }

  public getPlaylist = async ({ id, userId }: GetPlaylist): Promise<PlaylistEntity> => {
    return new Promise((resolve, reject) => {
      const getPlaylistRequest = new GetPlaylistRequest();
      getPlaylistRequest.setId(id);
      getPlaylistRequest.setUserid(userId);

      this.client.getPlaylist(getPlaylistRequest, (error: Error | null, playlist: Playlist) => {
        if (error != null) reject(error);
        else resolve(translatePlaylistEntity(playlist));
      });
    });
  };

  public getPlaylists = async ({ userId }: GetPlaylists): Promise<Array<PlaylistEntity>> => {
    return new Promise((resolve, reject) => {
      const getPlaylistsRequest = new GetPlaylistsRequest();
      getPlaylistsRequest.setUserid(userId);

      this.client.getPlaylists(getPlaylistsRequest, (error: Error | null, playlistsList: PlaylistsList) => {
        if (error != null) reject(error);
        else resolve(translatePlaylistEntityList(playlistsList));
      });
    });
  };

  public createPlaylist = async ({ name, userId }: CreatePlaylist): Promise<PlaylistEntity> => {
    return new Promise((resolve, reject) => {
      const createPlaylistRequest = new CreatePlaylistRequest();
      createPlaylistRequest.setName(name);
      createPlaylistRequest.setUserid(userId);

      this.client.createPlaylist(createPlaylistRequest, (error: Error | null, playlist: Playlist) => {
        if (error != null) reject(error);
        else resolve(translatePlaylistEntity(playlist));
      });
    });
  };

  public updatePlaylist = async ({ id, userId, name }: UpdatePlaylist): Promise<PlaylistEntity> => {
    return new Promise((resolve, reject) => {
      const updatePlaylistRequest = new UpdatePlaylistRequest();
      updatePlaylistRequest.setId(id);
      updatePlaylistRequest.setUserid(userId);
      updatePlaylistRequest.setName(name ? name : '');

      this.client.updatePlaylist(updatePlaylistRequest, (error: Error | null, playlist: Playlist) => {
        if (error != null) reject(error);
        else resolve(translatePlaylistEntity(playlist));
      });
    });
  };

  public deletePlaylist = async ({ id, userId }: DeletePlaylist): Promise<void> => {
    return new Promise((resolve, reject) => {
      const deletePlaylistRequest = new DeletePlaylistRequest();
      deletePlaylistRequest.setId(id);
      deletePlaylistRequest.setUserid(userId);

      this.client.deletePlaylist(deletePlaylistRequest, (error: Error | null) => {
        if (error != null) reject(error);
        else resolve();
      });
    });
  };

  public addTrack = async ({ musicId, playlistId, userId }: AddTrack): Promise<TrackEntity> => {
    return new Promise((resolve, reject) => {
      const addTrackRequest = new AddTrackRequest();
      addTrackRequest.setMusicid(musicId);
      addTrackRequest.setPlaylistid(playlistId);
      addTrackRequest.setUserid(userId);

      this.client.addTrack(addTrackRequest, (error: Error | null, track: Track) => {
        if (error != null) reject(error);
        else resolve(translateTrackEntity(track));
      });
    });
  };

  public updateTrack = async ({ id, index, playlistId, userId }: UpdateTrack): Promise<TrackEntity> => {
    return new Promise((resolve, reject) => {
      const updateTrackRequest = new UpdateTrackRequest();
      updateTrackRequest.setId(id);
      updateTrackRequest.setIndex(index);
      updateTrackRequest.setPlaylistid(playlistId);
      updateTrackRequest.setUserid(userId);

      this.client.updateTrack(updateTrackRequest, (error: Error | null, track: Track) => {
        if (error != null) reject(error);
        else resolve(translateTrackEntity(track));
      });
    });
  };

  public removeTrack = async ({ id, playlistId, userId }: RemoveTrack): Promise<void> => {
    return new Promise((resolve, reject) => {
      const removeTrackRequest = new RemoveTrackRequest();
      removeTrackRequest.setId(id);
      removeTrackRequest.setPlaylistid(playlistId);
      removeTrackRequest.setUserid(userId);

      this.client.removeTrack(removeTrackRequest, (error: Error | null) => {
        if (error != null) reject(error);
        else resolve();
      });
    });
  };
}
