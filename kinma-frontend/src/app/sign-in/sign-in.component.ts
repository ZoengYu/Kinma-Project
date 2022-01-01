import { Component, OnInit,DoCheck } from '@angular/core';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.scss']
})
export class SignInComponent implements OnInit,DoCheck {

  loginPageActive:boolean = true;

  constructor(
    private _authService:AuthService
  ) { }

  ngOnInit(): void {
  }

  ngDoCheck(): void {
    this.loginPageActive = this._authService.isLoginPageActive();
  }

  isLoginPageActive(){
    return this.loginPageActive;
  }

  isRegisterPageActive(){
    return !this.loginPageActive;
  }
}
