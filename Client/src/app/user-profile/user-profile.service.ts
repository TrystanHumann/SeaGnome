import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { ActiveEventResponse } from '../models/ActiveEventResponse.model';
import { Leaderboard } from '../models/Leaderboard.model';
import { Nextgames } from '../models/Nextgames.model';
import { Prediction } from '../models/Prediction.model';
import { Runner } from '../models/Runner.model';
import { GameRunners } from '../models/GameRunners.model';

@Injectable()
export class UserProfileService {

  private baseurl: string;

  constructor(private http: HttpClient)  {
    if (environment.production) {
      this.baseurl = window.location.href.substr(0, window.location.href.indexOf("momam.moe")) + environment.Base_URL;
    } else {
      this.baseurl = environment.Base_URL;
    }
  }

  getActiveEvent(): Observable<Array<ActiveEventResponse>> {
    return this.http.get<Array<ActiveEventResponse>>(this.baseurl + 'activeevent');
  }

  getLeaderBoard(eventID: number): Observable<Array<Leaderboard>> {
    return this.http.get<Array<Leaderboard>>(`${this.baseurl}score?event=${eventID}&user=-1`);
  }

  getGames(eventID: number): Observable<Array<Nextgames>> {
    return this.http.get<Array<Nextgames>>(`${this.baseurl}game?past=false&event=${eventID}`);
  }

  getGamesResult(eventID: number): Observable<Array<GameRunners>> {
    return this.http.get<Array<GameRunners>>(`${this.baseurl}game?past=true&event=${eventID}`);
  }

  getPredictions(eventID: number, user: string): Observable<Array<Prediction>> {
    return this.http.get<Array<Prediction>>(`${this.baseurl}predictions?event=${eventID}&user=${user}`);
  }
}
