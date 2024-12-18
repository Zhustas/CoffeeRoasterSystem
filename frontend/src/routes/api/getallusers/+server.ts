import * as db from '$lib/server/database';

export async function GET({ cookies }) {
	const sessionToken = cookies.get('session_token');

	const users = await db.getUsers(sessionToken);

	if (!users) {
		return new Response(JSON.stringify({ message: 'Error in getting orders' }), {
			status: 520
		});
	}

	return new Response(JSON.stringify({ users }), {
		status: 200
	});
}
