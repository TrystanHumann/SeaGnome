import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { UserProfileComponent } from '../user-profile/user-profile.component';
import { LandingComponent } from '../landing/landing.component';
import { AdminComponent } from '../admin/admin.component';

const routes: Routes = [
  { path: '', component: LandingComponent },
  { path: 'user/:user', component: UserProfileComponent },
  { path: 'admin', component: AdminComponent },
  { path: '*', component: LandingComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class RouteRoutingModule { }
