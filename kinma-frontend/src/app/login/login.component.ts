import { isEmptyExpression } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import {NgForm} from '@angular/forms';
import { AuthService } from '../auth.service'

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  constructor(private _authService: AuthService) { }

  ngOnInit(): void {
  }

  onSubmit(data:NgForm){
    this._authService.loginUser(data)
      .subscribe(
        res => {
          console.log('login success:',res)
        },
        err => {
          console.log('err:',err)
          return
        }
      );
  }

  openRegisterPage(){
    this._authService.loginPageActive = false;
  }
}
