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
  public leaderPercent = new Array<Leaderboard>();
  public gamesBoard = new Array<Nextgames>();
  public RunnerOne: Runner = { username: '', score: 0, gamesPlayed: 0 };
  public RunnerTwo: Runner = { username: '', score: 0, gamesPlayed: 0 };
  public RunnerThree: Runner = { username: '', score: 0, gamesPlayed: 0 };
  public RunnerFour: Runner = { username: '', score: 0, gamesPlayed: 0 };
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
          if (res.length > 0 && res[0].id) {
            this.activeEvent = res[0];

            // TESTING REMOVE AND PUT ACTIVE EVENT
            // kk
            this.getLeaderboard(this.activeEvent.id);
            this.getGames(this.activeEvent.id);
            this.getMatchReults(this.activeEvent.id);
            this.getPredictions(this.activeEvent.id, this.user);
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
          this.leaderPercent = res.slice(0);
          // sort by % and fix leaderboard
          this.leaderPercent = this.leaderPercent.sort(function (a, b) {
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
          this.RunnerOne.username = this.gamesBoard[0].First.Competitor;
          this.RunnerTwo.username = this.gamesBoard[0].Second.Competitor;
          this.RunnerThree.username = this.gamesBoard[0].Third.Competitor;
          this.RunnerFour.username = this.gamesBoard[0].Second.Competitor;
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
      },
      (err) => {
        this.userPrediction = [];
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
    this.getPredictions(this.activeEvent.id, this.user);
    this.findUserCard(user);
  }

  findUserCard(user: string) {
    const found = this.leaderboard.find(function (element) {
      return element.User.toLowerCase() === user.toLowerCase();
    });
    if (found != null) {
      this.userCard.User = user;
      this.userCard.Total = found.Total;
      this.userCard.Percent = found.Percent;
      this.userCard.LeaderboardPlacement = this.leaderboard.indexOf(found) + 1;
    } else {
      this.userCard = { User: 'N/A', Total: 0, Percent: 0, LeaderboardPlacement: 0 };
      this.userPrediction = [];
    }
  }

  searchUser() {
    this.getPredictions(this.activeEvent.id, this.user);
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

  public enterPress() {
    this.searchUser();
  }

}
