interface Order {
	readonly id: number;
	readonly user_id: number;
	total_amount: number;
	status: string;
	updated_at: Date;
	created_at: Date;
}

export type { Order };
