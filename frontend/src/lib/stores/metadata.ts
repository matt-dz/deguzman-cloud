import { readable } from 'svelte/store';

const prod = import.meta.env.PROD;

export const sessionCookieName = readable(prod ? '__Secure-sessionId' : 'sessionId');
