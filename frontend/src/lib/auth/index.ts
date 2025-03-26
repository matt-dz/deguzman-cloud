export const sessionCookieName = 'session';

export enum AuthRole {
	Admin = 'admin',
	User = 'user'
}

export type AuthPayload = {
	role: AuthRole;
};
