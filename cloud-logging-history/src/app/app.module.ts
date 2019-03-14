import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule }   from '@angular/forms';
import { HttpClientModule }    from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { CompanyComponent } from './company/company.component';
import { AgentComponent } from './agent/agent.component';
import { DeviceComponent } from './device/device.component';
import { DataComponent } from './data/data.component';
import { TagComponent } from './tag/tag.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    CompanyComponent,
    AgentComponent,
    DeviceComponent,
    DataComponent,
    TagComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
