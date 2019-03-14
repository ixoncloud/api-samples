import { Component, OnInit } from '@angular/core';
import { IxonService } from '../_services/ixon.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';

export interface IxonApi{
  count: number,
  data: any,
  links: any,
  status: string,
  type: string
}
// Here will all the requests to the data api be made
// We need some data before we can make a request
// first is better authentication we already set that in ixon.service
// second is a time zone this will be made in setTime()
// third is a request body this will be made in setBody()
@Component({
  selector: 'app-data',
  templateUrl: './data.component.html',
  styleUrls: ['./data.component.scss']
})
export class DataComponent implements OnInit {

  data: any
  today = new Date()

  constructor(
    public ixonService: IxonService,
    private http: HttpClient
  ) { }
  ngOnInit() {
    this.getData()
  }

  setTime(){
    // set the timezone from now till yesterday
    // so we get data from the last 24 hours
    var dd = this.today.getDate();
    var mm = this.today.getMonth() + 1; //January is 0!
    var yyyy = this.today.getFullYear();
    var hh = this.today.getHours();
    var mms = this.today.getMinutes();
    var ss = this.today.getSeconds();
    var ms = this.today.getMilliseconds();
    let to = yyyy + "-" + mm + "-" + dd + "T" + hh + ":" + mms + ":" + ss + "." + ms;
    var dd = this.today.getDate() - 1;
    let from = yyyy + "-" + mm + "-" + dd + "T" + hh + ":" + mms + ":" + ss + "." + ms;
    return `data?from=${from}&to=${to}`
  }

  setBody(){
    // for more information ask IXON
    return {[this.ixonService.chosenDevice]: {[this.ixonService.chosenTag.tagId]: {raw: [{ref: "b", extendedBoundary: true, limit: 1000, postAggr: "last"}]}}}
  }

  headers(){
    return new HttpHeaders({
      'Accept': 'application/json',
      'Authorization':  this.ixonService.authToken,
    });
  }

  
// Here will the request be made with all the data in it
// response wil be set in this.data
  getData(){
    this.http.post<IxonApi[]>("https://api.lsi.ams.dsn.ixon.net//" + this.setTime(), this.setBody(), { headers: this.headers() })
    .toPromise().then(data =>{
      this.data = data
      console.log(data);
    })
  }
}
