import { AbstractControl, ValidatorFn } from "@angular/forms";

export function forbiddenNameValidator(forbiddenName: RegExp): ValidatorFn{
	return (control: AbstractControl): {[key:string]:any} | null => {
		const forbidden = forbiddenName.test(control.value);
		return forbidden ? { 'forbiddenName': {value: control.value}} : null;
	}
}

export function forbiddenPhoneValidator(forbiddenPhone: RegExp): ValidatorFn{
	return (control: AbstractControl): {[key:string]:any} | null => {
		const forbidden = forbiddenPhone.test(control.value);
		return forbidden ? null : { 'fordiddenPhoneFormat': {value: control.value}};
	}
}

/**
 * if  the email regex matched, IsMatchEmail = True and return null, no need 
 * to show the warning on register form
 */
 export function MatchEmailValidator(matchEmailRegex: RegExp): ValidatorFn{
	return (control: AbstractControl): {[key:string]:any} | null => {
		const IsMatchEmail = matchEmailRegex.test(control.value);
		return IsMatchEmail ? null : { 'fordiddenEmailFormat': {value: control.value}};
	}
}

/**
 * if  the password regex matched, IsMatchPassword = True and return null, no need 
 * to show the warning on register form
 */
export function MatchPasswordValidator(matchPasswordRegex: RegExp): ValidatorFn{
	return (control: AbstractControl): {[key:string]:any} | null => {
		const IsMatchPassword = matchPasswordRegex.test(control.value);
		return IsMatchPassword ? null : { 'fordiddenPasswordFormat': {value: control.value}};
	}
}