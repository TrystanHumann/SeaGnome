import { Component, OnInit, ViewContainerRef } from '@angular/core';
import { AdminService } from './admin.service';
import { Streamer, ButtonStyle } from '../models/Streamer.model';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
import { HttpParams, HttpHeaders } from '@angular/common/http';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { EventRequest } from '../models/EventRequest.model';
import { EventResponse } from '../models/EventResponse.model';
import { StreamerSetRequest } from '../models/StreamerSetRequest.model';
import { ToastsManager } from 'ng2-toastr/ng2-toastr';
import { timeout } from 'q';
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
  
  public buttonStyleArray : ButtonStyle[];
  public buttonStyleSelected : ButtonStyle;

  constructor(public adminservice: AdminService,
    private sanitizer: DomSanitizer,
    private modalService: NgbModal,
    private toastsManager: ToastsManager) {
  }


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
    this.getButtonStyles();
  }

  public getButtonStyles() {
    this.adminservice.getButtonStyles().subscribe(res => {
      this.buttonStyleArray = res;
      if (this.buttonStyleArray.length > 0)
      {
        this.buttonStyleSelected = this.buttonStyleArray[0];
      }
    }, err => this.toastsManager.error('Unable to fetch button styles', 'Error'));
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
        this.toastsManager.success('Streamers have been updated!');
      },
      (err) => {
        this.toastsManager.error('Error updating streamers. Make sure the names are valid twitch streamers.');
      }
    );
  }


  public createEvent() {
    this.adminservice.CreateEvent(this.eventRequest).subscribe(
      (res) => {
        this.getEvents();
        this.toastsManager.success('Event has been created!');
      },
      (err) => {
        this.toastsManager.error('Error creating event.');
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
        this.toastsManager.success('Event has been completed!');
      },
      (err) => {
        this.toastsManager.error('Error completing event.');
      }
    );
  }

  public deleteEvent() {
    this.adminservice.DeleteEvent(this.manageEventsObject.deleteEvent).subscribe(
      (res) => {
        this.eventList = this.eventList.filter((elm) => elm.id !== this.manageEventsObject.deleteEvent);
        this.toastsManager.success('Event has been removed!');
      },
      (err) => {
        this.toastsManager.error('Error removing event.');
      }
    );

  }

  public activateEvent() {
    this.adminservice.ActivateEvent(this.manageEventsObject.activateEvent).subscribe(
      (res) => {
        this.toastsManager.success('Event has been activated!');
      },
      (err) => {
        this.toastsManager.error('Error activating event.');
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
      Timeout: 1000
    };

    this.adminservice.uploadExcel(formData, options).subscribe(
      (res) => {
        this.toastsManager.success('Event has been uploaded, leave it 45 minutes to process!', null, { toastLife: 300000 });
      },
      (err) => {
        this.toastsManager.success('Event has been uploaded, leave it 45 minutes to process!', null, { toastLife: 300000 });
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
        this.toastsManager.success('Event results have been uploaded!', null, { toastLife: 180000 });
      },
      (err) => {
        if (err.status == 200) {
          this.toastsManager.success('Results uploaded!');
        } else {
          this.toastsManager.error('Error uploading results, please try again.');
        }
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
        this.toastsManager.success('Password changed!');
      },
      (err) => {
        if (err.status == 200) {
          this.toastsManager.success('Password changed!');
        } else {
          this.toastsManager.error('Error changing password, please try again.');
        }
      }
    );
  }

  public uploadBackground(event) {
    const fileList: FileList = event.target.files;
    if (fileList.length === 0) {
      return;
    }
    const file = fileList[0];

    const formData = new FormData();
    formData.append('img', file, file.name);
    const headers = new Headers();
    /** No need to include Content-Type in Angular 4 */
    const params = new HttpParams();
    const options = {
      headers: headers,
      params: params,
      reportProgress: true,
      Timeout: 1000
    };
    this.adminservice.uploadBackground(formData, options).subscribe(
      (res) => {
        this.toastsManager.success('background has been uploaded!', null, { toastLife: 180000 });
      },
      (err) => {
        this.toastsManager.error('Error uploading background, please try again.');
      }
    );
  }

  public submitStyle() {
    this.adminservice.updateButtonStyles(this.buttonStyleSelected).subscribe(
      res => {
        this.toastsManager.success(`successfully updated button with guid ${this.buttonStyleSelected.button_id}`, 'Success');
      },
      err => {
        this.toastsManager.error(`error updating button with guid ${this.buttonStyleSelected.button_id}`, 'Error');
      }
    )
  }
} 
