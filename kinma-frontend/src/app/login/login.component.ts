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
    this._authService.loginUser(data);
    console.log(data.value)
    console.log(data.valid)
  }

  openRegisterPage(){
    this._authService.loginPageActive = false;
  }
}
