import { VITE_SERVER_ADDRESS } from '$env/static/private';
import type { Coffee } from '../Coffee';

async function getCoffees() {
	const response = await fetch(`${VITE_SERVER_ADDRESS}/coffeeinventory`, {
		method: 'GET'
	});

	const { coffeeList, title }: { coffeeList: Coffee[]; title: string } = await response.json();

	return coffeeList;
}

export { getCoffees };
