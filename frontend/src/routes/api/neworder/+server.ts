import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ request, cookies }) {
	const requestBody = await request.json();
	const sessionToken = cookies.get('session_token');

	const response = await fetch(`${VITE_SERVER_ADDRESS}/order`, {
		method: 'POST',
		headers: {
			Cookie: `session_token=${sessionToken}`
		},
		body: JSON.stringify(requestBody)
	});

	if (!response.ok) {
		const errorData = await response.json();
		return new Response(JSON.stringify({ message: 'Error in making an order', body: errorData }), {
			status: response.status
		});
	}

	return new Response(
		JSON.stringify({ message: 'Order created successfully!', body: await response.json() }),
		{
			status: 200
		}
	);
}
