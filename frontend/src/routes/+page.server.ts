import * as db from '$lib/server/database';
import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ cookies }) => {
	const sessionToken = cookies.get('session_token');

	if (sessionToken) {
		const user = await db.getUserBySessionToken(sessionToken);

		if (user) {
			redirect(302, '/main');
		}
	}
};
