import { VITE_SERVER_ADDRESS } from '$env/static/private';
import { hashWithSHA256 } from '$lib/Functions';

export async function POST({ request, cookies }) {
	const { username, password } = await request.json();
	let response;

	try {
		const hashedPassword = await hashWithSHA256(password);
		response = await fetch(VITE_SERVER_ADDRESS + '\\login', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ username: username, password: hashedPassword })
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

	const body = await response.json();
	const { expires_at, session_token } = body;

	cookies.set('session_token', session_token, {
		httpOnly: true,
		path: '/'
	});

	return new Response(JSON.stringify({ message: 'Login successful!', body: body }), {
		status: 200
	});
}
