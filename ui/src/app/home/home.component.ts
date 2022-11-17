import { Component, OnInit } from '@angular/core';
import { MusicService, ChartPlaylist } from '../music.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  chartPlaylists: ChartPlaylist[] | undefined;

  constructor(private musicService: MusicService) { }

  ngOnInit() {
    this.getChartPlaylists();
  }

  getChartPlaylists() {
    this.musicService.getChartPlaylists().subscribe((data: ChartPlaylist[]) => {
      this.chartPlaylists = data;
    });
  }


}
