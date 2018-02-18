import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Streamer } from '../models/Streamer.model';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class AdminService {

  constructor(private http: HttpClient) { }

  public getStreamers(): Observable<Streamer[]> {
    return this.http.get<Streamer[]>(environment.Base_URL + 'streamer');
  }

  public putStreamers(streamers: Streamer[]): any {
    return this.http.put(environment.Base_URL + 'streamer', JSON.stringify(streamers));
  }

  public postStreamers(streamers: Streamer[]): any {
    return this.http.post(environment.Base_URL + 'streamer', JSON.stringify(streamers));
  }

  public getEvents(): Observable<Array<string>> {
    return this.http.get<string[]>(environment.Base_URL + 'events');
  }

  public CreateEvent(): Observable<Array<string>> {
    return this.http.get<string[]>(environment.Base_URL + 'events');
  }

  public uploadExcel(excelFile: FormData, options): any {
    return this.http.put(environment.Base_URL + 'eventsUpload', excelFile, options);
  }
}
