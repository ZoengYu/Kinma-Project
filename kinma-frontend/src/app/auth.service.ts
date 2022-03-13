import { Injectable } from '@angular/core';
import {FormGroup,NgForm} from '@angular/forms';
import {HttpClient} from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private apiServiceUrl = "http://localhost:8080";

  public loginPageActive = false;
  public response:string = "Success";
  constructor(private http: HttpClient) {
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
    return this.http.post<any>(this.apiServiceUrl + "/users/login", user.value)
  }

  isLoginPageActive() {
    return this.loginPageActive;
  }
}
