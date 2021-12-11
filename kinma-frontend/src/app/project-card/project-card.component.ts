import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-project-card',
  templateUrl: './project-card.component.html',
  styleUrls: ['./project-card.component.scss']
})
export class ProjectCardComponent implements OnInit {
  @Input() labels:string[];
  @Input() title:string;
  @Input() owner:string;
  @Input() price:string;
  @Input() date:string;
  @Input() img:string;
  constructor() {
    this.labels=[];
    this.title='';
    this.owner='';
    this.price='';
    this.date='';
    this.img = '';
   }

  ngOnInit(): void {
    this.checkDate()
  }

  checkDate(){
    if (this.date=='0') {
      this.date = '已結束';
    } else {
      this.date ='剩餘'+this.date+'天';
    }
  }
}
