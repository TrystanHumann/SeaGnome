import { Component, OnInit } from '@angular/core';
import { Leaderboard } from '../models/Leaderboard.model';
import { Nextgames } from '../models/Nextgames.model';
import { Runner } from '../models/Runner.model';


@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})

export class UserProfileComponent implements OnInit {

  public testLeader = new Array<Leaderboard>(
    { username: 'FreeCrabs69', score: 4200 }
    , { username: 'NeverSubNeverDonateGimmie', score: 1337 }
    , { username: 'FreeCrabs6969', score: 420 }
    , { username: 'Reb', score: 22 }
    , { username: 'JohnLee', score: 11 }
    , { username: 'WonderBoy#1(andrew)', score: 1 }
    , { username: 'Nick', score: 0 }
    , { username: 'NeverSubNeverDonateGimmie', score: 1337 }
    , { username: 'FreeCrabs6969', score: 420 }
    , { username: 'Reb', score: 22 }
  );
  public gamesBoard = new Array<Nextgames>(
    { title: 'a new crab', favorOnePercent: 69, favorTwoPercent: 420, abstainPercentage: 31 }
    , { title: 'kill crabs 2 return of the crab', favorOnePercent: 69, favorTwoPercent: 420, abstainPercentage: 31 }
    , { title: 'jimmy newtrons great adventure', favorOnePercent: 69, favorTwoPercent: 420, abstainPercentage: 31 }
    , { title: 'Christian the crab revenger', favorOnePercent: 32, favorTwoPercent: 10, abstainPercentage: 11 }
    , { title: 'The fallen crab killer: Christian, final cut', favorOnePercent: 31, favorTwoPercent: 220, abstainPercentage: 311 }
  );
  public RunnerOne: Runner = { username: 'Spikevegeta', score: 21, gamesPlayed: 34 };
  public RunnerTwo: Runner = { username: 'Iateyourpie', score: 13, gamesPlayed: 34 };

  constructor() { }

  ngOnInit() {
  }

}
