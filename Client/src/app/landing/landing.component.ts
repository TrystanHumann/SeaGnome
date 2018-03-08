import { Component, OnInit } from '@angular/core';
import { Streamer } from '../models/Streamer.model';
import { SafeUrl, DomSanitizer } from '@angular/platform-browser';
import { AdminService } from '../admin/admin.service';
import { RouterLink, NavigationEnd } from '@angular/router';
import { Router } from '@angular/router';

declare var ga: Function;


@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.css']
})
export class LandingComponent implements OnInit {
  public streamers: Streamer[];
  public trustedUrl: Array<SafeUrl> = [];
  public trustedUrlChat: Array<SafeUrl> = [];
  public user: string;
  public ga: any;
  constructor(public adminservice: AdminService,
    private sanitizer: DomSanitizer, private router: Router) {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        ga('set', 'page', event.urlAfterRedirects);
        ga('send', 'pageview');
      }
    });
  }

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

        streamers.forEach(streamer => {
          // tslint:disable-next-line:max-line-length
          this.trustedUrlChat.push(this.sanitizer.bypassSecurityTrustResourceUrl('http://www.twitch.tv/' + streamer.tag + '/chat?darkpopout'));
        });
      }
    );
  }

  public enterPress() {
    this.router.navigate(['user', this.user]);
  }
}
