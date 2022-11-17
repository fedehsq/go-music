import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class MusicService {
  constructor(private httpClient: HttpClient) { }

  getChartPlaylists() {
    return this.httpClient.get<ChartPlaylist[]>(environment.gateway + '/chart-playlists');
  }

}

export class Artist {
  id: number | undefined;
  name: string | undefined;
  picture: string | undefined;
  tracklist: string | undefined;
}

export class Track {
  title: string | undefined;
  preview: string | undefined;
  mp3: string | undefined;
  duration: string | undefined;
  rank: string | undefined;
  url: string | undefined;
  artist: Artist | undefined;
}

export class ChartPlaylist {
  id: number | undefined;
  title: string | undefined;
  picture_big: string | undefined;
  link: string | undefined;
  tracklist: string | undefined;
}

export class Playlist {
  id: number | undefined;
  title: string | undefined;
  pictureBig: string | undefined;
  duration: number | undefined;
  nbtracks: number | undefined;
  description: string | undefined;
  tracks: Track[] | undefined;
}

export class Album {
  id: number | undefined;
  title: string | undefined;
  cover: string | undefined;
  artist: Artist | undefined;
  tracks: Track[] | undefined;
}