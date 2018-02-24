import { Component, OnInit } from '@angular/core';
import { Streamer } from '../models/Streamer.model';
import { SafeUrl, DomSanitizer } from '@angular/platform-browser';
import { AdminService } from '../admin/admin.service';
import { RouterLink } from '@angular/router';
import { Router } from '@angular/router';


@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.css']
})
export class LandingComponent implements OnInit {
  public streamers: Streamer[];
  public trustedUrl: Array<SafeUrl> = [];
  public user: string;

  constructor(public adminservice: AdminService,
    private sanitizer: DomSanitizer , private router: Router) { }

  ngOnInit() {
    this.setStreamers();
  }

  public setStreamers() {
    this.adminservice.getStreamers().subscribe(
      (streamers: Array<Streamer>) => {
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

  public enterPress() {
    this.router.navigate(['user', this.user]);
  }
}
