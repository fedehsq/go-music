import { Component } from '@angular/core';
import { MusicService, Song } from './music.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
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

