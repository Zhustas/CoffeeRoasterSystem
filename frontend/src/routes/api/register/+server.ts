import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ request }) {
	const user = await request.json();
	let response;

	try {
		response = await fetch(VITE_SERVER_ADDRESS + '\\register', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(user)
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

		if (errorData.error === 'This username or email is already taken. Please choose another.') {
			return new Response(null, { status: 409 });
		}

		if (errorData.error === 'Failed to register user. Please try again later.') {
			return new Response(null, { status: 500 });
		}
	}

	return new Response(JSON.stringify({ message: 'Registration successful!' }), {
		status: 200
	});
}
