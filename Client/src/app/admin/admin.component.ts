import { Component, OnInit } from '@angular/core';
import { AdminService } from './admin.service';
import { Streamer } from '../models/Streamer.model';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
import { HttpParams } from '@angular/common/http';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { EventRequest } from '../models/EventRequest.model';
import { EventResponse } from '../models/EventResponse.model';
import { StreamerSetRequest } from '../models/StreamerSetRequest.model';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})



export class AdminComponent implements OnInit {
  public authenticated: boolean;
  public oldPassword: string;
  public newPassword: string;
  public streamers: Streamer[];
  public streamerOne: string;
  public streamerTwo: string;
  public trustedUrl: SafeUrl[] = [];
  public EventName: string;
  public eventName: string;
  public eventList = new Array<EventResponse>();
  public startInsertRows = { id: 0, game: 'Insert Game', competitor: ['Iateyourpie', 'Spikevegeta'] };
  public closeResult: string;
  public eventCreateRows: any[] = [];
  public showEditable: boolean;
  public editRowId: any;
  public editColumnId: any;
  public eventRequest: EventRequest = { Name: '' };
  // Password and credential html models
  public Username: string;
  public Password: string;
  // tslint:disable-next-line:max-line-length
  public manageEventsObject = { updatePredictions: null, updateEventResults: null, completeEvent: null, deleteEvent: null, activateEvent: null };
  constructor(public adminservice: AdminService,
    private sanitizer: DomSanitizer,
    private modalService: NgbModal) { }


  open(content) {
    this.modalService.open(content, { size: 'lg' }).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
    });
  }

  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return 'by clicking on a backdrop';
    } else {
      return `with: ${reason}`;
    }
  }



  ngOnInit() {
    this.getStreamers();
    this.eventCreateRows.push(this.startInsertRows);
    this.getEvents();
  }

  public getStreamers() {
    this.adminservice.getStreamers().subscribe(
      (streamers: Streamer[]) => {
        if (!streamers) {
          return;
        }
        streamers.forEach(streamer => {
          // tslint:disable-next-line:max-line-length
          this.trustedUrl.push(this.sanitizer.bypassSecurityTrustResourceUrl('http://player.twitch.tv/?channel=' + streamer.tag + '&muted=true'));
        });
      }
    );
  }

  public updateStreamers() {
    const streamRequest: StreamerSetRequest = { streamerOne: this.streamerOne, streamerTwo: this.streamerTwo };
    this.adminservice.putStreamers(streamRequest).subscribe(
      (res) => {
      }
    );
  }


  public createEvent() {
    this.adminservice.CreateEvent(this.eventRequest).subscribe(
      (res) => {
        this.getEvents();
      }
    );
  }

  public getEvents() {
    this.adminservice.getEvents().subscribe(
      (res) => {
        if (!res) {
          return;
        }
        this.eventList = res.filter((event) => {
          return event.complete === false;
        });
      }
    );
  }

  public completeEvent() {
    this.adminservice.CompeleteEvent(this.manageEventsObject.completeEvent, true).subscribe(
      (res) => {
        this.eventList = this.eventList.filter((elm) => elm.id !== this.manageEventsObject.completeEvent);
      }
    );
  }

  public deleteEvent() {
    this.adminservice.DeleteEvent(this.manageEventsObject.deleteEvent).subscribe(
      (res) => {
        this.eventList = this.eventList.filter((elm) => elm.id !== this.manageEventsObject.deleteEvent);
      }
    );

  }

  public activateEvent() {
    this.adminservice.ActivateEvent(this.manageEventsObject.activateEvent).subscribe(
      (res) => {
        
      }
    );
  }

  public uploadExcel(event) {
    const fileList: FileList = event.target.files;
    if (fileList.length === 0) {
      return;
    }
    const file = fileList[0];

    const formData = new FormData();
    formData.append('uploadFile', file, file.name);
    formData.append('eventID', this.manageEventsObject.updatePredictions);
    const headers = new Headers();
    /** No need to include Content-Type in Angular 4 */

    const params = new HttpParams();
    const options = {
      headers: headers,
      params: params,
      reportProgress: true,
    };

    this.adminservice.uploadExcel(formData, options).subscribe(
      (res) => {
      }
    );
  }


  public uploadResults(event) {
    const fileList: FileList = event.target.files;
    if (fileList.length === 0) {
      return;
    }
    const file = fileList[0];

    const formData = new FormData();
    formData.append('uploadFile', file, file.name);
    formData.append('eventID', this.manageEventsObject.updateEventResults);
    const headers = new Headers();
    /** No need to include Content-Type in Angular 4 */

    const params = new HttpParams();
    const options = {
      headers: headers,
      params: params,
      reportProgress: true,
    };

    this.adminservice.uploadResults(formData, options).subscribe(
      (res) => {
      }
    );
  }


  public authenticateUser() {
    
    if (this.Password.trim() !== '' && this.Username.trim() !== '') {
      // Check if user should be authenticated
      this.adminservice.basicAuthenticateUser(this.Username, this.Password).subscribe(res => {
        this.authenticated = true;
      }, err => {
        console.log(err);
        this.authenticated = false;
      });
    } else {
      // invalid input
      this.authenticated = false;
    }
  }

  // addNewRow Expands the modal's row
  public addNewRow() {
    this.eventCreateRows.push({
      id: this.eventCreateRows.length + 1,
      game: 'Insert Game',
      competitor: ['Iateyourpie', 'Spikevegeta'],
    });
  }

  public toggle(val) {
    this.editRowId = val;
  }

  public toggleColumn(y) {
    this.editColumnId = y;
  }

  public changePassword() {
    this.adminservice.changePassword(this.oldPassword, this.newPassword).subscribe(
      (res) => {
      },
      (err) => {
        console.log(err);
      }
    );
  }
}
