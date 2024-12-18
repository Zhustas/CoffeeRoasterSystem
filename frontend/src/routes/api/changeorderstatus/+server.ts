import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ request, cookies }) {
	const sessionToken = cookies.get('session_token');
	const { ID, status }: { ID: number; status: string } = await request.json();

	const response = await fetch(`${VITE_SERVER_ADDRESS}/updateorders/${ID}`, {
		method: 'POST',
		headers: {
			Cookies: `session_token=${sessionToken}`,
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ status })
	});

	if (!response.ok) {
		const errorData = await response.json();
		return new Response(
			JSON.stringify({ message: 'Error in updating order status', body: errorData }),
			{
				status: 400
			}
		);
	}

	return new Response(
		JSON.stringify({ message: 'Order status updated successfully!', body: await response.json() }),
		{
			status: 200
		}
	);
}
