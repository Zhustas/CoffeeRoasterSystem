import { VITE_SERVER_ADDRESS } from '$env/static/private';

export async function POST({ request }) {
	const user = await request.json();

	const response = await fetch(VITE_SERVER_ADDRESS + '\\register', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(user)
	});

	if (!response.ok) {
		const errorData = await response.json();
		return new Response(JSON.stringify({ message: errorData }), { status: 400 });
	}

	return new Response(JSON.stringify({ message: 'Registration successful!' }), { status: 200 });
}
