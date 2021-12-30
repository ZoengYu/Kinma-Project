import { Component, OnInit } from '@angular/core';
import { LoginComponent } from '../login/login.component'
import {MatDialog} from '@angular/material/dialog';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  constructor(
    public dialog: MatDialog
  ) { }

  ngOnInit(): void {
  }


  openLoginDialog(){
    console.log('login page open!')
    this.dialog.open(LoginComponent);
  }

}
