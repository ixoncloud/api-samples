import { Component, OnInit } from '@angular/core';
import { IxonService } from '../_services/ixon.service';

@Component({
  selector: 'app-agent',
  templateUrl: './agent.component.html',
  styleUrls: ['./agent.component.scss']
})
export class AgentComponent implements OnInit {

  allAgents: any

  constructor(
    private ixonService: IxonService
  ) { }

  ngOnInit() {
    this.ixonService.makeRequest('agents')
    .toPromise().then(data =>{
      this.allAgents = data
      console.log(data);
    })
  }

  onAgent(agentId){
    console.log(agentId);
    this.ixonService.chosenAgent = agentId
  }
}
