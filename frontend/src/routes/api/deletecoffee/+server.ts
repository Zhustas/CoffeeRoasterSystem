import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ request }) {
	const { ID }: { ID: number } = await request.json();

	const response = await fetch(`${VITE_SERVER_ADDRESS}/deletecoffee/${ID}`, {
		method: 'POST'
	});

	if (!response.ok) {
		const errorData = await response.json();
		return new Response(JSON.stringify({ message: 'Error in deleting coffee', body: errorData }), {
			status: 400
		});
	}

	return new Response(
		JSON.stringify({ message: 'Coffee deleted successfully!', body: await response.json() }),
		{
			status: 200
		}
	);
}
