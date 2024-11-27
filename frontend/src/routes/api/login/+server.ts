import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ request }) {
	const { username, password } = await request.json();

	const response = await fetch(VITE_SERVER_ADDRESS + '\\login', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ username: username, password: password })
	});

	if (!response.ok) {
		const errorData = await response.json();
		return new Response(JSON.stringify({ message: errorData }), { status: 400 });
	}

	return new Response(
		JSON.stringify({ message: 'Login successful!', body: await response.json() }),
		{
			status: 200
		}
	);
}
