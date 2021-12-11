import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-block-header',
  templateUrl: './block-header.component.html',
  styleUrls: ['./block-header.component.scss']
})
export class BlockHeaderComponent implements OnInit {

  @Input() title: string;
  constructor() {
    this.title = '';
   }

  ngOnInit(): void {
  }

}
