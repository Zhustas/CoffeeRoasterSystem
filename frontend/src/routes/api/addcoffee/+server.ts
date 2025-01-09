import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ request, cookies }) {
	const sessionToken = cookies.get('session_token');
	const body = await request.json();

	const response = await fetch(`${VITE_SERVER_ADDRESS}/addcoffee`, {
		method: 'POST',
		headers: {
			Cookie: `session_token=${sessionToken}`
		},
		body: JSON.stringify(body)
	});

	if (!response.ok) {
		const errorData = await response.json();
		return new Response(JSON.stringify({ message: 'Error in adding coffee', body: errorData }), {
			status: 400
		});
	}

	return new Response(
		JSON.stringify({ message: 'Coffee added successfully!', body: await response.json() }),
		{
			status: 200
		}
	);
}
