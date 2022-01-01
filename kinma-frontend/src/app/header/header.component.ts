import { Component, OnInit } from '@angular/core';
import {MatDialog} from '@angular/material/dialog';
import { SignInComponent } from '../sign-in/sign-in.component'
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  constructor(
    private dialog: MatDialog,
    private _authService: AuthService,
  ) { }

  ngOnInit(): void {
  }


  openLoginDialog(){
    this._authService.loginPageActive = true;
    console.log('login page open!')
    const dialogRef = this.dialog.open(SignInComponent);
    dialogRef.afterClosed().subscribe(result => {
      console.log(`header: Dialog close: ${result}`)
    })
  }

  openRegisterDialog(){
    this._authService.loginPageActive = false;
    console.log('Register page open!')
    const dialogRef = this.dialog.open(SignInComponent);
    dialogRef.afterClosed().subscribe(result => {
      console.log(`header: Dialog close: ${result}`)
    })
  }

}
