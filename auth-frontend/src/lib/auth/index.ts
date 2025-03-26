export const sessionCookieName = 'session';

export enum AuthRole {
	Admin = 'admin',
	User = 'user'
}

export type AuthPayload = {
	role: AuthRole;
};

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
	redirect: string;
};
