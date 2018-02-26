import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Streamer } from '../models/Streamer.model';
import { Observable } from 'rxjs/Observable';
import { EventRequest } from '../models/EventRequest.model';
import { EventResponse } from '../models/EventResponse.model';
import { StreamerSetRequest } from '../models/StreamerSetRequest.model';

@Injectable()
export class AdminService {

  private baseurl: string;

  constructor(private http: HttpClient) {
    if (environment.production) {
      this.baseurl = window.location.href.substr(0, window.location.href.indexOf("momam.moe")) + environment.Base_URL;
    } else {
      this.baseurl = environment.Base_URL;
    }
  }

  public getStreamers(): Observable<Streamer[]> {
    return this.http.get<Streamer[]>(this.baseurl + 'streamer');
  }

  public putStreamers(streamerRequest: StreamerSetRequest): any {
    return this.http.put(this.baseurl + 'streamer', JSON.stringify(streamerRequest));
  }

  public getEvents(): Observable<Array<EventResponse>> {
    return this.http.get<Array<EventResponse>>(this.baseurl + 'events');
  }

  public CreateEvent(eventRequest: EventRequest): Observable<any> {
    return this.http.put(this.baseurl + 'events', eventRequest);
  }

  public CompeleteEvent(eventID: number, completedStatus: boolean): Observable<any> {
    return this.http.post(`${this.baseurl}events?id=${eventID}&completed=${completedStatus}`, null);
  }

  public DeleteEvent(eventID: number): Observable<any> {
    return this.http.delete(`${this.baseurl}events?id=${eventID}`);
  }

  public ActivateEvent(eventID: number): Observable<any> {  
    return this.http.post(`${this.baseurl}activeevent`, {eventID});
  }

  public uploadExcel(excelFile: FormData, options): any {
    return this.http.put(this.baseurl + 'predictions/upload', excelFile, options);
  }

  public uploadResults(excelFile: FormData, options): any {
    return this.http.put(this.baseurl + 'results/upload', excelFile, options);
  }
  // Address what the return type should be
  public basicAuthenticateUser(username: string, password: string): Observable<any> {
    let headers: HttpHeaders = new HttpHeaders();
    headers = headers.append('Authorization', 'Basic ' + btoa(username + ':' + password));
    headers = headers.append('Content-Type', 'application/x-www-form-urlencoded');
    // withCredentials should use cookie ?
    return this.http.get(this.baseurl + 'auth', { headers: headers, withCredentials: true });
  }

  public changePassword(password, newpassword) {
    // withCredentials should use cookie ?
    // tslint:disable-next-line:max-line-length
    return this.http.post(this.baseurl + 'password/change', { oldPassword: btoa(password), newPassword: btoa(newpassword) }, { withCredentials: true });
  }
}
