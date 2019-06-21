import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { IxonService } from '../services/ixon.service';
import { Router } from '@angular/router';
import * as  base64 from "base-64";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup

  constructor(
    public fb: FormBuilder,
    private ixonService: IxonService,
    private router: Router,
  ) { 
    this.loginForm = this.fb.group({
      "email": [''],
      "password": [''],
      "otpassword": ['']
    })
  }
  
  ngOnInit() {  }

  set_authorization_header(method, val){
    var header: string = method + ' ' + val;
    this.ixonService.authToken = header
  }

  onLogin(){
    if(this.loginForm.invalid){
      return
    }
    this.set_authorization_header('Basic', base64.encode(this.loginForm.value.email + ':' + (this.loginForm.value.otpassword ? this.loginForm.value.otpassword : '') + ':' + this.loginForm.value.password));
    
    this.ixonService.makeRequest("access/tokens")
      .toPromise().then(data =>{
        let test: any = data
        if (test.status === "success") {
          this.ixonService.isLoggedIn = true
          this.router.navigate(['company']);
        }
      })
  }
}
