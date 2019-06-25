import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptModule } from "nativescript-angular/nativescript.module";
import { ReactiveFormsModule }   from '@angular/forms';


import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";
import { TagComponent } from './tag/tag.component';
import { LoginComponent } from './login/login.component';
import { DeviceComponent } from './device/device.component';
import { DataComponent } from './data/data.component';
import { CompanyComponent } from './company/company.component';
import { AgentComponent } from './agent/agent.component';


import { NativeScriptFormsModule } from "nativescript-angular/forms";
import { NativeScriptHttpClientModule } from "nativescript-angular/http-client";

@NgModule({
    bootstrap: [
        AppComponent
    ],
    imports: [
        NativeScriptModule,
        AppRoutingModule,
        NativeScriptHttpClientModule,
        NativeScriptFormsModule,
        ReactiveFormsModule,
    ],
    declarations: [
        AppComponent,
        TagComponent,
        LoginComponent,
        DeviceComponent,
        DataComponent,
        CompanyComponent,
        AgentComponent,
    ],
    providers: [],
    schemas: [
        NO_ERRORS_SCHEMA
    ]
})
/*
Pass your application module to the bootstrapModule function located in main.ts to start your app
*/
export class AppModule { }
