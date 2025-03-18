export interface Build {
	arguments: string[];
	image: string;
}

export interface Service {
	port: number;
	target_port: number;
}

export interface Config {
	build: Build;
	env: Record<string, string>;
	mounts: string[];
	name: string;
	services: Service[];
}

export interface Deployment {
	id: string;
	name: string;
	config: Config;
	user_id: string;
	status: string;
	created_at: string;
	updated_at: string;
}
