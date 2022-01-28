import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service'
import {FormGroup, FormBuilder, Validators , FormControl} from '@angular/forms';
import {forbiddenNameValidator,MatchPasswordValidator,MatchEmailValidator} from '../validator/form.validator'
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})

export class RegisterComponent implements OnInit {
  backendResponse = ''
  inconsistentPassword= false
  
  invalidUserName = /admin|password/
  matchEmailRegex = /^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$/
  matchPasswordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d).*$/

  constructor(
    private _auth:AuthService,
    private formBuilder: FormBuilder
    ) { }
    get userName() {
      return this.registrationForm.get('userName') as FormControl;
    }
    get userEmail(){
      return this.registrationForm.get('userEmail') as FormControl;
    }

    get password(){
      return this.registrationForm.get('password') as FormControl;
    }
  
  registrationForm = this.formBuilder.group({
    userName: ['',[Validators.required, Validators.minLength(2),forbiddenNameValidator(this.invalidUserName)]],
    userEmail: ['',[Validators.required,MatchEmailValidator(this.matchEmailRegex)]],
    password: ['',[Validators.required,Validators.minLength(8),MatchPasswordValidator(this.matchPasswordRegex)]],
    confirmPassword: '*至少8個字元與1個數字',
    phoneNumber:Validators.minLength(8)
  })

  ngOnInit(): void {
  }
  
  onSubmit(data:FormGroup){
    this.backendResponse = this._auth.registerUser(data);
    if (this.backendResponse == 'password inconsistent'){
      this.inconsistentPassword = true
      console.log("註冊失敗:",this.backendResponse);
    } else if(this.backendResponse == 'Success') {
      this.inconsistentPassword = false
      console.log("註冊成功?:",data.valid);
    }
  }

  openLoginDialog(){
    this._auth.loginPageActive = true;
  }
  
/**
 * When invalid password notification show up, invalidPassword=True
 * Turn off the alerm by user click and set it back to "false"
 */
  closeAlert(){
    if (this.inconsistentPassword){
      this.inconsistentPassword = !this.inconsistentPassword;
    }
  }
}
