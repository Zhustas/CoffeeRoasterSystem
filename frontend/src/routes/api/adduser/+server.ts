import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ request, cookies }) {
	const sessionToken = cookies.get('session_token');
	const body = await request.json();

	const response = await fetch(`${VITE_SERVER_ADDRESS}/registernewuser`, {
		method: 'POST',
		headers: {
			Cookies: `session_token=${sessionToken}`,
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(body)
	});

	if (!response.ok) {
		const errorData = await response.json();
		return new Response(JSON.stringify({ message: 'Error in adding user', body: errorData }), {
			status: 400
		});
	}

	return new Response(
		JSON.stringify({ message: 'User added successfully!', body: await response.json() }),
		{
			status: 200
		}
	);
}
