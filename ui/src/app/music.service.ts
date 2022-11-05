import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class MusicService {
  constructor(private httpClient: HttpClient) { }

  getTopItalians() {
    return this.httpClient.get<Song[]>(environment.gateway + '/get-top-italians');
  }

}

export class Artist {
  id: number | undefined;
  name: string | undefined;
  picture: string | undefined;
}

export class Song {
  title: string | undefined;
  preview: string | undefined;
  mp3: string | undefined;
  duration: string | undefined;
  rank: string | undefined;
  url: string | undefined;
  artist: Artist | undefined;
}