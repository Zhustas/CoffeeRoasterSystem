import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ params, cookies }) {
	const sessionToken = cookies.get('session_token');

	const response = await fetch(`${VITE_SERVER_ADDRESS}/deleteuser/${params.slug}`, {
		method: 'POST',
		headers: {
			Cookie: `session_token=${sessionToken}`
		}
	});

	if (!response.ok) {
		const errorData = await response.json();
		return new Response(JSON.stringify({ message: 'Error in deleting user', body: errorData }), {
			status: 400
		});
	}

	return new Response(
		JSON.stringify({ message: 'User deleted successfully!', body: await response.json() }),
		{
			status: 200
		}
	);
}
