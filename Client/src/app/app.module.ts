import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { UserProfileComponent } from './user-profile/user-profile.component';
import { RouteRoutingModule } from './route/route-routing.module';
import { LandingComponent } from './landing/landing.component';
import { AdminComponent } from './admin/admin.component';
import { AdminService } from './admin/admin.service';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { UserProfileService } from './user-profile/user-profile.service';
import { ToastModule } from 'ng2-toastr/ng2-toastr';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { Observable } from 'rxjs/Observable';
import { TimeoutError } from 'rxjs/util/TimeoutError';
import 'rxjs/add/operator/timeout';
import { InjectionToken, Injectable, Inject } from '@angular/core';
import { HttpHandler, HttpRequest, HttpInterceptor, HttpEvent } from '@angular/common/http';

const DEFAULT_TIMEOUT = new InjectionToken<number>('defaultTimeout');
const defaultTimeout = 5000000;

@Injectable()
export class TimeoutInterceptor implements HttpInterceptor {
  // tslint:disable-next-line:no-shadowed-variable
  constructor(@Inject(DEFAULT_TIMEOUT) protected defaultTimeout) { }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const timeout = Number(req.headers.get('timeout')) || this.defaultTimeout;
    return next.handle(req).timeout(timeout);
  }
}

@NgModule({
  declarations: [
    AppComponent,
    UserProfileComponent,
    LandingComponent,
    AdminComponent
  ],
  imports: [
    BrowserModule,
    NgbModule.forRoot(),
    FormsModule,
    RouteRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ToastModule.forRoot()
  ],
  providers: [AdminService, UserProfileService,
    [{ provide: HTTP_INTERCEPTORS, useClass: TimeoutInterceptor, multi: true }],
    [{ provide: DEFAULT_TIMEOUT, useValue: defaultTimeout }]],
  bootstrap: [AppComponent]
})
export class AppModule { }
