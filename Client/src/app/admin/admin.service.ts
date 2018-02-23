import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Streamer } from '../models/Streamer.model';
import { Observable } from 'rxjs/Observable';
import { EventRequest } from '../models/EventRequest.model';
import { EventResponse } from '../models/EventResponse.model';
import { StreamerSetRequest } from '../models/StreamerSetRequest.model';

@Injectable()
export class AdminService {

  constructor(private http: HttpClient) { }

  public getStreamers(): Observable<Streamer[]> {
    return this.http.get<Streamer[]>(environment.Base_URL + 'streamer');
  }

  public putStreamers(streamerRequest: StreamerSetRequest): any {
    return this.http.put(environment.Base_URL + 'streamer', JSON.stringify(streamerRequest));
  }

  public getEvents(): Observable<Array<EventResponse>> {
    return this.http.get<Array<EventResponse>>(environment.Base_URL + 'events');
  }

  public CreateEvent(eventRequest: EventRequest): Observable<any> {
    return this.http.put(environment.Base_URL + 'events', eventRequest);
  }

  public CompeleteEvent(eventID: number, completedStatus: boolean): Observable<any> {
    return this.http.post(`${environment.Base_URL}events?id=${eventID}&completed=${completedStatus}`, null);
  }

  public DeleteEvent(eventID: number): Observable<any> {
    return this.http.delete(`${environment.Base_URL}events?id=${eventID}`);
  }

  public ActivateEvent(eventID: number): Observable<any> {
    return this.http.delete(`${environment.Base_URL}events?id=${eventID}`);
  }


  public uploadExcel(excelFile: FormData, options): any {
    return this.http.put(environment.Base_URL + 'predictions/upload', excelFile, options);
  }

  public uploadResults(excelFile: FormData, options): any {
    return this.http.put(environment.Base_URL + 'results/upload', excelFile, options);
  }
}
