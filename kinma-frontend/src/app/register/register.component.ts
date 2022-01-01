import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service'
import {NgForm} from '@angular/forms';
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})

export class RegisterComponent implements OnInit {

  constructor(private _auth:AuthService) { }

  ngOnInit(): void {
  }
  
  onSubmit(data:NgForm){
    this._auth.registerUser(data);
    console.log("註冊成功?:",data.valid)
  }

  openLoginDialog(){
    this._auth.loginPageActive = true;
  }
}
