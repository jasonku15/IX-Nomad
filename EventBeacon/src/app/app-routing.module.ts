import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {EventListComponent} from './event-list/event-list.component';
import {EventLandingComponent} from './event-landing/event-landing.component';
import {EventCreateComponent} from './event-create/event-create.component';

const routes: Routes = [
  {
    path: 'event/show',
    component: EventListComponent
  },
  {
    path: 'event/create',
    component: EventCreateComponent
  },
  {
    path: 'event/land',
    component: EventLandingComponent
  }

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
