import { Component, OnInit } from '@angular/core';
import { IxonService } from '../services/ixon.service';

@Component({
  selector: 'app-company',
  templateUrl: './company.component.html',
  styleUrls: ['./company.component.css']
})
export class CompanyComponent implements OnInit {

  allCompanies: any

  constructor(
    public ixonService: IxonService
  ) { }

  ngOnInit() {
    this.ixonService.makeRequest('companies')
      .toPromise().then(data =>{
        this.allCompanies = data
        console.log(data);
      })
  }

  onCompany(companyId){
    console.log(companyId);
    this.ixonService.chosenCompany = companyId
    this.ixonService.getToken()
  }
}
