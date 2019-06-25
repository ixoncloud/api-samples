import { NgModule } from "@angular/core";
import { NativeScriptRouterModule } from "nativescript-angular/router";
import { Routes } from "@angular/router";

import { LoginComponent } from './login/login.component';
import { CompanyComponent } from './company/company.component';
import { AgentComponent } from './agent/agent.component';
import { DeviceComponent } from './device/device.component';
import { AuthGuard } from './auth.guard';
import { DataComponent } from './data/data.component';
import { TagComponent } from './tag/tag.component';

export const routes: Routes = [
  {
      path: 'login',
      component: LoginComponent,
  },
  {
    path: '',
    canActivate: [AuthGuard],
    children: [
      {
          path: 'company',
          component: CompanyComponent,
      },
      {
          path: 'agent',
          component: AgentComponent,
      },
      {
          path: 'device',
          component: DeviceComponent,
      },
      {
          path: 'tag',
          component: TagComponent,
      },
      {
          path: 'data',
          component: DataComponent,
      },
    ]
  },
  { path: '**', redirectTo: '/company' }
];
@NgModule({
    imports: [NativeScriptRouterModule.forRoot(routes)],
    exports: [NativeScriptRouterModule]
})
export class AppRoutingModule { }
