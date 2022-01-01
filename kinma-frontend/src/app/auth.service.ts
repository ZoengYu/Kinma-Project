import { Injectable } from '@angular/core';
import {NgForm} from '@angular/forms';
@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private _registerUrl = "";
  private _loginUrl = "";
  public loginPageActive = false;
  constructor() { }

  registerUser(user:NgForm) {
    if (user.valid){
      console.log("Register Success!");
      console.log(user.value);
    }
  }

  loginUser(user:NgForm) {
    if (user.valid){
      console.log("Login Success!");
      console.log(user.value);
    }
  }

  isLoginPageActive() {
    return this.loginPageActive;
  }
}
