import { Component, OnInit } from '@angular/core';
import { IxonService } from '../_services/ixon.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-tag',
  templateUrl: './tag.component.html',
  styleUrls: ['./tag.component.scss']
})
export class TagComponent implements OnInit {

  allTags: any

  constructor(
    private ixonService: IxonService,
    private router: Router
  ) { }

  ngOnInit() {
    this.ixonService.makeRequest(`agents/${this.ixonService.chosenAgent}/devices/${this.ixonService.chosenDevice}/data-tags`)
    .toPromise().then(data =>{
      this.allTags = data
      // console.log(data);
    })
  }

  onTag(data){
    // console.log(data);
    this.ixonService.chosenTag = data
    this.router.navigate(['data']);
  }

}
