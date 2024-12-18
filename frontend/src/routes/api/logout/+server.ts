// src/routes/api/logout/+server.js
export async function GET() {
	return new Response(null, {
		status: 200,
		headers: {
			'Set-Cookie': 'session_token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT'
		}
	});
}
