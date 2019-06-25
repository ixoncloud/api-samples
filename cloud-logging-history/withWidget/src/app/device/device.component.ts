import { Component, OnInit } from '@angular/core';
import { IxonService } from '../services/ixon.service';

@Component({
  selector: 'app-device',
  templateUrl: './device.component.html',
  styleUrls: ['./device.component.css']
})
export class DeviceComponent implements OnInit {

  allDevices: any

  constructor(
    private ixonService: IxonService
  ) { }

  ngOnInit() {
    this.ixonService.makeRequest(`agents/${this.ixonService.chosenAgent}/devices`)
    .toPromise().then(data =>{
      this.allDevices = data
      console.log(data);
    })
  }

  onDevice(deviceId){
    console.log(deviceId);
    this.ixonService.chosenDevice = deviceId
  }

}
