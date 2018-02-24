import { Component, OnInit } from '@angular/core';
import { Leaderboard } from '../models/Leaderboard.model';
import { Nextgames } from '../models/Nextgames.model';
import { Runner } from '../models/Runner.model';
import { ActivatedRoute } from '@angular/router';
import { UserProfileService } from './user-profile.service';
import { ActiveEventResponse } from '../models/ActiveEventResponse.model';
import { UserCard } from '../models/UserCard.model';
import { Prediction } from '../models/Prediction.model';
import { GameRunners } from '../models/GameRunners.model';


@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})

export class UserProfileComponent implements OnInit {
  public user: string;
  public userCard: UserCard = { User: 'N/A', Total: 0, Percent: 0, LeaderboardPlacement: 0 };
  public activeEvent: ActiveEventResponse;

  public leaderboard = new Array<Leaderboard>();
  public leaderboardPercent = new Array<Leaderboard>();
  public gamesBoard = new Array<Nextgames>();
  public RunnerOne: Runner = { username: '', score: 0, gamesPlayed: 0 };
  public RunnerTwo: Runner = { username: '', score: 0, gamesPlayed: 0 };
  public runners = new Array<GameRunners>();
  public totalresults: Array<{ username: '', won: 0, total: 0 }>;
  public userPrediction: Array<Prediction>;
  public leaderFilter: boolean;

  constructor(private activatedRoute: ActivatedRoute, private userProfileService: UserProfileService) { }

  ngOnInit() {
    this.getUserFromRoute();
    this.getActiveEvent();
  }

  getActiveEvent() {
    this.userProfileService.getActiveEvent().subscribe(
      (res: Array<ActiveEventResponse>) => {
        if (res != null) {
          if (res.length > 0) {
            this.activeEvent = res[0];

            // TESTING REMOVE AND PUT ACTIVE EVENT
            this.getLeaderboard(45);
            this.getGames(45);
            this.getMatchReults(45);
            this.getPredictions(45, this.user);
          }
        }
      }
    );
  }

  getLeaderboard(eventID: number) {
    this.userProfileService.getLeaderBoard(eventID).subscribe(
      (res) => {
        if (res != null) {
          this.leaderboard = res;
          // sort by % and fix leaderboard
          this.leaderboardPercent = this.leaderboard.sort(function (a, b) {
            if (a.Percent < b.Percent) {
              return 1;
            }
            if (a.Percent > b.Percent) {
              return -1;
            }
            return 0;
          });
          this.findUserCard(this.user);
        }
      }
    );
  }
  getGames(eventID: number) {
    this.userProfileService.getGames(eventID).subscribe(
      (res) => {
        if (res != null) {
          this.gamesBoard = res;
        }
      }
    );
  }

  getPredictions(eventID: number, user: string) {
    this.userProfileService.getPredictions(eventID, user).subscribe(
      (res) => {
        if (res != null) {
          this.userPrediction = res;
        }
      }
    );
  }

  getUserFromRoute() {
    this.activatedRoute.params.subscribe(params => {
      const user = params['user'];
      this.user = user;
    });
  }


  selectUser(user, index: number) {
    this.user = user;
    this.getPredictions(45, this.user);
    this.findUserCard(user);
  }

  findUserCard(user: string) {
    const found = this.leaderboard.find(function (element) {
      return element.User === user;
    });
    if (found != null) {
      this.userCard.User = user;
      this.userCard.Total = found.Total;
      this.userCard.Percent = found.Percent;
      this.userCard.LeaderboardPlacement = this.leaderboard.indexOf(found) + 1;
    }
  }

  searchUser() {
    this.getPredictions(45, this.user);
    this.findUserCard(this.user);
  }

  swapLeaderboard() {
    this.leaderFilter = !this.leaderFilter;
  }

  getMatchReults(eventID: number) {
    this.userProfileService.getGamesResult(eventID).subscribe(
      (res) => {
        this.runners = res;
      }
    );
  }

}
