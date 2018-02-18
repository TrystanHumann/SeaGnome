import { Component, OnInit } from '@angular/core';
import { AdminService } from './admin.service';
import { Streamer } from '../models/Streamer.model';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
import { HttpParams } from '@angular/common/http';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';

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
  public fakeListOfData = new Array<String>('MOM1', 'MOM2', 'MOM3');
  public startInsertRows = { id: 0, game: 'Insert Game', competitor: ['Iateyourpie', 'Spikevegeta'] };
  public closeResult: string;
  public eventCreateRows: any[] = [];
  public showEditable: boolean;
  public editRowId: any;
  public editColumnId: any;

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
    console.log(this.eventName);
    console.log(this.eventCreateRows);
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
    // console.log(this.editRowId, this.editColumnId);
    console.log(this.eventCreateRows);
  }
}
