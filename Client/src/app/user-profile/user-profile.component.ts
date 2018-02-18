import { Component, OnInit } from '@angular/core';
import { Leaderboard } from '../models/Leaderboard.model';
import { Nextgames } from '../models/Nextgames.model';
import { Runner } from '../models/Runner.model';
import { ActivatedRoute } from '@angular/router';


@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})

export class UserProfileComponent implements OnInit {
  public user: string;

  public testLeader = new Array<Leaderboard>(
    { username: 'ReignSupremeSc2', score: 4600 }
    , { username: 'L337G4M3PR3D1CT0R', score: 1337 }
    , { username: 'TheLegend27', score: 330 }
    , { username: 'SeaGnome', score: 322 }
    , { username: 'Mario', score: 34 }
    , { username: 'WonderBoy#1', score: 31 }
    , { username: 'Iateyourpiefan1', score: 27 }
    , { username: 'IHopeIWinThePredictions', score: 21 }
    , { username: 'GoodGuesser12', score: 12 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
    , { username: 'SpyroTheDragon', score: 1 }
  );
  public gamesBoard = new Array<Nextgames>(
    { title: 'Tomb Raider 1', favorOnePercent: 12, favorTwoPercent: 88, abstainPercentage: 0 }
    , { title: 'Starcraft 2 Campaign', favorOnePercent: 35, favorTwoPercent: 60, abstainPercentage: 5 }
    , { title: 'jimmy newtrons great adventure', favorOnePercent: 61, favorTwoPercent: 39, abstainPercentage: 0 }
    , { title: 'Deadliest CatFish: The Game', favorOnePercent: 12, favorTwoPercent: 77, abstainPercentage: 11 }
    , { title: 'Skyrim, any %', favorOnePercent: 31, favorTwoPercent: 55, abstainPercentage: 14 }
    , { title: 'jimmy newtrons great adventure', favorOnePercent: 61, favorTwoPercent: 39, abstainPercentage: 0 }
    , { title: 'Deadliest CatFish: The Game', favorOnePercent: 12, favorTwoPercent: 77, abstainPercentage: 11 }
    , { title: 'Skyrim, any %', favorOnePercent: 31, favorTwoPercent: 55, abstainPercentage: 14 }
    , { title: 'jimmy newtrons great adventure', favorOnePercent: 61, favorTwoPercent: 39, abstainPercentage: 0 }
    , { title: 'Deadliest CatFish: The Game', favorOnePercent: 12, favorTwoPercent: 77, abstainPercentage: 11 }
    , { title: 'Skyrim, any %', favorOnePercent: 31, favorTwoPercent: 55, abstainPercentage: 14 }
    , { title: 'jimmy newtrons great adventure', favorOnePercent: 61, favorTwoPercent: 39, abstainPercentage: 0 }
    , { title: 'Deadliest CatFish: The Game', favorOnePercent: 12, favorTwoPercent: 77, abstainPercentage: 11 }
    , { title: 'Skyrim, any %', favorOnePercent: 31, favorTwoPercent: 55, abstainPercentage: 14 }
    , { title: 'jimmy newtrons great adventure', favorOnePercent: 61, favorTwoPercent: 39, abstainPercentage: 0 }
    , { title: 'Deadliest CatFish: The Game', favorOnePercent: 12, favorTwoPercent: 77, abstainPercentage: 11 }
    , { title: 'Skyrim, any %', favorOnePercent: 31, favorTwoPercent: 55, abstainPercentage: 14 }
    , { title: 'jimmy newtrons great adventure', favorOnePercent: 61, favorTwoPercent: 39, abstainPercentage: 0 }
    , { title: 'Deadliest CatFish: The Game', favorOnePercent: 12, favorTwoPercent: 77, abstainPercentage: 11 }
    , { title: 'Skyrim, any %', favorOnePercent: 31, favorTwoPercent: 55, abstainPercentage: 14 }
  );
  public RunnerOne: Runner = { username: 'Spikevegeta', score: 21, gamesPlayed: 34 };
  public RunnerTwo: Runner = { username: 'Iateyourpie', score: 13, gamesPlayed: 34 };

  constructor(private activatedRoute: ActivatedRoute) { }

  ngOnInit() {
    this.getUserFromRoute();
  }

  getUserFromRoute() {
    this.activatedRoute.params.subscribe(params => {
      const user = params['user'];
      this.user = user;
      this.searchUser(user);
    });
  }

  searchUser(user: string) {
    console.log(user);
  }

  selectUser(user) {
    this.user = user;
  }

}
