import { Component, OnInit } from '@angular/core';
import { MusicService, Song } from '../music.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  topItalians: Song[] | undefined;

  constructor(private musicService: MusicService) { }

  ngOnInit() {
    this.getTopItalians();
  }

  getTopItalians() {
    this.musicService.getTopItalians().subscribe((data: Song[]) => {
      this.topItalians = data;
    });
  }


}
