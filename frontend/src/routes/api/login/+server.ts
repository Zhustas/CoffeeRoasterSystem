import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ request }) {
	const { username, password } = await request.json();
	let response;

	try {
		response = await fetch(VITE_SERVER_ADDRESS + '\\login', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ username: username, password: password })
		});
	} catch (e: unknown) {
		const error = e as Error;
		if (error.name === 'TypeError') {
			// Tegu būna serverio klaida
			return new Response(null, { status: 500 });
		}
	}

	if (!response.ok) {
		const errorData = await response.json();

		return new Response(null, { status: response.status });
	}

	return new Response(
		JSON.stringify({ message: 'Login successful!', body: await response.json() }),
		{
			status: 200
		}
	);
}
