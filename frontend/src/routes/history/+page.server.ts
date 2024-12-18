import { USER_CUSTOMER } from '$lib/constants/UserTypeConstants';
import type { Order } from '$lib/order';
import * as db from '$lib/server/database';
import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ cookies, url }) => {
	const sessionToken = cookies.get('session_token');

	// Gauti vartotojo ID
	const user = await db.getUserBySessionToken(sessionToken);

	if (user) {
		if (user.role !== USER_CUSTOMER) {
			redirect(302, '/main');
		}

		const id = user.Id;
		const orders: Order[] = await db.getUserOrders(sessionToken, id);

		const fakeOrderItems = await Promise.all(
			orders.map((order) => db.getOrderItems(sessionToken, order.id))
		);

		let orderItems = fakeOrderItems.map((orderItem) => orderItem.order_items);

		let coffees = await Promise.all(
			orderItems.map((orderItem) => db.getCoffee(sessionToken, orderItem.coffee_id))
		);

		return { orders, orderItems, coffees };
	}

	redirect(302, `/?redirectTo=${url.pathname}`);
};
