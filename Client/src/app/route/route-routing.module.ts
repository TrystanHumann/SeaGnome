import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { UserProfileComponent } from '../user-profile/user-profile.component';
import { LandingComponent } from '../landing/landing.component';

const routes: Routes = [
  { path: '', component: LandingComponent},
  { path: 'user', component: UserProfileComponent},
  { path: '*', component: LandingComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class RouteRoutingModule { }
