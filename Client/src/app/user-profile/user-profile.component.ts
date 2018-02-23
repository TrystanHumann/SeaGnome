import { Component, OnInit } from '@angular/core';
import { Leaderboard } from '../models/Leaderboard.model';
import { Nextgames } from '../models/Nextgames.model';
import { Runner } from '../models/Runner.model';
import { ActivatedRoute } from '@angular/router';
import { UserProfileService } from './user-profile.service';
import { ActiveEventResponse } from '../models/ActiveEventResponse.model';
import { UserCard } from '../models/UserCard.model';
import { Prediction } from '../models/Prediction.model';


@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})

export class UserProfileComponent implements OnInit {
  public user: string;
  public userCard: UserCard = { User: 'LOADING', Total: 0, Percent: 0, LeaderboardPlacement: 0 };
  public activeEvent: ActiveEventResponse;

  public leaderboard = new Array<Leaderboard>();
  public gamesBoard = new Array<Nextgames>();
  public RunnerOne: Runner = { username: 'Spikevegeta', score: 21, gamesPlayed: 34 };
  public RunnerTwo: Runner = { username: 'Iateyourpie', score: 13, gamesPlayed: 34 };
  public totalresults: Array<{ username: '', won: 0, total: 0 }>;
  public userPrediction: Array<Prediction>;
  constructor(private activatedRoute: ActivatedRoute, private userProfileService: UserProfileService) { }

  ngOnInit() {
    this.getUserFromRoute();
    this.getActiveEvent();
  }

  getActiveEvent() {
    this.userProfileService.getActiveEvent().subscribe(
      (res: Array<ActiveEventResponse>) => {
        if (res != null) {
          console.log(res);
          if (res.length > 0) {
            this.activeEvent = res[0];
            console.log(this.activeEvent);

            // TESTING REMOVE AND PUT ACTIVE EVENT
            this.getLeaderboard(45);
            this.getGames(45);
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
          console.log(res);
          this.leaderboard = res;
          this.findUserCard(this.user);
        }
      }
    );
  }
  getGames(eventID: number) {
    this.userProfileService.getGames(eventID).subscribe(
      (res) => {
        console.log('heres some data bitch', res);
        if (res != null) {
          this.gamesBoard = res;
          console.log('gamerboard', this.gamesBoard);
        }
      }
    );
  }

  getPredictions(eventID: number, user: string) {
    this.userProfileService.getPredictions(eventID, user).subscribe(
      (res) => {
        console.log('prediciton data', res);
        if (res != null) {
          this.userPrediction = res;
          console.log('prediction set', res);
        }
      }
    );
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

  sumResults() {
    // sum results from game board
  }

}
