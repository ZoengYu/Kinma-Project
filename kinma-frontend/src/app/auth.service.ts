import { Injectable } from '@angular/core';
import {FormGroup,NgForm} from '@angular/forms';
import {HttpClient} from '@angular/common/http';

export interface RegisterResponse{
  username          : string;
  email             : string;
  phone             : string;
  passwordChangeAt  : any;
  createdAt         : any;
};

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
    return this.http.post<any>(this.apiServiceUrl + "/users", 
      { "username"  :user.value.userName, 
        "password"  :user.value.password,
        "email"     :user.value.userEmail,
        "phone"     :user.value.phoneNumber
      })
  }

  loginUser(user:NgForm) {
    return this.http.post<RegisterResponse>(this.apiServiceUrl + "/users/login", user.value)
  }

  isLoginPageActive() {
    return this.loginPageActive;
  }
}
