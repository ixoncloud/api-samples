import { Component, OnInit } from '@angular/core';
import { IxonService } from '../services/ixon.service';

@Component({
  selector: 'app-agent',
  templateUrl: './agent.component.html',
  styleUrls: ['./agent.component.css']
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
