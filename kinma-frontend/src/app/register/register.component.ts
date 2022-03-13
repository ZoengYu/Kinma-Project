import { Component, OnInit } from '@angular/core';
import { AuthService, RegisterResponse } from '../auth.service'
import {FormGroup, FormBuilder, Validators , FormControl} from '@angular/forms';
import {forbiddenNameValidator,MatchPasswordValidator,MatchEmailValidator,forbiddenPhoneValidator} from '../validator/form.validator'
import { HttpErrorResponse } from '@angular/common/http';
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})

export class RegisterComponent implements OnInit {
  inconsistentPassword = false
  formInputInvalid = false
  
  invalidUserName = /admin|password/
  matchEmailRegex = /^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$/
  matchPasswordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d).*$/
  //Taiwan phone number format
  matchPhoneNumber = /^09\d{2}-?\d{3}-?\d{3}$/

  constructor(
    private _auth:AuthService,
    private fb: FormBuilder
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
    get phoneNumber(){
      return this.registrationForm.get('phoneNumber') as FormControl;
    }
  
  registrationForm = this.fb.group({
    userName        : ['',[Validators.required, Validators.minLength(2), forbiddenNameValidator(this.invalidUserName)]],
    userEmail       : ['',[Validators.required, MatchEmailValidator(this.matchEmailRegex)]],
    password        : ['',[Validators.required, Validators.minLength(8), MatchPasswordValidator(this.matchPasswordRegex)]],
    confirmPassword : '*至少8個字元與1個數字',
    phoneNumber     : ['',[Validators.required, Validators.minLength(8), forbiddenPhoneValidator(this.matchPhoneNumber)]]
  })

  ngOnInit(): void {
  }
  
  onSubmit(data:FormGroup){
    if (data.invalid){
      this.formInputInvalid = true
      return
    }
    
    this._auth.registerUser(data)
      .subscribe(
        (res: RegisterResponse) => {
          console.log('註冊成功:',res)
          this.openLoginDialog()
        },
        (error: HttpErrorResponse) => {
          console.log('error response', error)
          return
        },
      )
  };

  openLoginDialog(){
    this._auth.loginPageActive = true;
  }
  
/**
 * When invalid password notification show up, invalidPassword=True
 * Turn off the alerm by user click and set it back to "false"
 */
  closePasswordAlert(){
    if (this.inconsistentPassword){
      this.inconsistentPassword = !this.inconsistentPassword;
    }
  }

  /**
 * When form is not fill out then notification show up, formInputInvalid=True
 * Turn off the alerm by user clicked and set it back to "false"
 */
  closeInvalidAlert(){
    if (this.formInputInvalid){
      this.formInputInvalid = !this.formInputInvalid;
    }
  }
}
