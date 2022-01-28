import { Injectable } from '@angular/core';
import {FormGroup,NgForm} from '@angular/forms';
@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private _registerUrl = "";
  private _loginUrl = "";
  public loginPageActive = false;
  public response:string = "Success";
  constructor() {
   }

  registerUser(user:FormGroup) :any {

    if (user.valid && user.value.password != user.value.confirmPassword){
      this.response = 'password inconsistent';
      console.log("Error:請確認輸入密碼一致！");
      return this.response
    } else if (user.valid){
      this.response = 'Success';
      console.log("Register Success!");
      console.log(user.value);
      return this.response
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
