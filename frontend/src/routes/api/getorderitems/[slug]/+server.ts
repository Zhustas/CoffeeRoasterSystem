import * as db from '$lib/server/database';

export async function GET({ params, cookies }) {
	const sessionToken = cookies.get('session_token');

	const orderItems = await db.getOrderItems(sessionToken, params.slug);

	if (!orderItems) {
		return new Response(JSON.stringify({ message: 'Error in getting order items' }), {
			status: 520
		});
	}

	return new Response(JSON.stringify({ orderItems }), {
		status: 200
	});
}
