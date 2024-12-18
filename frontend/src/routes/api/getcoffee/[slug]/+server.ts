import * as db from '$lib/server/database';

export async function POST({ params, cookies }) {
	const sessionToken = cookies.get('session_token');

	const coffeeObject = await db.getCoffee(sessionToken, params.slug);

	if (!coffeeObject) {
		return new Response(JSON.stringify({ message: 'Error in getting coffee' }), {
			status: 520
		});
	}

	return new Response(JSON.stringify({ coffee: coffeeObject }), {
		status: 200
	});
}
