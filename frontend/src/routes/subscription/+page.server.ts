import { USER_CUSTOMER } from '$lib/constants/UserTypeConstants';
import * as db from '$lib/server/database';
import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params, cookies, url }) => {
	const sessionToken = cookies.get('session_token');

	const user = await db.getUserBySessionToken(sessionToken);

	if (!user) {
		redirect(302, `/?redirectTo=${url.pathname}`);
	}

	if (user.role !== USER_CUSTOMER) {
		redirect(302, '/main');
	}
};
