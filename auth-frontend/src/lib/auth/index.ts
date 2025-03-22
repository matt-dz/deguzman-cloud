export enum AuthRole {
	ADMIN = 'admin',
	USER = 'user'
}

export type SignupPayload = {
	first_name: string;
	last_name: string;
	email: string;
	password: string;
};

export type LoginPayload = {
	email: string;
	password: string;
};

export type LoginResponse = {
	redirectUrl: string;
};
