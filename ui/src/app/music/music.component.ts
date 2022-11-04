import { Component, OnInit } from '@angular/core';
import { MusicService, Song } from '../music.service';

@Component({
  selector: 'app-music',
  templateUrl: './music.component.html',
  styleUrls: ['./music.component.css']
})
export class MusicComponent implements OnInit {

  songs: Song[] | undefined;

  constructor(private musicService: MusicService) { }

  ngOnInit() {
    this.getTopItalians();
  }

  getTopItalians() {
    this.musicService.getTopItalians().subscribe((data: any) => {
      this.songs = data;
    });
  }

}
