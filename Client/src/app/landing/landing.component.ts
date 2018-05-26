import { Component, OnInit } from '@angular/core';
import { Streamer, ButtonStyle } from '../models/Streamer.model';
import { SafeUrl, DomSanitizer } from '@angular/platform-browser';
import { AdminService } from '../admin/admin.service';
import { RouterLink, NavigationEnd } from '@angular/router';
import { Router } from '@angular/router';
import { style } from '@angular/animations';

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
  public streamerNames: Array<string> = [];
  public user: string;
  public singleStreamerDisplay;
  public ga: any;
  public buttonStyleArray : ButtonStyle[];
  public buttonStyleMap : Map<string, ButtonStyle>;
  public webpageTitle: string;

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
    this.singleStreamerDisplay = null;
    this.buttonStyleMap = new Map<string, ButtonStyle>();
    this.getWebpageTitle();
    this.setStreamers();
    this.getButtonStyles();
  }

  public getWebpageTitle(): void {
    this.adminservice.getTitle().subscribe(res => {
      if (res.length > 0) {
        this.webpageTitle = res[0];
      }
    });
  }

  public getButtonStyles() : void {
    this.adminservice.getButtonStyles().subscribe(
      res => {
        this.buttonStyleArray = res;

        // Creating map to easily access values by id
        this.buttonStyleArray.forEach(style => {
          this.buttonStyleMap.set(style.button_id, style);
        });
      },
      err => {
        this.buttonStyleArray = [];
      });
  }

  public setStreamers() {
    this.adminservice.getStreamers().subscribe(
      (streamers: Array<Streamer>) => {
        if (!streamers) {
          return;
        }
        streamers.forEach(streamer => {
          // tslint:disable-next-line:max-line-length
          this.trustedUrl.push(this.sanitizer.bypassSecurityTrustResourceUrl('https://player.twitch.tv/?channel=' + streamer.tag + '&muted=true'));
        });

        streamers.forEach(streamer => {
          // tslint:disable-next-line:max-line-length
          this.trustedUrlChat.push(this.sanitizer.bypassSecurityTrustResourceUrl('http://www.twitch.tv/embed/' + streamer.tag + '/chat'));
        });

        streamers.forEach(streamer => {
          // tslint:disable-next-line:max-line-length
          this.streamerNames.push(streamer.tag);
        });
      }
    );
  }

  public setsingleStreamerDisplay(streamer : string): void {
    this.singleStreamerDisplay = streamer;
  }

  public getSingleStreamerStreamChannelURL(): SafeUrl {
    return this.sanitizer.bypassSecurityTrustResourceUrl(`https://player.twitch.tv/?channel=${this.singleStreamerDisplay}`);
  }

  public getSingleStreamerChatURL(): SafeUrl {
    return this.sanitizer.bypassSecurityTrustResourceUrl(`http://www.twitch.tv/embed/${this.singleStreamerDisplay}/chat`);
  }
  
  public enterPress() {
    this.router.navigate(['user', this.user]);
  }
}
