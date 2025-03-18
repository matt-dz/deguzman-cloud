export interface LoginForm {
	email: string;
	password: string;
}

export interface LoginErrors {
	email: boolean;
	password: boolean;
}

export interface SignupForm {
	firstName: string;
	lastName: string;
	email: string;
	password: string;
}

export interface SignupErrors {
	firstName: boolean;
	lastName: boolean;
	email: boolean;
	password: boolean;
}
