import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { ActiveEventResponse } from '../models/ActiveEventResponse.model';
import { Leaderboard } from '../models/Leaderboard.model';
import { Nextgames } from '../models/Nextgames.model';
import { Prediction } from '../models/Prediction.model';

@Injectable()
export class UserProfileService {

  constructor(private http: HttpClient) { }

  getActiveEvent(): Observable<Array<ActiveEventResponse>> {
    return this.http.get<Array<ActiveEventResponse>>(environment.Base_URL + 'activeevent');
  }

  getLeaderBoard(eventID: number): Observable<Array<Leaderboard>> {
    return this.http.get<Array<Leaderboard>>(`${environment.Base_URL}score?event=${eventID}&user=-1`);
  }

  getGames(eventID: number): Observable<Array<Nextgames>> {
    return this.http.get<Array<Nextgames>>(`${environment.Base_URL}game?past=false&event=${eventID}`);
  }

  getPredictions(eventID: number, user: string): Observable<Array<Prediction>> {
    return this.http.get<Array<Prediction>>(`${environment.Base_URL}predictions?event=${eventID}&user=${user}`);
  }
}
