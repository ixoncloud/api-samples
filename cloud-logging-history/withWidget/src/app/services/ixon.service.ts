import { Injectable } from '@angular/core';

import { HttpHeaders, HttpClient } from '@angular/common/http';
import * as appSettings from "tns-core-modules/application-settings";


export interface IxonApi{
  // This is the data we expect to get from the api
  status: string,
  data: any,
  count: number,
  links: any,
  type: string,
}

@Injectable({
  providedIn: 'root'
})
export class IxonService {

  isLoggedIn: boolean = false

  authToken: any
  authToken2: any
  // Here are all the variables set to make a request to the official ixon api
  // First we need an url to make the request to, this will look like this
  url: string = "https://api.ixon.net:443/"
  // Then we need all the headers to make a succesfull request
  header: string[] = [
    // Response
    'application/json',
    // Api version
    '1',
    // Application id Must be requested at https://www.ixon.cloud/support but is neccesary to make requests!
    // PASTE YOURE APPLICATION ID HERE
    ''

    // Autorization header will be set after login
    // Company id will also be set after login
  ]

  chosenCompany: string
  chosenAgent: string
  chosenDevice: string
  chosenTag: any
  
  constructor(
    public http: HttpClient,
  ) {
    // appSettings.setString("test", "1234 hoedje van hoedje van papier")
    // console.log("string is set");
   }

  setHeader(){
    if (!this.chosenCompany) {
    return new HttpHeaders({
      'Accept': this.header[0],
      'IXapi-Version' : this.header[1],
      'IXapi-Application': this.header[2],
      'Authorization':  this.authToken
    });
  }else{
    return new HttpHeaders({
      'Accept': this.header[0],
      'IXapi-Version' : this.header[1],
      'IXapi-Application': this.header[2],
      'Authorization':  this.authToken,
      'IXapi-Company': this.chosenCompany
    });
  }
  }

  makeRequest(data){
    return this.http.get<IxonApi[]>(this.url + data + "?fields=*", { headers: this.setHeader()})
  }

  // To make an request to the data api(api2) from ixon we need higher authentication
  // To get this we ask for an accestoken from api1
  getToken(){
    var body = {
      "expiresIn": 3600
    }
    return this.http.post("https://api.ixon.net:443/auth/tokens?fields=*", body,{ headers: this.setHeader()})
    .toPromise()
        .then(data => {
      var token: any = data
      // And set the token so both api's can make use of it
      this.authToken2 = "Bearer " + token.data.token;
    })
  }
}
