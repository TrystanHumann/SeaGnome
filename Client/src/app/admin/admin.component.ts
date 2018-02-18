import { Component, OnInit } from '@angular/core';
import { AdminService } from './admin.service';
import { Streamer } from '../models/Streamer.model';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
import { HttpParams } from '@angular/common/http';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { EventRequest } from '../models/EventRequest.model';
import { EventResponse } from '../models/EventResponse.model';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})



export class AdminComponent implements OnInit {
  public streamers: Streamer[];
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
  public manageEventsObject = { updatePredictions: null, updateEventResults: null, completeEvent: null, deleteEvent: null };
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
    this.setStreamers();
    this.eventCreateRows.push(this.startInsertRows);
    this.getEvents();
  }

  public setStreamers() {
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

  public uploadExcel(event) {
    console.log(this.eventName);
    console.log('hello');
    console.log(event);
    const fileList: FileList = event.target.files;
    if (fileList.length === 0) {
      return;
    }
    const file = fileList[0];

    const formData = new FormData();
    formData.append('file', file, file.name);
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
        console.log(res);
      }
    );
  }

  // addNewRow Expands the modal's row
  public addNewRow() {
    console.log('clicky');
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
}
