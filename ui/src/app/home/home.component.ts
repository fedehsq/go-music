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
    console.log('qua!');
    this.getTopItalians();
  }

  getTopItalians() {
    this.musicService.getTopItalians().subscribe((data: any) => {
      this.topItalians = data;
    });
  }


}
